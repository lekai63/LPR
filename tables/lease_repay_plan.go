package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	"github.com/lekai63/lpr/models"
)

func GetLeaseRepayPlanTable(ctx *context.Context) table.Table {

	leaseRepayPlan := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := leaseRepayPlan.GetInfo().HideFilterArea()

	dbGorm := models.Gorm
	var lesseeContractGorm models.LeaseContract

	info.AddField("Id", "id", db.Int).FieldHide()
	info.AddField("项目简称", "lease_contract_id", db.Int).FieldDisplay(func(model types.FieldModel) interface{} {
		dbGorm.First(&lesseeContractGorm, "id = $1", model.Value)
		return lesseeContractGorm.Abbreviation
	})
	info.AddField("Period", "period", db.Int2)
	info.AddField("Plan_date", "plan_date", db.Date)
	info.AddField("Plan_amount", "plan_amount", db.Int8)
	info.AddField("Plan_principal", "plan_principal", db.Int8)
	info.AddField("Plan_interest", "plan_interest", db.Int8)
	info.AddField("Actual_date", "actual_date", db.Date)
	info.AddField("Actual_amount", "actual_amount", db.Int8)
	info.AddField("Actual_principal", "actual_principal", db.Int8)
	info.AddField("Actual_interest", "actual_interest", db.Int8)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("fzzl.lease_repay_plan").SetTitle("LeaseRepayPlan").SetDescription("LeaseRepayPlan")

	formList := leaseRepayPlan.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default)
	formList.AddField("Lease_contract_id", "lease_contract_id", db.Int4, form.Number)
	formList.AddField("Period", "period", db.Int2, form.Text)
	formList.AddField("Plan_date", "plan_date", db.Date, form.Datetime)
	formList.AddField("Plan_amount", "plan_amount", db.Int8, form.Text)
	formList.AddField("Plan_principal", "plan_principal", db.Int8, form.Text)
	formList.AddField("Plan_interest", "plan_interest", db.Int8, form.Text)
	formList.AddField("Actual_date", "actual_date", db.Date, form.Datetime)
	formList.AddField("Actual_amount", "actual_amount", db.Int8, form.Text)
	formList.AddField("Actual_principal", "actual_principal", db.Int8, form.Text)
	formList.AddField("Actual_interest", "actual_interest", db.Int8, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.lease_repay_plan").SetTitle("LeaseRepayPlan").SetDescription("LeaseRepayPlan")

	return leaseRepayPlan
}
