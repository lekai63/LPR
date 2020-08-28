package lpr

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lekai63/lpr/models"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

func Test() {
	var brp models.BankRepayPlan
	df, _ := InitDataframe(brp)
	fmt.Print(df.Table())
}

// Gen dataframe from BankRepayPlan Struct
func InitDataframe(brp models.BankRepayPlan) (df *dataframe.DataFrame, err error) {
	typ := reflect.TypeOf(brp)
	val := reflect.ValueOf(brp)
	// sliceInt64 := make([]dataframe.SeriesInt64, 0)
	// sliceTime := make([]dataframe.SeriesTime, 0)
	for i := 0; i < val.NumField(); i++ {
		// name:=typ.Field(i).Name
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
		default:
			err = fmt.Errorf("无法识别的gorm type")
			return
		}

	}

	// df = dataframe.NewDataFrame(&sliceInt64[0], &sliceInt64[1], &sliceTime[1])

	return

}

// getColnameWithType 返回gormTag 对应的colname,type(只判断int,date)
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
			if strings.HasPrefix(temp[1], "INT") {
				typename = "int"
				continue
			} else if strings.HasPrefix(temp[1], "DATE") || strings.HasPrefix(temp[1], "TIMESTAMP") {
				typename = "time"
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
