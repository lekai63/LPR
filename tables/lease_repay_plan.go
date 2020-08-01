package tables

import (
	"strconv"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/lekai63/lpr/models"
)

func GetLeaseRepayPlanTable(ctx *context.Context) table.Table {

	leaseRepayPlan := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := leaseRepayPlan.GetInfo()

	info.AddField("Id", "id", db.Int).FieldHide()
	info.AddField("项目简称", "abbreviation", db.Varchar).FieldJoin(types.Join{
		Table:     "lease_contract",    // 连表的表名，对应的外面的表的表名
		Field:     "lease_contract_id", // 要连表的字段,即本表的字段
		JoinField: "id",                // 连表的表的字段，对应的外面的表字段
	})

	info.AddField("期次", "period", db.Int2).FieldDisplay(func(model types.FieldModel) interface{} {
		return "第" + model.Value + "期"
	})
	info.AddField("计划还款日期", "plan_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("计划还款租金合计", "plan_amount", db.Int8).FieldDisplay(showMoney)
	info.AddField("计划还款本金", "plan_principal", db.Int8).FieldDisplay(showMoney)
	info.AddField("计划还款利息", "plan_interest", db.Int8).FieldDisplay(showMoney)
	info.AddField("实际还款日期", "actual_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("实际还款租金合计", "actual_amount", db.Int8).
		FieldDisplay(func(model types.FieldModel) interface{} {
			row := model.Row
			if row["actual_amount"] != nil {
				return showMoney(model)
			} else {
				return ""
			}
		})
	info.AddField("实际还款本金", "actual_principal", db.Int8).
		FieldDisplay(func(model types.FieldModel) interface{} {
			row := model.Row
			if row["actual_principal"] != nil {
				return showMoney(model)
			} else {
				return ""
			}
		})
	info.AddField("实际还款利息", "actual_interest", db.Int8).FieldDisplay(showMoney).
		FieldDisplay(func(model types.FieldModel) interface{} {
			row := model.Row
			if row["actual_interest"] != nil {
				return showMoney(model)
			} else {
				return ""
			}
		})
	info.AddField("Created_at", "created_at", db.Timestamp).FieldHide()
	info.AddField("Updated_at", "updated_at", db.Timestamp).FieldHide()

	info.HideNewButton()
	info.SetTable("fzzl.lease_repay_plan").SetTitle("租金台账").SetDescription("LeaseRepayPlan")

	formList := leaseRepayPlan.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).FieldHide().FieldNotAllowEdit().FieldNotAllowAdd()

	dbGorm := models.Gorm
	var leaseContractGorm models.LeaseContract

	formList.AddField("项目简称", "lease_contract_id", db.Int, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().
		FieldDisplay(func(model types.FieldModel) interface{} {
			id, _ := strconv.Atoi(model.Value)
			dbGorm.First(&leaseContractGorm, id)
			a := leaseContractGorm.Abbreviation
			if a.Valid {
				return a.String
			} else {
				return ""
			}
		})

	formList.AddField("期次", "period", db.Int2, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().
		FieldDisplay(func(model types.FieldModel) interface{} {
			return "第" + model.Value + "期"
		})
	formList.AddField("计划还款日期", "plan_date", db.Date, form.Date).FieldNotAllowAdd().FieldNotAllowEdit().FieldSubstr(0, 10)
	formList.AddField("计划还款租金合计", "plan_amount", db.Int8, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().FieldDisplay(showMoney)
	formList.AddField("计划还款本金", "plan_principal", db.Int8, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().FieldDisplay(showMoney)
	formList.AddField("计划还款利息", "plan_interest", db.Int8, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().FieldDisplay(showMoney)
	formList.AddField("实际还款日期", "actual_date", db.Date, form.Date)
	formList.AddField("实际还款金额", "actual_amount", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(money2bigint)
	formList.AddField("实际还款本金", "actual_principal", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(money2bigint)
	formList.AddField("实际还款利息", "actual_interest", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(money2bigint)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).FieldHide().FieldNotAllowEdit().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			if value.Value == nil {
				return time.Now().Format("2006-01-02 15:04:05")
			}
			return value.Value.Value()
		})
	formList.AddField("修改时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			return time.Now().Format("2006-01-02 15:04:05")
		})

	formList.SetTable("fzzl.lease_repay_plan").SetTitle("LeaseRepayPlan").SetDescription("LeaseRepayPlan")

	return leaseRepayPlan
}
