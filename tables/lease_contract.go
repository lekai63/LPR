package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetLeaseContractTable(ctx *context.Context) table.Table {

	leaseContract := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := leaseContract.GetInfo().HideFilterArea()

	info.AddField("Cid", "cid", db.Int4)
	info.AddField("Contract_no", "contract_no", db.Varchar)
	info.AddField("Lessee", "lessee", db.Varchar)
	info.AddField("Abbreviation", "abbreviation", db.Varchar)
	info.AddField("Start_date", "start_date", db.Date)
	info.AddField("End_date", "end_date", db.Date)
	info.AddField("Fee", "fee", db.Int8)
	info.AddField("Margin", "margin", db.Int8)
	info.AddField("Contract_principal", "contract_principal", db.Int8)
	info.AddField("Actual_principal", "actual_principal", db.Int8)
	info.AddField("Term_month", "term_month", db.Int2)
	info.AddField("Subject_matter", "subject_matter", db.Varchar)
	info.AddField("Irr", "irr", db.Int4)
	info.AddField("Is_lpr", "is_lpr", db.Bool)
	info.AddField("Current_reprice_day", "current_reprice_day", db.Date)
	info.AddField("Current_LPR", "current_LPR", db.Int4)
	info.AddField("Lpr_plus", "lpr_plus", db.Int4)
	info.AddField("Current_rate", "current_rate", db.Int4)
	info.AddField("Next_reprice_day", "next_reprice_day", db.Date)
	info.AddField("Received_principal", "received_principal", db.Int8)
	info.AddField("Received_interest", "received_interest", db.Int8)
	info.AddField("Is_finished", "is_finished", db.Bool)
	info.AddField("Customer_id", "customer_id", db.Int4)
	info.AddField("Create_time", "create_time", db.Timestamp)
	info.AddField("Modify_time", "modify_time", db.Timestamp)

	info.SetTable("fzzl.lease_contract").SetTitle("LeaseContract").SetDescription("LeaseContract")

	formList := leaseContract.GetForm()
	formList.AddField("Cid", "cid", db.Int4, form.Number)
	formList.AddField("Contract_no", "contract_no", db.Varchar, form.Text)
	formList.AddField("Lessee", "lessee", db.Varchar, form.Text)
	formList.AddField("Abbreviation", "abbreviation", db.Varchar, form.Text)
	formList.AddField("Start_date", "start_date", db.Date, form.Datetime)
	formList.AddField("End_date", "end_date", db.Date, form.Datetime)
	formList.AddField("Fee", "fee", db.Int8, form.Text)
	formList.AddField("Margin", "margin", db.Int8, form.Text)
	formList.AddField("Contract_principal", "contract_principal", db.Int8, form.Text)
	formList.AddField("Actual_principal", "actual_principal", db.Int8, form.Text)
	formList.AddField("Term_month", "term_month", db.Int2, form.Text)
	formList.AddField("Subject_matter", "subject_matter", db.Varchar, form.Text)
	formList.AddField("Irr", "irr", db.Int4, form.Number)
	formList.AddField("Is_lpr", "is_lpr", db.Bool, form.Text)
	formList.AddField("Current_reprice_day", "current_reprice_day", db.Date, form.Datetime)
	formList.AddField("Current_LPR", "current_LPR", db.Int4, form.Number)
	formList.AddField("Lpr_plus", "lpr_plus", db.Int4, form.Number)
	formList.AddField("Current_rate", "current_rate", db.Int4, form.Number)
	formList.AddField("Next_reprice_day", "next_reprice_day", db.Date, form.Datetime)
	formList.AddField("Received_principal", "received_principal", db.Int8, form.Text)
	formList.AddField("Received_interest", "received_interest", db.Int8, form.Text)
	formList.AddField("Is_finished", "is_finished", db.Bool, form.Text)
	formList.AddField("Customer_id", "customer_id", db.Int4, form.Number)
	formList.AddField("Create_time", "create_time", db.Timestamp, form.Datetime)
	formList.AddField("Modify_time", "modify_time", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.lease_contract").SetTitle("LeaseContract").SetDescription("LeaseContract")

	return leaseContract
}
