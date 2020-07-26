package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBankLoanContractTable(ctx *context.Context) table.Table {

	bankLoanContract := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := bankLoanContract.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).FieldFilterable()
	info.AddField("Lease_contract_id", "lease_contract_id", db.Int4)
	info.AddField("Bank_contract_no", "bank_contract_no", db.Varchar)
	info.AddField("Bank_contract_name", "bank_contract_name", db.Varchar)
	info.AddField("Bank_account", "bank_account", db.Varchar)
	info.AddField("Interest_calc_method", "interest_calc_method", db.Varchar)
	info.AddField("Bank_name", "bank_name", db.Varchar)
	info.AddField("Bank_branch", "bank_branch", db.Varchar)
	info.AddField("Loan_principal", "loan_principal", db.Int8)
	info.AddField("Loan_method", "loan_method", db.Varchar)
	info.AddField("Contract_start_date", "contract_start_date", db.Date)
	info.AddField("Contract_end_date", "contract_end_date", db.Date)
	info.AddField("Actual_start_date", "actual_start_date", db.Date)
	info.AddField("Is_lpr", "is_lpr", db.Bool)
	info.AddField("Current_reprice_day", "current_reprice_day", db.Date)
	info.AddField("Current_lpr", "current_lpr", db.Int4)
	info.AddField("Lpr_plus", "lpr_plus", db.Int4)
	info.AddField("Current_rate", "current_rate", db.Int4)
	info.AddField("Next_reprice_day", "next_reprice_day", db.Date)
	info.AddField("All_repaid_principal", "all_repaid_principal", db.Int8)
	info.AddField("All_repaid_interest", "all_repaid_interest", db.Int8)
	info.AddField("Is_finished", "is_finished", db.Bool)
	info.AddField("Contact_person", "contact_person", db.Varchar)
	info.AddField("Contact_tel", "contact_tel", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("fzzl.bank_loan_contract").SetTitle("BankLoanContract").SetDescription("BankLoanContract")

	formList := bankLoanContract.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default)
	formList.AddField("Lease_contract_id", "lease_contract_id", db.Int4, form.Number)
	formList.AddField("Bank_contract_no", "bank_contract_no", db.Varchar, form.Text)
	formList.AddField("Bank_contract_name", "bank_contract_name", db.Varchar, form.Text)
	formList.AddField("Bank_account", "bank_account", db.Varchar, form.Text)
	formList.AddField("Interest_calc_method", "interest_calc_method", db.Varchar, form.Text)
	formList.AddField("Bank_name", "bank_name", db.Varchar, form.Text)
	formList.AddField("Bank_branch", "bank_branch", db.Varchar, form.Text)
	formList.AddField("Loan_principal", "loan_principal", db.Int8, form.Text)
	formList.AddField("Loan_method", "loan_method", db.Varchar, form.Text)
	formList.AddField("Contract_start_date", "contract_start_date", db.Date, form.Datetime)
	formList.AddField("Contract_end_date", "contract_end_date", db.Date, form.Datetime)
	formList.AddField("Actual_start_date", "actual_start_date", db.Date, form.Datetime)
	formList.AddField("Is_lpr", "is_lpr", db.Bool, form.Text)
	formList.AddField("Current_reprice_day", "current_reprice_day", db.Date, form.Datetime)
	formList.AddField("Current_lpr", "current_lpr", db.Int4, form.Number)
	formList.AddField("Lpr_plus", "lpr_plus", db.Int4, form.Number)
	formList.AddField("Current_rate", "current_rate", db.Int4, form.Number)
	formList.AddField("Next_reprice_day", "next_reprice_day", db.Date, form.Datetime)
	formList.AddField("All_repaid_principal", "all_repaid_principal", db.Int8, form.Text)
	formList.AddField("All_repaid_interest", "all_repaid_interest", db.Int8, form.Text)
	formList.AddField("Is_finished", "is_finished", db.Bool, form.Text)
	formList.AddField("Contact_person", "contact_person", db.Varchar, form.Text)
	formList.AddField("Contact_tel", "contact_tel", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.bank_loan_contract").SetTitle("BankLoanContract").SetDescription("BankLoanContract")

	return bankLoanContract
}
