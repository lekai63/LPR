package inscalc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	"github.com/antlabs/deepcopy"

	"github.com/lekai63/lpr/models"
	dataframe "github.com/rocketlaunchr/dataframe-go"

	// "github.com/rocketlaunchr/dataframe-go/exports"

	"github.com/rocketlaunchr/dataframe-go/imports"
)

// BankRepayPlanCalcModel  定义单个合同的计算模型
type BankRepayPlanCalcModel struct {
	// Bc BankLoanContractMini
	Bc models.BankLoanContract
	// Brps存储Bc合同项下的所有还款记录, Brp []BankRepayPlan
	Brps *dataframe.DataFrame
}

/* // BankLoanContractMini 提取BankLoanContract与利息计算相关的字段
type BankLoanContractMini struct {
	// ID,InterestCalcMethod ,BankName,LoanMethod ,CurrentRate对应bank_loan_contract 中的同名字段
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: true   col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:INT4;" json:"id"`
	//[ 4] interest_calc_method                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	InterestCalcMethod null.String `gorm:"column:interest_calc_method;type:VARCHAR;size:255;" json:"interest_calc_method"`
	//[ 5] bank_name                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	BankName string `gorm:"column:bank_name;type:VARCHAR;size:255;" json:"bank_name"`
	//[ 8] loan_method                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	LoanMethod null.String `gorm:"column:loan_method;type:VARCHAR;size:255;" json:"loan_method"`
	//[11] actual_start_date                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	ActualStartDate null.Time `gorm:"column:actual_start_date;type:DATE;" json:"actual_start_date"`
	//[16] current_rate                                   INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	CurrentRate int32 `gorm:"column:current_rate;type:INT4;" json:"current_rate"`
} */

// NewModel 根据bankLoanContractID 从数据库中获取数据 并生成 BankRepayPlanCalcModel
func NewModel(bankLoanContractID int32) (model BankRepayPlanCalcModel, err error) {
	// conn := models.GlobalConn
	// to change back to models.GlobalConn
	db, _ := gormInitForTest()

	// gen model.Bc
	var bc models.BankLoanContract
	bc.ID = bankLoanContractID
	db.Table("bank_loan_contract").First(&bc)
	// db.First(&bc)
	model.Bc = bc

	// gen model.Brps
	sqldb, _ := db.DB()
	tx, _ := sqldb.Begin()
	op := imports.SQLLoadOptions{
		// KnownRowCount: &[]int{13}[0],
		DictateDataType: map[string]interface{}{
			"id":                    int64(0),
			"bank_loan_contract_id": int64(0),
			"plan_date":             time.Unix(0, 0),
			"plan_amount":           int64(0),
			"plan_principal":        int64(0),
			"plan_interest":         int64(0),
			"actual_date":           time.Unix(0, 0),
			"actual_amount":         int64(0),
			"actual_principal":      int64(0),
			"actual_interest":       int64(0),
			"created_at":            time.Unix(0, 0),
			"updated_at":            time.Unix(0, 0),
			"accrued_principal":     int64(0),
		},
		Database: imports.PostgreSQL,
		Query:    `select * from "bank_repay_plan"` + `where bank_loan_contract_id =` + strconv.Itoa(int(bc.ID)),
	}
	brps, err := imports.LoadFromSQL(ctx, tx, &op)
	if err != nil {
		fmt.Printf("从数据库中读取数据组装dataframe时发生错误：%+v", err)
		return model, err
	}
	model.Brps = brps
	model.convTimeToDate()

	return model, nil

}

// AfterDay 仅提取某一日之后的dataframe
func (model *BankRepayPlanCalcModel) AfterDay(day civil.Date) (*BankRepayPlanCalcModel, error) {
	model.Sort("plan_date")
	df := model.Brps
	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true}) // Don't apply read lock because we are write locking from outside.
	df.Lock()
	n := -1
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}

		if d := (vals["plan_date"]).(civil.Date); d.After(day) {
			n = *row
			break
		}
	}
	df.Unlock()
	if n < 0 {
		return model, fmt.Errorf("不存在%s之后的还款计划", day)
	}
	df.Lock()
	newDf := df.Copy(dataframe.Range{&n, nil})
	df.Unlock()
	model.Brps = newDf
	return model, nil
}

