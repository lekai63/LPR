package lpr

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"cloud.google.com/go/civil"
	"github.com/lekai63/lpr/models"
	dataframe "github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/exports"
)

// CreateInsToDB 将利息计划写入数据库（Create）
func (model *BankRepayPlanCalcModel) CreateInsToDB() error {
	df := model.Brps
	filterFn := dataframe.FilterDataFrameFn(func(vals map[interface{}]interface{}, row, nRows int) (dataframe.FilterAction, error) {
		if vals["id"] == nil {
			return dataframe.KEEP, nil
		}
		return dataframe.DROP, nil
	})
	newDf := df.Copy()
	newDf.Names()
	_, err := dataframe.Filter(context.TODO(), newDf, filterFn, dataframe.FilterOptions{InPlace: true})
	if err != nil {
		return err
	}
	fmt.Println("newDf:\n")
	fmt.Print(newDf.Table())
	ctx := context.TODO()
	conn := models.GlobalConn
	tx := conn.BeginTx()

	m := map[string]*string{
		"id":                    nil,
		"bank_loan_contract_id": &[]string{"bank_loan_contract_id"}[0],
		"plan_date":             &[]string{"plan_date"}[0],
		"plan_insterest":        &[]string{"plan_insterest"}[0],
		"plan_amount":           &[]string{"plan_amount"}[0],
		// 设为nil的字段不导出至sql
		"actual_date":      nil,
		"actual_principal": nil,
		"actual_amount":    nil,
		"actual_interest":  nil,
	}
	op := exports.SQLExportOptions{
		NullString: &[]string{"0"}[0],
		Range:      dataframe.Range{},
		/* PrimaryKey: &exports.PrimaryKey{
			PrimaryKey: "id",
			Value: func(row int, n int) *string {
				return nil
			},
		}, */
		BatchSize:      &[]uint{50}[0],
		SeriesToColumn: m,
		Database:       exports.PostgreSQL,
	}
	err = exports.ExportToSQL(ctx, tx, newDf, "bank_repay_plan", op)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// 将计划还款总额更新到数据库，仅针对之前已有principal的数据，之前如无principal，则计划还款总额plan_amount数据已通过CreateInsToDB写入数据库
func UpdateAmountToDB(model *BankRepayPlanCalcModel) (*BankRepayPlanCalcModel, error) {
	return nil, nil
}

// InitDataframe  Gen dataframe from BankRepayPlan Struct
func InitDataframe(brp models.BankRepayPlan) (df *dataframe.DataFrame, err error) {
	typ := reflect.TypeOf(brp)
	val := reflect.ValueOf(brp)
	for i := 0; i < val.NumField(); i++ {
		gormTag := typ.Field(i).Tag.Get("gorm")
		colname, typename := getColnameWithType(gormTag)
		switch typename {
		case "int":
			se := dataframe.NewSeriesInt64(colname, nil)
			if df == nil {
				df = dataframe.NewDataFrame(se)
			} else {
				df.AddSeries(se, nil)
			}

			// sliceInt64 = append(sliceInt64, *se)
		case "time":
			se := dataframe.NewSeriesTime(colname, nil)
			if df == nil {
				df = dataframe.NewDataFrame(se)
			} else {
				df.AddSeries(se, nil)
			}
		case "date":
			se := dataframe.NewSeriesGeneric(colname, civil.Date{}, nil)

			// 定义比较函数，以便用于排序
			f := func(a, b interface{}) bool {
				a1 := a.(civil.Date)
				b1 := b.(civil.Date)
				return a1.Before(b1)
			}
			se.SetIsLessThanFunc(f)
			if df == nil {
				df = dataframe.NewDataFrame(se)
			} else {
				df.AddSeries(se, nil)
			}

			/* // 定义string函数，用于显示和最后的exports to sql
			fs := func(val interface{}) string {
				if val == nil {
					return "NaN"
				}
				v := val.(civil.Date)
				return v.String()
			}
			se.SetValueToStringFormatter(fs) */

		default:
			err = fmt.Errorf("无法识别:%s", typename)
			return
		}

	}

	return

}

// NewCalcModel 从原始dataframe（含所有已还款记录），抽离出最近一期还款(row=0)+未还款记录
func NewCalcModel(model *BankRepayPlanCalcModel) (*BankRepayPlanCalcModel, error) {
	i, err := getLatestNilActualRowNum(model.Brps)
	if err != nil {
		model.Brps = nil
	}
	r := make([]dataframe.Range, 1)
	withLastPaidRcord := i - 1
	r[0] = dataframe.Range{&withLastPaidRcord, nil}
	df := model.Brps.Copy(r...)
	model.Brps = df
	return model, nil
}

// AddAccruedPrincipalSeries 添加应计本金列,用于计算此row的plan_interest
//假定截至9月10日，应计本金为100万，每季还本，9月11日根据还款计划归还10万，则9月11日row的应计本金仍写作100万，10月11日row应计本金写作90万。
func (model *BankRepayPlanCalcModel) AddAccruedPrincipalSeries(ctx context.Context) *BankRepayPlanCalcModel {
	brps := model.Brps
	errorColl := dataframe.NewErrorCollection()
	i, err := brps.NameToColumn("plan_principal")
	if err != nil {
		errorColl.AddError(err)
	}

	copiedSerie, ok := brps.Series[i].Copy().(*dataframe.SeriesInt64)
	sums := dataframe.NewSeriesInt64("accrued_principal", nil)
	if ok {
		for i = 0; i < copiedSerie.NRows(); copiedSerie.Remove(i) {
			sumfloat, err := copiedSerie.Sum(ctx)
			sum := int64(sumfloat)
			if err != nil {
				errorColl.AddError(err)
			}
			sums.Append(sum)
		}
		brps.AddSeries(sums, nil)
	}
	fmt.Printf("%s", errorColl)
	return model
}

// getLatestNilActualRowNum 返回第一笔实际未付的记录序号，如全部已付，则返回-1
func getLatestNilActualRowNum(df *dataframe.DataFrame) (int, error) {
	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
	df.Lock()
	defer df.Unlock()
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		if vals["actual_amount"] == nil {
			return *row, nil
		}
	}

	return -1, fmt.Errorf("无未还款记录，请检查合同是否已结束")

}

/* getColnameWithType 返回gormTag 对应的colname,type
type转换关系:
int-->int64
date-->civil.date
timestamp-->time.time
*/
func getColnameWithType(gormTag string) (colname, typename string) {
	slice := strings.Split(gormTag, ";")
	for _, s := range slice {
		if strings.HasPrefix(s, "column:") {
			temp := strings.Split(s, ":")
			colname = temp[1]
			continue
		}
		if strings.HasPrefix(s, "type:") {
			temp := strings.Split(s, ":")
			t := temp[1]
			switch {
			case strings.HasPrefix(t, "INT"):
				typename = "int"
				continue
			case strings.HasPrefix(t, "DATE"):
				typename = "date"
				continue
			case strings.HasPrefix(t, "TIMESTAMP"):
				typename = "time"
				continue
			default:
				typename = "unsupport type : " + t
				continue
			}

		}

	}
	return
}

// call-gorm-alias-properly ref https://stackoverflow.com/questions/53444434/how-to-call-gorm-alias-properly
/* func getGormAlias(obj interface{}, fieldName string) string {
	if field, ok := reflect.TypeOf(obj).FieldByName(fieldName); ok {
		return field.Tag.Get("gorm")
	}

	return ""
} */