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

	info.AddField("Id", "id", db.Int4).FieldFilterable()
	info.AddField("Repaid_date", "repaid_date", db.Date)
	info.AddField("Repaid_amount", "repaid_amount", db.Int8)
	info.AddField("Repaid_principal", "repaid_principal", db.Int8)
	info.AddField("Repaid_interest", "repaid_interest", db.Int8)
	info.AddField("Shareholder_loan_contract_id", "shareholder_loan_contract_id", db.Int4)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("fzzl.shareholder_loan_repaid_record").SetTitle("ShareholderLoanRepaidRecord").SetDescription("ShareholderLoanRepaidRecord")

	formList := shareholderLoanRepaidRecord.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default)
	formList.AddField("Repaid_date", "repaid_date", db.Date, form.Datetime)
	formList.AddField("Repaid_amount", "repaid_amount", db.Int8, form.Text)
	formList.AddField("Repaid_principal", "repaid_principal", db.Int8, form.Text)
	formList.AddField("Repaid_interest", "repaid_interest", db.Int8, form.Text)
	formList.AddField("Shareholder_loan_contract_id", "shareholder_loan_contract_id", db.Int4, form.Number)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.shareholder_loan_repaid_record").SetTitle("ShareholderLoanRepaidRecord").SetDescription("ShareholderLoanRepaidRecord")

	return shareholderLoanRepaidRecord
}
