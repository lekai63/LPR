package inscalc

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"cloud.google.com/go/civil"
	"github.com/guregu/null"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

// TODO:修改db获取方式
var ctx = context.Background()
var db, _ = gormInitForTest()

// Option 利息计算的参数
type Option struct {
	// 计息方式:
	// yearly: 年利率/360计息
	// monthly: 月利率/30 计息
	// halfyearly: 年利率/2 计息半年
	Method  string
	ExeRate int32 //执行利率
	LprPlus null.Int
}

// LprRecord 表"lpr_record"的结构体模型
type LprRecord struct {
	ID int32 `gorm:"primaryKey;autoIncrement;column:id;type:INT4;"`
	// json: cannot unmarshal string into Go struct field Record.records.1Y of type float64
	OneY string `json:"1Y" gorm:"column:one_y;type:FLOAT8"`
	// ShowDateEN time.Time `json:"showDateEN,omitempty"`
	FiveY      string `json:"5Y" gorm:"column:five_y;type:FLOAT8"`
	ShowDateCN string `json:"showDateCN" gorm:"column:show_date;type:DATE"`
}

// TableName sets the insert table name for this struct type
func (b *LprRecord) TableName() string {
	return "lpr_record"
}

// reprice 输入day日，输出重定价后的利率=day日执行的LPR+LprPlus
// 如day日当天公布LPR，该LPR会在day+1日执行。故day日当天适用此前的LPR
func (option *Option) reprice(day civil.Date) *Option {
	var r []LprRecord
	// 取一年期LPR. LPR公布当日执行的是上一日的LPR，故查询语句只需要写"<" 而非"<="
	db.Where("show_date < ? ", day.String()).Order("show_date desc").Find(&r)
	lpr := floatStr2int64(r[0].OneY)
	option.ExeRate = int32(lpr) + int32(option.LprPlus.ValueOrZero())
	return option

}

// getLpr 获得 day日之前的lpr, 数值乘以10000 后返回int64（即单位为0.0001%） ，以便与insCalc利率单位一致
// 注意：T日公布的LPR，将在T+1日生效。故day日事实上执行的是day日之前的LPR
func getLpr(day civil.Date) int64 {
	dpast := day.AddDays(-31)
	dyestoday := day.AddDays(-1) //构造查询语句时，不包含今日，以免混入本应明日生效的LPR
	var records []LprRecord
	result := db.Where("show_date BETWEEN ? AND ?", dpast.String(), dyestoday.String()).Find(&records)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	switch result.RowsAffected {
	case 0:
		fmt.Printf("未查询到距离%s(含)之前30天内的LPR值,取最新一期LPR", day.String())
		var r LprRecord
		restemp := db.Order("show_date desc").First(&r)
		if restemp.Error != nil {
			log.Fatalln(restemp.Error)
		}
		return floatStr2int64(r.OneY)
	case 1:
		return floatStr2int64(records[0].OneY)
	default:
		//取日期最近的一期记录
		m := records[0]
		for _, val := range records {
			valday, err := civil.ParseDate(val.ShowDateCN[:10])
			if err != nil {
				log.Fatalln(err)
			}
			mday, err := civil.ParseDate(m.ShowDateCN[:10])
			if err != nil {
				log.Fatalln(err)
			}

			if valday.After(mday) {
				m = val
			}
		}
		return floatStr2int64(m.OneY)
	}
}

// floatStr2int64 浮点数×10000后取整
func floatStr2int64(floatStr string) int64 {
	f, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int64(f * 10000)
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

// genMonthlyInsPlanDate 生成利息还款计划，默认每月21日扣息.
// 若起始日为21日以后（含21日当日），则生成的第一个扣息日为下月21日
// 若起始日为21日以前，则生成的第一个扣息日为本月21日
// 最后一期利随本清,不额外生成一个利息还款计划
func genMonthlyInsPlanDate(start, end civil.Date) (dates []civil.Date) {
	termMonths := 12*(end.Year-start.Year) + (int(end.Month) - int(start.Month))
	plus := 0
	if start.Day >= 21 {
		plus = 1
	}

	for i := 0; i < termMonths; i++ {
		mInt := (int(start.Month) + i + plus) % 12
		var m time.Month
		if mInt == 0 {
			m = time.December
		} else {
			m = time.Month(mInt)
		}
		d := civil.Date{
			Year:  start.Year + (i+int(start.Month))/12,
			Month: m,
			Day:   21,
		}

		if d.Before(end) {
			dates = append(dates, d)
		}
	}
	// 不额外生成最后一期利随本清还息日期
	// dates = append(dates, end)

	return
}

// genSeasonlyInsPlanDate 生成利息还款计划，默认每季度（3、6、9、12月）21日扣息.
// 第一个扣息日为距离start_date 最近的一个季度末月的21日
// 最后一期利随本清,不额外生成一个利息还款计划
func genSeasonlyInsPlanDate(start, end civil.Date) (dates []civil.Date) {
	//d1 为第一个扣息日
	var d1 civil.Date
	if start.Month == time.December && start.Day >= 21 {
		// start 日期处于12月21日至12月31日时
		d1 = civil.Date{
			Year:  start.Year + 1,
			Month: time.March,
			Day:   21,
		}
	} else {
		for i := time.March; i <= time.December; i = i + 3 {
			m := civil.Date{
				Year:  start.Year,
				Month: i,
				Day:   21,
			}
			if start.Before(m) {
				d1 = m
				break
			}
		}

	}

	termMonths := 12*(end.Year-start.Year) + (int(end.Month) - int(start.Month))
	for i := 0; i < termMonths; i = i + 3 {
		mInt := (int(d1.Month) + i) % 12
		var m time.Month
		var d civil.Date
		if mInt == 0 {
			d = civil.Date{
				Year:  d1.Year + (i+int(d1.Month))/12 - 1,
				Month: time.December,
				Day:   21,
			}
		} else {
			m = time.Month(mInt)
			d = civil.Date{
				Year:  d1.Year + (i+int(d1.Month))/12,
				Month: m,
				Day:   21,
			}
		}

		if d.Before(end) {
			dates = append(dates, d)
		}
	}
	// 不额外生成最后一期利随本清还息日期
	// dates = append(dates, end)

	return
}

// fliterDates 筛选planDates，若planDates中的元素日期与series中的任一元素日期相同，则删除planDates中的此元素。
func fliterDates(planDates []civil.Date, se dataframe.Series) []interface{} {
	iterator := se.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
	se.Lock()
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		n := len(planDates)
		for i := 0; i < n; i++ {
			x := planDates[i]
			if (vals.(civil.Date)).DaysSince(x) == 0 {
				planDates = append(planDates[:i], planDates[i+1:]...)
				i = i - 1
				n = n - 1
			}
		}
	}
	se.Unlock()

	// convert a []T to an []interface{}
	// https://golang.org/doc/faq#convert_slice_of_interface
	s := make([]interface{}, len(planDates))
	for i, v := range planDates {
		s[i] = v
	}

	return s

}

// 四舍五入
func rounding(p int64) (res int64) {
	pString := strconv.FormatInt(p, 10)
	if len(pString) < 2 {
		res = 0
		return
	}
	if d, _ := strconv.Atoi(pString[len(pString)-2 : len(pString)-1]); d < 5 {
		pString = pString[:len(pString)-2] + "00"
		res, _ = strconv.ParseInt(pString, 10, 64)
	} else {
		res, _ = strconv.ParseInt(pString[:len(pString)-2], 10, 64)
		res = (res + 1) * 100
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