// Update 筛选id为nil的值，插入到数据库
func (model *BankRepayPlanCalcModel) Update() error {

	df := model.Brps
	df.Lock()
	newDf := df.Copy()
	df.Unlock()

	fnToUpdate := dataframe.FilterDataFrameFn(func(vals map[interface{}]interface{}, row, nRows int) (dataframe.FilterAction, error) {
		if vals["id"] == nil {
			return dataframe.DROP, nil
		}
		return dataframe.KEEP, nil
	})

	newDf.Names()
	_, err := dataframe.Filter(ctx, newDf, fnToUpdate, dataframe.FilterOptions{InPlace: true})
	if err != nil {
		return err
	}
	/* 	fmt.Println("toCreateDf:")
	   	fmt.Print(newDf.Table()) */

	//  conn := models.GlobalConn
	//  tx := conn.BeginTx()

	sqldb, err := db.DB()
	check(err)
	tx, err := sqldb.Begin()
	check(err)

	m := map[string]*string{
		"id":                    &[]string{"id"}[0],
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
	op := SQLExportOptions{
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
		Database:       PostgreSQL,
	}

	err = ExportToSQL(ctx, tx, newDf, "bank_repay_plan", true, op)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err

}

// Insert 筛选id为nil的值，插入到数据库
func (model *BankRepayPlanCalcModel) Insert() error {

	df := model.Brps
	df.Lock()
	newDf := df.Copy()
	df.Unlock()

	fnToInsert := dataframe.FilterDataFrameFn(func(vals map[interface{}]interface{}, row, nRows int) (dataframe.FilterAction, error) {
		if vals["id"] == nil {
			return dataframe.KEEP, nil
		}
		return dataframe.DROP, nil
	})

	newDf.Names()
	_, err := dataframe.Filter(ctx, newDf, fnToInsert, dataframe.FilterOptions{InPlace: true})
	if err != nil {
		return err
	}

	//  conn := models.GlobalConn
	//  tx := conn.BeginTx()

	sqldb, err := db.DB()
	check(err)
	tx, err := sqldb.Begin()
	check(err)

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
	op := SQLExportOptions{
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
		Database:       PostgreSQL,
	}

	err = ExportToSQL(ctx, tx, newDf, "bank_repay_plan", false, op)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err

}

// AddAccruedPrincipal 重新计算应计本金并替换原Series
func (model *BankRepayPlanCalcModel) AddAccruedPrincipal() *BankRepayPlanCalcModel {
	model.Sort("plan_date")
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
				fmt.Printf("%s", errorColl)
			}
			sums.Append(sum)
		}
		brps.RemoveSeries("accrued_principal")
		brps.AddSeries(sums, nil)
	}

	return model
}

// TODO:根据计息方式不同 生成不同的计划还款日期
func (model *BankRepayPlanCalcModel) FillInsPlanDate() *BankRepayPlanCalcModel {
	method := model.Bc.InterestCalcMethod
	switch method.String {
	case "按月扣息":
		model.FillPlanDateMonthly()
	case "按季扣息":
		model.FillPlanDateSeasonly()
	default:

	}

	return model
}

// Sort 根据 model 的 fieldname 字段升序排列
func (model *BankRepayPlanCalcModel) Sort(fieldname string) *BankRepayPlanCalcModel {
	df := model.Brps

	df.Sort(ctx, []dataframe.SortKey{
		{Key: fieldname, Desc: false},
	})

	model.Brps = df
	return model
}

