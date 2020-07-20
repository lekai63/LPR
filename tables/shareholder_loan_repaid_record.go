package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetShareholderLoanRepaidRecordTable(ctx *context.Context) table.Table {

	shareholderLoanRepaidRecord := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := shareholderLoanRepaidRecord.GetInfo().HideFilterArea()

	info.AddField("Rid", "rid", db.Int4)
	info.AddField("Repaid_date", "repaid_date", db.Date)
	info.AddField("Repaid_amount", "repaid_amount", db.Int8)
	info.AddField("Repaid_principal", "repaid_principal", db.Int8)
	info.AddField("Repaid_interest", "repaid_interest", db.Int8)
	info.AddField("Sl_cid", "sl_cid", db.Int4)
	info.AddField("Create_time", "create_time", db.Timestamp)
	info.AddField("Modify_time", "modify_time", db.Timestamp)

	info.SetTable("fzzl.shareholder_loan_repaid_record").SetTitle("ShareholderLoanRepaidRecord").SetDescription("ShareholderLoanRepaidRecord")

	formList := shareholderLoanRepaidRecord.GetForm()
	formList.AddField("Rid", "rid", db.Int4, form.Number)
	formList.AddField("Repaid_date", "repaid_date", db.Date, form.Datetime)
	formList.AddField("Repaid_amount", "repaid_amount", db.Int8, form.Text)
	formList.AddField("Repaid_principal", "repaid_principal", db.Int8, form.Text)
	formList.AddField("Repaid_interest", "repaid_interest", db.Int8, form.Text)
	formList.AddField("Sl_cid", "sl_cid", db.Int4, form.Number)
	formList.AddField("Create_time", "create_time", db.Timestamp, form.Datetime)
	formList.AddField("Modify_time", "modify_time", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.shareholder_loan_repaid_record").SetTitle("ShareholderLoanRepaidRecord").SetDescription("ShareholderLoanRepaidRecord")

	return shareholderLoanRepaidRecord
}
