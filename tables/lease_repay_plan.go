package tables

import (
	// "fmt"
	"fmt"
	"time"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/cast"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
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

	dbGorm := models.Gormv2
	var leaseContractGorm models.LeaseContract
	formList := leaseRepayPlan.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldHide().FieldNotAllowEdit().FieldNotAllowAdd()

	formList.AddField("项目简称", "lease_contract_id", db.Int, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().
		FieldDisplay(func(model types.FieldModel) interface{} {
			// 获取第一条匹配的记录
			id := model.Value
			dbGorm.Where("id = ?", id).First(&leaseContractGorm)
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

	formList.SetUpdateFn(func(values form2.Values) (err error) {

		var leaseRepayPlanGorm models.LeaseRepayPlan

		dbGorm.Where("id= ?", values.Get("id")).First(&leaseRepayPlanGorm)

		formatDate := "2006-01-02"
		tempDate, _ := time.Parse(formatDate, values.Get("actual_date"))
		leaseRepayPlanGorm.ActualDate.SetValid(tempDate)
		leaseRepayPlanGorm.ActualAmount.SetValid(cast.ToInt64(floatStr2BigintStr(values.Get("actual_amount"))))
		leaseRepayPlanGorm.ActualPrincipal.SetValid(cast.ToInt64(floatStr2BigintStr(values.Get("actual_principal"))))
		leaseRepayPlanGorm.ActualInterest.SetValid(cast.ToInt64(floatStr2BigintStr(values.Get("actual_interest"))))

		err = dbGorm.Transaction(func(tx *gorm.DB) error {

			if e := tx.Model(&leaseRepayPlanGorm).
				// Gorm的Select字段 对应 Struct字段名，而非数据库字段名。
				Select("ActualDate", "ActualAmount", "ActualPrincipal", "ActualInterest").
				Updates(leaseRepayPlanGorm).Error; e != nil {
				return e
			}

			// result 存储中间步骤，结构体内需为大写，否则后面Scan赋值失败。
			type result struct {
				Lcid           int32
				Totalprincipal int64
				Totalinterest  int64
			}

			var r result

			tx.Session(&gorm.Session{WithConditions: false}).
				Table("lease_repay_plan").
				Select("lease_contract_id as lcid  , sum(actual_principal) as totalprincipal ,sum(actual_interest) as totalinterest ").
				Where("lease_contract_id = ?", leaseRepayPlanGorm.LeaseContractID).
				Group("lease_contract_id").
				Scan(&r)

			// fmt.Printf("%+v", r)

			// 更新LeaseContract已收本金、已收利息
			var lc models.LeaseContract
			lc.ID = r.Lcid
			lc.ReceivedInterest.SetValid(r.Totalinterest)
			lc.ReceivedPrincipal.SetValid(r.Totalprincipal)
			// fmt.Printf("%+v", lc)
			if e := tx.Session(&gorm.Session{WithConditions: false}).
				Model(&lc).
				Select("ReceivedPrincipal", "ReceivedInterest").
				Updates(lc).Error; e != nil {
				return e
			}

			return nil

		})

		return nil

	})

	formList.SetInsertFn(func(values form2.Values) error {
		// return nil
		return fmt.Errorf("请直接通过《导入租金台账》导入数据")
	})

	formList.SetTable("fzzl.lease_repay_plan").SetTitle("LeaseRepayPlan").SetDescription("LeaseRepayPlan")

	return leaseRepayPlan
}