// FillPlanDateMonthly 对还未还款的记录（actual_date 为nil），生成每月21日的还息日期plan_date，并按planDate升序排序。
// 其他字段以nil进行填充 （bank_loan_contract_id 用 bc.id填充）
func (model *BankRepayPlanCalcModel) FillPlanDateMonthly() *BankRepayPlanCalcModel {
	brps := model.Brps
	col, _ := brps.NameToColumn("plan_date")
	se := brps.Series[col]
	se.Sort(ctx, dataframe.SortOptions{Desc: false})

	n, err := getLatestNilActualRowNum(brps)
	nrow := se.NRows()
	check(err)
	startDate := se.Value(n).(civil.Date)
	if n > 0 {
		startDate = se.Value(n - 1).(civil.Date)
	}

	endDate := se.Value(nrow - 1).(civil.Date)
	planDates := genMonthlyInsPlanDate(startDate, endDate)
	planDateSlice := fliterDates(planDates, se)
	// 填充生成的planDate，并对其他字段进行填充
	model.Brps = brps
	maps := model.slice2maps("plan_date", planDateSlice...)

	// 组装dataframe
	// 注意maps中字段要与series一一对应，否则报错"no. of args not equal to no. of series"
	for _, val := range maps {
		brps.Append(nil, val)
	}

	model.Brps = brps

	return model
}

// FillPlanDateSeasonly 对还未还款的记录（actual_date 为nil），生成每季度21日的还息日期plan_date，并按planDate升序排序。
// 其他字段以nil进行填充 （bank_loan_contract_id 用 bc.id填充）
func (model *BankRepayPlanCalcModel) FillPlanDateSeasonly() *BankRepayPlanCalcModel {
	brps := model.Brps
	col, _ := brps.NameToColumn("plan_date")
	se := brps.Series[col]
	se.Sort(ctx, dataframe.SortOptions{Desc: false})

	n, err := getLatestNilActualRowNum(brps)
	nrow := se.NRows()
	check(err)
	startDate := se.Value(n).(civil.Date)
	if n > 0 {
		startDate = se.Value(n - 1).(civil.Date)
	}
	endDate := se.Value(nrow - 1).(civil.Date)
	planDates := genSeasonlyInsPlanDate(startDate, endDate)
	planDateSlice := fliterDates(planDates, se)
	// 填充生成的planDate，并对其他字段进行填充
	model.Brps = brps
	maps := model.slice2maps("plan_date", planDateSlice...)

	// 组装dataframe
	// 注意maps中字段要与series一一对应，否则报错"no. of args not equal to no. of series"
	for _, val := range maps {
		brps.Append(nil, val)
	}

	model.Brps = brps

	return model
}

// ConvTimeToDate 将 model 中含有_date字段 的 time转换为civil.date：time2date
func (model *BankRepayPlanCalcModel) convTimeToDate() *BankRepayPlanCalcModel {
	df := model.Brps
	for _, name := range df.Names() {
		if strings.Contains(name, "_date") {
			n, err := df.NameToColumn(name)
			check(err)
			se, err := timeSerie2dateSerie(&df.Series[n])
			check(err)
			err = df.RemoveSeries(name)
			check(err)
			err = df.AddSeries(se, nil)
			check(err)
		}
	}
	model.Brps = df
	return model
}

// AddPlanAmount 加入列“计划还款总额”
func (model *BankRepayPlanCalcModel) AddPlanAmount() *BankRepayPlanCalcModel {
	df := model.Brps
	nrows := df.NRows()
	// 对原始df的直接迭代更新，只能更新部分数据，故每次都另外生成slice，再添加到df
	planAmount := make([]interface{}, nrows)

	// 计算当前行的计划还款总额planAmount
	rowPlanAmountCalc := func(vals map[interface{}]interface{}) (planAmount int64) {
		planInterest := vals["plan_interest"].(int64)
		if vals["plan_principal"] == nil {
			planAmount = planInterest
		} else {
			planAmount = planInterest + vals["plan_principal"].(int64)
		}
		return
	}

	applyDfFn := dataframe.ApplyDataFrameFn(func(val map[interface{}]interface{}, row, nRows int) map[interface{}]interface{} {
		planAmount[row] = rowPlanAmountCalc(val)
		return val
	})

	dataframe.Apply(ctx, df, applyDfFn, dataframe.FilterOptions{InPlace: true})

	se2 := dataframe.NewSeriesInt64("plan_amount", nil, planAmount...)
	df.RemoveSeries("plan_amount")
	df.AddSeries(se2, nil)
	model.Brps = df
	return model
}

