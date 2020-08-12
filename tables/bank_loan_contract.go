package tables

import (
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBankLoanContractTable(ctx *context.Context) table.Table {

	bankLoanContract := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := bankLoanContract.GetInfo().HideFilterArea()

	info.AddField("序号", "id", db.Int).FieldFilterable()
	// 项目简称 模糊查询
	info.AddField("项目简称", "abbreviation", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldFilterProcess(func(s string) string {
			// 即使前端错误输入带空格，在这里可以过滤空格进行sql查询
			return strings.TrimSpace(s)
		})

	info.AddField("银行借款合同号", "bank_contract_no", db.Varchar)
	info.AddField("银行借款合同名称", "bank_contract_name", db.Varchar)
	info.AddField("保理户账号", "bank_account", db.Varchar)
	info.AddField("计息方式", "interest_calc_method", db.Varchar)
	info.AddField("银行", "bank_name", db.Varchar)
	info.AddField("支行", "bank_branch", db.Varchar)
	info.AddField("借款本金", "loan_principal", db.Int8)
	info.AddField("贷款方式", "loan_method", db.Varchar)
	info.AddField("合同起始日", "contract_start_date", db.Date)
	info.AddField("合同到期日", "contract_end_date", db.Date)
	info.AddField("提款日", "actual_start_date", db.Date)
	info.AddField("定价模式", "is_lpr", db.Bool).FieldBool("LPR", "基准定价")
	info.AddField("Current_reprice_day", "current_reprice_day", db.Date).FieldHide()
	info.AddField("Current_lpr", "current_lpr", db.Int).FieldHide()
	info.AddField("Lpr_plus", "lpr_plus", db.Int).FieldHide()
	info.AddField("当前执行利率", "current_rate", db.Int).FieldDisplay(showMoney)
	info.AddField("下一重定价日", "next_reprice_day", db.Date)
	info.AddField("已还本金", "all_repaid_principal", db.Int8)
	info.AddField("已还利息", "all_repaid_interest", db.Int8)
	info.AddField("合同执行", "is_finished", db.Bool).FieldBool("已结束", "")
	info.AddField("银行联系人", "contact_person", db.Varchar)
	info.AddField("联系电话", "contact_tel", db.Varchar)
	// 对应的租赁合同id数组
	// info.AddField("lease_contract_ids", "lease_contract_ids", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp).FieldHide()
	info.AddField("Updated_at", "updated_at", db.Timestamp).FieldHide()

	info.SetTable("fzzl.bank_loan_contract").SetTitle("BankLoanContract").SetDescription("BankLoanContract")

	formList := bankLoanContract.GetForm()
	formList.AddField("序号", "id", db.Int4, form.Default).FieldHide().FieldNotAllowEdit().FieldNotAllowAdd()

	// 需处理对应1个融资合同对应2个租赁合同的情况，拟数据库表改为数组。
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
