package inscalc

import (
	"errors"

	"cloud.google.com/go/civil"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

// ToCMB 生成招商银行还款计划
func (model *BankRepayPlanCalcModel) ToCMB(isFirst bool) (*BankRepayPlanCalcModel, error) {
	if model.Bc.BankName != "招商银行" {
		return nil, errors.New("输入模型的银行名称不是招商银行，请检查")
	}

	// 招行每期利息与本金一起支付，不需额外生成利息还款计划
	// if isFirst {
	// 	model.FillInsPlanDate()
	// }

	model.AddAccruedPrincipal()
	model.AddCMBFactoringIns()
	model.AddPlanAmount()
	return model, nil
}

// AddCMBFactoringIns 计算招商银行保理利息并添加到列，本函数将df.lock 注意避免与其他函数形成死锁
// 招商银行保理利息方案为：本金利息均根据应收款偿还计划在6月、12月的20日偿还。还本付息时若遇节假日，仍按之前计算好的本息还款，不额外加收多占天数的利息。
// 故在使用rowInsCalc函数前，需调整传入的参数vals和upperVals
func (model *BankRepayPlanCalcModel) AddCMBFactoringIns() *BankRepayPlanCalcModel {
	model.Sort("plan_date")
	df := model.Brps
	// upperVal 用于存储上一个row的信息
	upperVals := make(map[interface{}]interface{})

	nrows := df.NRows()
	// planInsterest Slice用于存储计算后的计划利息,因最后要将 planInsterest 传入NewSeries，故类型直接选择[]interface{} ,如定义为[]int,还需再循环转换为[]interface{}
	// https://golang.org/doc/faq#convert_slice_of_interface
	planInsterest := make([]interface{}, nrows)
	var e error
	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
	df.Lock()
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		switch {
		// 第0行为第一笔利息还款计划
		case (*row) == 0:
			upperVals["plan_date"] = civil.DateOf(model.Bc.ActualStartDate.ValueOrZero())
			upperVals["actual_date"] = nil
			vals["actual_date"] = nil
			planInsterest[*row], e = model.rowInsCalc(vals, upperVals)
			check(e)
		default:
			// 不需要根据实际还款日期重新计算当期利息，故传参前调整vals，upperVals
			vals["actual_date"] = nil
			upperVals["actual_date"] = nil
			planInsterest[*row], e = model.rowInsCalc(vals, upperVals, "halfyearly")
			check(e)
		}
		// 本次循环结束时，将本行赋值给upperVal用于下次循环
		upperVals = vals
	}
	df.Unlock()
	// 移除旧的plan_interest 以及不需要用的几个字段
	df.RemoveSeries("plan_interest")
	df.RemoveSeries("created_at")
	df.RemoveSeries("updated_at")
	// df.RemoveSeries("accrued_principal")

	//添加新的plan_interest
	se := dataframe.NewSeriesInt64("plan_interest", nil, planInsterest...)
	df.AddSeries(se, nil)

	model.Brps = df
	return model
}
