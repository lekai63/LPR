package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetShareholderLoanContractTable(ctx *context.Context) table.Table {

	shareholderLoanContract := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := shareholderLoanContract.GetInfo().HideFilterArea()

	info.AddField("Sl_cid", "sl_cid", db.Int4)
	info.AddField("Creditor", "creditor", db.Varchar)
	info.AddField("Abbreviation", "abbreviation", db.Varchar)
	info.AddField("Loan_principal", "loan_principal", db.Int8)
	info.AddField("Loan_rate", "loan_rate", db.Int4)
	info.AddField("Loan_contract_no", "loan_contract_no", db.Varchar)
	info.AddField("Loan_start_date", "loan_start_date", db.Date)
	info.AddField("Loan_end_date", "loan_end_date", db.Date)
	info.AddField("All_repaid_principal", "all_repaid_principal", db.Int8)
	info.AddField("All_repaid_interest", "all_repaid_interest", db.Int8)
	info.AddField("Is_finished", "is_finished", db.Bool)
	info.AddField("Create_time", "create_time", db.Timestamp)
	info.AddField("Modify_time", "modify_time", db.Timestamp)

	info.SetTable("fzzl.shareholder_loan_contract").SetTitle("ShareholderLoanContract").SetDescription("ShareholderLoanContract")

	formList := shareholderLoanContract.GetForm()
	formList.AddField("Sl_cid", "sl_cid", db.Int4, form.Number)
	formList.AddField("Creditor", "creditor", db.Varchar, form.Text)
	formList.AddField("Abbreviation", "abbreviation", db.Varchar, form.Text)
	formList.AddField("Loan_principal", "loan_principal", db.Int8, form.Text)
	formList.AddField("Loan_rate", "loan_rate", db.Int4, form.Number)
	formList.AddField("Loan_contract_no", "loan_contract_no", db.Varchar, form.Text)
	formList.AddField("Loan_start_date", "loan_start_date", db.Date, form.Datetime)
	formList.AddField("Loan_end_date", "loan_end_date", db.Date, form.Datetime)
	formList.AddField("All_repaid_principal", "all_repaid_principal", db.Int8, form.Text)
	formList.AddField("All_repaid_interest", "all_repaid_interest", db.Int8, form.Text)
	formList.AddField("Is_finished", "is_finished", db.Bool, form.Text)
	formList.AddField("Create_time", "create_time", db.Timestamp, form.Datetime)
	formList.AddField("Modify_time", "modify_time", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.shareholder_loan_contract").SetTitle("ShareholderLoanContract").SetDescription("ShareholderLoanContract")

	return shareholderLoanContract
}
