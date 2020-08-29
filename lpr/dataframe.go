package lpr

import (
	"fmt"
	"reflect"
	"strings"

	"cloud.google.com/go/civil"
	"github.com/lekai63/lpr/models"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

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
		default:
			err = fmt.Errorf("无法识别:%s", typename)
			return
		}

	}

	return

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
