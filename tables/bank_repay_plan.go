package tables

import (
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/lekai63/lpr/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func GetBankRepayPlanTable(ctx *context.Context) table.Table {

	bankRepayPlan := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := bankRepayPlan.GetInfo().HideFilterArea()

	info.AddField("序号", "id", db.Int)

	info.AddField("项目简称", "abbreviation", db.Varchar).FieldJoin(types.Join{
		Table:     "bank_loan_contract",    // 连表的表名，对应的外面的表的表名
		Field:     "bank_loan_contract_id", // 要连表的字段,即本表的字段
		JoinField: "id",                    // 连表的表的字段，对应的外面的表字段
	})
	info.AddField("计划还款日期", "plan_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("计划还款本息合计", "plan_amount", db.Int8).FieldDisplay(showMoney)
	info.AddField("计划还款本金", "plan_principal", db.Int8).FieldDisplay(showMoney)
	info.AddField("计划还款利息", "plan_interest", db.Int8).FieldDisplay(showMoney)
	info.AddField("实际还款日期", "actual_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("实际还款本息合计", "actual_amount", db.Int8).
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
	info.AddField("实际还款利息", "actual_interest", db.Int8).
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
	info.HideNewButton() //隐藏新增按钮
	info.SetTable("fzzl.bank_repay_plan").SetTitle("BankRepayPlan").SetDescription("BankRepayPlan")

	dbGorm := models.Gormv2
	var BankLoanContractGorm models.BankLoanContract
	formList := bankRepayPlan.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).FieldHide().FieldNotAllowEdit().FieldNotAllowAdd()

	formList.AddField("项目简称", "bank_loan_contract_id", db.Int, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().
		FieldDisplay(func(model types.FieldModel) interface{} {
			// 获取第一条匹配的记录
			id := model.Value
			dbGorm.Where("id = ?", id).First(&BankLoanContractGorm)
			a := BankLoanContractGorm.Abbreviation
			if a.Valid {
				return a.String
			} else {
				return ""
			}
		})
	formList.AddField("计划还款日期", "plan_date", db.Date, form.Date).FieldNotAllowAdd().FieldNotAllowEdit().FieldSubstr(0, 10)
	formList.AddField("计划还款本息合计", "plan_amount", db.Int8, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().FieldDisplay(showMoney)
	formList.AddField("计划还款本金", "plan_principal", db.Int8, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().FieldDisplay(showMoney)
	formList.AddField("计划还款利息", "plan_interest", db.Int8, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().FieldDisplay(showMoney)
	formList.AddField("实际还款日期", "actual_date", db.Date, form.Date)
	formList.AddField("实际还款本息合计", "actual_amount", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元")
	formList.AddField("实际还款本金", "actual_principal", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元")
	formList.AddField("实际还款利息", "actual_interest", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元")
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).FieldHide().FieldNotAllowEdit()
	formList.AddField("修改时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide()

	formList.SetTable("fzzl.bank_repay_plan").SetTitle("BankRepayPlan").SetDescription("BankRepayPlan")

	formList.SetUpdateFn(func(values form2.Values) (err error) {
		var bankRepayPlanGorm models.BankRepayPlan
		dbGorm.Where("id= ?", values.Get("id")).First(&bankRepayPlanGorm)

		formatDate := "2006-01-02"
		tempDate, _ := time.Parse(formatDate, values.Get("actual_date"))
		bankRepayPlanGorm.ActualDate.SetValid(tempDate)
		bankRepayPlanGorm.ActualAmount.SetValid(cast.ToInt64(floatStr2BigintStr(values.Get("actual_amount"))))
		bankRepayPlanGorm.ActualPrincipal.SetValid(cast.ToInt64(floatStr2BigintStr(values.Get("actual_principal"))))
		bankRepayPlanGorm.ActualInterest.SetValid(cast.ToInt64(floatStr2BigintStr(values.Get("actual_interest"))))

		err = dbGorm.Transaction(func(tx *gorm.DB) error {
			if e := tx.Model(&bankRepayPlanGorm).
				Select("ActualDate", "ActualAmount", "ActualPrincipal", "ActualInterest").
				Updates(bankRepayPlanGorm).Error; e != nil {
				return e
			}
			// result 存储中间步骤，结构体内需为大写，否则后面Scan赋值失败。
			type result struct {
				Blcid          int32
				Totalprincipal int64
				Totalinterest  int64
			}
			var r result

			tx.Session(&gorm.Session{WithConditions: false}).Debug().
				Table("bank_repay_plan").
				Select("bank_loan_contract_id as blcid, sum(actual_principal) as totalprincipal ,sum(actual_interest) as totalinterest ").
				Where("bank_loan_contract_id = ?", bankRepayPlanGorm.BankLoanContractID).
				Group("bank_loan_contract_id").
				Scan(&r)

			fmt.Printf("\n result is:\n  %+v ", r)

			//更新BankLoanContract已还本金、已还利息
			var blc models.BankLoanContract
			blc.ID = r.Blcid
			blc.AllRepaidPrincipal.SetValid(r.Totalprincipal)
			blc.AllRepaidInterest.SetValid(r.Totalinterest)

			fmt.Printf("\n blc is:\n %+v ", blc)
			if e := tx.Session(&gorm.Session{WithConditions: false}).
				Model(&blc).
				Select("AllRepaidPrincipal", "AllRepaidInterest").
				Updates(blc).Error; e != nil {
				return e
			}

			return nil

		})

		return
	})

	formList.SetInsertFn(func(values form2.Values) error {
		// return nil
		return fmt.Errorf("请直接通过《导入银行还款台账》导入数据")
	})

	return bankRepayPlan
}