func calcDays(vals map[interface{}]interface{}, upperVals map[interface{}]interface{}) int {
	upperDay := genDay(upperVals)
	rowDay := genDay(vals)
	return (rowDay).DaysSince(upperDay)
}

func genDay(vals map[interface{}]interface{}) civil.Date {
	day := vals["plan_date"].(civil.Date)
	if vals["actual_date"] != nil {
		day = (vals["actual_date"].(civil.Date))
	}
	return day
}

// rowInsCalc 计算本行的计划利息。参数vals传入本行row值，upperVals传入上一行值
// 在传入本函数前，应该先完成sort排序 model.Sort("plan_date")， 不在本函数里做sort 是担心与其他调用函数形成死锁
// method不传参，为默认计息方式（年利率/360×天数）适用多数银行
// method传参"monthly",为按月利率计息(月利率/30×天数)，与默认计息方式的区别是利率的四舍五入,适用杭州银行
// method传参"halfyearly",按半年计息，适用招行
func (model *BankRepayPlanCalcModel) rowInsCalc(vals map[interface{}]interface{}, upperVals map[interface{}]interface{}, method ...string) (int64, error) {
	insCalcOptions := make([]InsCalcOption, 1)
	option := insCalcOptions[0]
	if method == nil {
		option.Method = "yearly"
	} else {
		option.Method = method[0]
	}
	option.ExeRate = model.Bc.CurrentRate // 默认取当前利率字段

	isLpr := model.Bc.IsLpr.Valid // 是否LPR定价
	if isLpr {
		option.LprPlus = model.Bc.LprPlus
	}
	crd := model.Bc.CurrentRepriceDay
	iscrd := crd.Valid   //重定价日本身是否有值
	isLprChange := false // 重定价日前后LPR是否发生变化
	isCrdIn := false     // 重定价日是否落在upperDay和rowDay之间
	upperDay := genDay(upperVals)
	rowDay := genDay(vals)
	if iscrd {
		d := civil.DateOf(crd.ValueOrZero())

		// 若重定价日落在upperDay之前，则执行利率按重定价日的LPR重新计算
		if d.Before(upperDay) {
			option.reprice(d)
		}

		// 重定价日前后LPR是否发生变化
		lprBefore := getLpr(d)
		lprAfter := getLpr(d.AddDays(1))
		if lprBefore != lprAfter {
			isLprChange = true
		}

		// 重定价日是否落在upperDay和rowDay之间
		if ((d.DaysSince(upperDay) == 0) || d.After(upperDay)) && d.Before(rowDay) {
			isCrdIn = true
		}

	}

	// 条件判断：在isLPR为true && 重定价日有值 && 重定价日前后LPR发生了变化 && 重定价日落在upperDay和rowDay之间 ,满足上述所有条件时，才需要分前后利率计息
	if isLpr && isLprChange && isCrdIn {
		d := civil.DateOf(crd.ValueOrZero())
		return segIns(d, vals, upperVals, option)
	} else {
		// 非LPR定价合同，暂认为就是固定利率（即认为未来人行基准将不会变化）
		// LPR定价，但重定价日在upperDay之前，也可认为计息期间内是固定利率
		return fixedIns(vals, upperVals, option)
	}
}

// segIns 分段计息
// TODO:未考虑upperDay和rowDay之间存在多个重定价日的情况，可能影响招行保理利息计算
func segIns(repriceDay civil.Date, vals map[interface{}]interface{}, upperVals map[interface{}]interface{}, option InsCalcOption) (int64, error) {
	// 做一个深拷贝
	var midVals map[interface{}]interface{}
	deepcopy.Copy(&midVals, &vals).Do()
	midVals["actual_date"] = repriceDay
	segins1, err := fixedIns(midVals, upperVals, option)
	check(err)
	segins2, err := fixedIns(vals, midVals, *option.reprice(repriceDay))
	check(err)
	return segins1 + segins2, nil
}

