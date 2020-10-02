package inscalc

import (
	"errors"

	"cloud.google.com/go/civil"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

// ToDefault 根据model生成默认的还本付息计划：利息在每月或每季度21日偿还，本金在其他日期偿还，且本金偿还时不额外付息
// 适用工行、农行
// 首次生成，则isFirst 为true,生成后写入数据库；非首次生成，则预期数据库中已有InsPlanDate信息
func (model *BankRepayPlanCalcModel) ToDefault(isFirst bool) (*BankRepayPlanCalcModel, error) {

	// 注意使用括号决定计算优先级，不要直接链式调用
	if isFirst {
		model.FillInsPlanDate()
	}
	model.AddAccruedPrincipal()
	model.AddDefaultFactoringIns()
	model.AddPlanAmount()
	return model, nil
}

// AddDefaultFactoringIns 计算默认保理利息并添加到列，本函数将df.lock 注意避免与其他函数形成死锁
// 默认保理利息方案为：利息在每月或每季度21日偿还，本金在其他日期偿还，且本金偿还时不额外付息
func (model *BankRepayPlanCalcModel) AddDefaultFactoringIns() *BankRepayPlanCalcModel {
	model.Sort("plan_date")
	df := model.Brps
	// upperVal 用于存储上一个row的信息
	upperVals := make(map[interface{}]interface{})
	var temp int64
	nrows := df.NRows()
	// planInsterest Slice用于存储计算后的计划利息,因最后要将 planInsterest 传入NewSeries，故类型直接选择[]interface{} ,如定义为[]int,还需再循环转换为[]interface{}
	// https://golang.org/doc/faq#convert_slice_of_interface
	planInsterest := make([]interface{}, nrows)

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
			planInsterest[*row] = model.rowInsCalc(vals, upperVals)

			// 最后一行利随本清
		case (*row) == nrows-1: // 如使用(*row) == df.NRows() 游标直接到最后，从而无法执行
			planInsterest[*row] = model.rowInsCalc(vals, upperVals)

		default:
			// 因农行保理利息在每季末21日扣，故本金还款日当天的利息，应加到下一季末的21日一并扣息
			// 本row还款日非21日，将计划还款利息暂存入temp
			if d := vals["plan_date"].(civil.Date); d.Day != 21 {
				temp = model.rowInsCalc(vals, upperVals)
				planInsterest[*row] = 0
			} else if x := upperVals["plan_date"].(civil.Date); x.Day != 21 {
				// 上一row为非21日，将temp提取出来，加入本row
				planInsterest[*row] = model.rowInsCalc(vals, upperVals) + temp
			} else {
				// 默认planInsterest算法
				planInsterest[*row] = model.rowInsCalc(vals, upperVals)
			}
		}

		// 存在数据竞争，如直接更新row，最后的df数据不全，故删去，转而在前面数据处理时，组装一个series，再将series整体加入到df
		// df.UpdateRow(*row, &dataframe.DontLock, vals)

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

func (model *BankRepayPlanCalcModel) ToABC(isFirst bool) (*BankRepayPlanCalcModel, error) {
	if model.Bc.BankName != "农业银行" {
		return nil, errors.New("输入模型的银行名称不是农业银行，请检查")
	}
	model.ToDefault(isFirst)
	return model, nil

}

func (model *BankRepayPlanCalcModel) ToCCB(isFirst bool) (*BankRepayPlanCalcModel, error) {
	if model.Bc.BankName != "建设银行" {
		return nil, errors.New("输入模型的银行名称不是建设银行，请检查")
	}
	model.ToDefault(isFirst)
	return model, nil

}

// ToICBC 根据model 生成工行还本付息计划
// 首次生成，则isFirst 为true；
func (model *BankRepayPlanCalcModel) ToICBC(isFirst bool) (*BankRepayPlanCalcModel, error) {
	if model.Bc.BankName != "工商银行" {
		return nil, errors.New("输入模型的银行名称不是工商银行，请检查")
	}
	model.ToDefault(isFirst)
	return model, nil
}
