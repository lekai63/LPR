package tables

import (
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetShareholderLoanContractTable(ctx *context.Context) table.Table {

	shareholderLoanContract := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := shareholderLoanContract.GetInfo().HideFilterArea()

	info.AddField("序号", "id", db.Int)
	info.AddField("出借人", "creditor", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "浙江省国有资本运营有限公司" {
			return "浙资运营"
		} else {
			return model.Value
		}
	})
	info.AddField("项目简称", "abbreviation", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldFilterProcess(func(s string) string {
			// 即使前端错误输入带空格，在这里可以过滤空格进行sql查询
			return strings.TrimSpace(s)
		})
	info.AddField("借款本金", "loan_principal", db.Int8).FieldDisplay(showMoney)
	info.AddField("借款利率", "loan_rate", db.Int).FieldDisplay(showMoney)
	info.AddField("借款合同号", "loan_contract_no", db.Varchar)
	info.AddField("起息日", "loan_start_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("到期日", "loan_end_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("已还本金", "all_repaid_principal", db.Int8).
		FieldDisplay(showMoney)
	info.AddField("已还利息", "all_repaid_interest", db.Int8).
		FieldDisplay(showMoney)
	info.AddField("合同执行", "is_finished", db.Bool).FieldBool("已结束", "")
	info.AddField("Created_at", "created_at", db.Timestamp).FieldHide()
	info.AddField("Updated_at", "updated_at", db.Timestamp).FieldHide()

	info.SetTable("fzzl.shareholder_loan_contract").SetTitle("ShareholderLoanContract").SetDescription("ShareholderLoanContract")

	formList := shareholderLoanContract.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default)
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
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.shareholder_loan_contract").SetTitle("ShareholderLoanContract").SetDescription("ShareholderLoanContract")

	return shareholderLoanContract
}