// fixedIns 固定利率计息。重定价日不落在vals 和upperVals之间时，适用此方法计息
func fixedIns(vals map[interface{}]interface{}, upperVals map[interface{}]interface{}, option InsCalcOption) (int64, error) {

	calcDays := calcDays(vals, upperVals)
	planInsB := big.NewInt(0)
	accruedB := big.NewInt(vals["accrued_principal"].(int64))
	yearRate := option.ExeRate
	switch option.Method {
	case "yearly":
		// 默认_计划利息 = 应计本金×年利率×期间天数/360 （因利率单位为0.01%，所以再除以1000000）
		rateB := big.NewInt(int64(yearRate))
		planInsB.Mul(accruedB, rateB).Mul(planInsB, big.NewInt(int64(calcDays)))
		planInsB.Div(planInsB, big.NewInt(360000000))
	case "monthly":
		// 月利率=年利率/12,精确到0.0001%
		rateMonthB := big.NewInt(int64(yearRate * 10 / 12)) //为了月利率的精度，小数点右移一位后再除以12
		// 计划利息 = 应计本金×月利率×期间天数/30
		planInsB = planInsB.Mul(accruedB, rateMonthB).Mul(planInsB, big.NewInt(int64(calcDays))).Div(planInsB, big.NewInt(300000000)) // 注意最后要多除以10，即把上面移动的小数点移回去
	case "halfyearly":
		// 半利率=年利率/2,精确到0.0001%
		rateHalfB := big.NewInt(int64(yearRate / 2))
		// 计划利息 = 应计本金×半利率
		planInsB.Mul(accruedB, rateHalfB).Div(planInsB, big.NewInt(1000000)) //因利率单位为0.01%，所以再除以1000000

	default:
		return -1, fmt.Errorf("未定义的计息方式")
	}

	// 四舍五入
	p := planInsB.Int64()
	return rounding(p), nil
}

func timeSerie2dateSerie(d *dataframe.Series) (*dataframe.SeriesGeneric, error) {

	typ := (*d).Type()
	if typ != "time" {
		fmt.Println(typ)
		return nil, fmt.Errorf("格式错误")
	}
	t := (*d).Copy()
	colname := (t).Name()
	// vals := []civil.Date{}
	x := (t).NRows()
	vals := make([]interface{}, x)
	// 将time转换为date并放到另一个slice中 （直接in place替换，可能会有更新不全的问题）
	fconvert := dataframe.ApplySeriesFn(func(val interface{}, row, nRows int) interface{} {
		if val == nil {
			vals[row] = nil
		} else {
			z := val.(time.Time)
			vals[row] = civil.DateOf(z)
		}
		return val
	})
	dataframe.Apply(ctx, t, fconvert, dataframe.FilterOptions{InPlace: true})

	// 生成dateSerie
	se := dataframe.NewSeriesGeneric(colname, civil.Date{}, nil, vals...)

	// 定义比较函数，以便用于排序
	f := func(a, b interface{}) bool {
		a1 := a.(civil.Date)
		b1 := b.(civil.Date)
		return a1.Before(b1)
	}
	se.SetIsLessThanFunc(f)

	return se, nil

}

// TODO:lpr中段变化时，分段计息的计算公式

// slice2maps 传入两个参数，生成map[fieldname]val . 其中 "bank_loan_contract_id" 字段固定为 model.Bc.ID; fieldname 字段值为val
// 该map用于添加至 model.Brps 的dataframe中
func (model *BankRepayPlanCalcModel) slice2maps(fieldname string, vals ...interface{}) []interface{} {
	brps := model.Brps
	names := brps.Names()
	bcID := model.Bc.ID
	maps := make([]interface{}, 0)
	if vals == nil {
		panic("get no vals")
	} else {
		for _, val := range vals {
			// 初始化a
			a := make(map[string]interface{})
			for _, name := range names {
				a[name] = nil
			}

			a["bank_loan_contract_id"] = int64(bcID)
			a[fieldname] = val

			maps = append(maps, a)
		}
	}
	// fmt.Printf("maps:\n %+v", maps)
	return maps

}
