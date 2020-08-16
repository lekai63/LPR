package tables

import (
	"fmt"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/lekai63/lpr/models"
)

func GetBankLoanContractTable(ctx *context.Context) table.Table {

	bankLoanContract := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := bankLoanContract.GetInfo().HideFilterArea()

	info.AddField("序号", "id", db.Int)
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
	info.AddField("借款本金", "loan_principal", db.Int8).FieldDisplay(showMoney)
	info.AddField("贷款方式", "loan_method", db.Varchar)

	info.AddField("合同起始日", "contract_start_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable()
	info.AddField("提款日", "actual_start_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("合同到期日", "contract_end_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})

	info.AddField("定价模式", "is_lpr", db.Bool).FieldDisplay(func(model types.FieldModel) interface{} {
		switch model.Value {
		case "true":
			return "LPR"
		case "false":
			return "基准"
		default:
			return "未定义"
		}
	})
	info.AddField("Current_reprice_day", "current_reprice_day", db.Date).FieldHide()
	info.AddField("Current_lpr", "current_lpr", db.Int).FieldHide()
	info.AddField("Lpr_plus", "lpr_plus", db.Int).FieldHide()
	info.AddField("当前执行利率", "current_rate", db.Int).FieldDisplay(showMoney)
	info.AddField("下一重定价日", "next_reprice_day", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("已还本金", "all_repaid_principal", db.Int8).FieldDisplay(showMoney)
	info.AddField("已还利息", "all_repaid_interest", db.Int8).FieldDisplay(showMoney)
	info.AddField("合同执行", "is_finished", db.Bool).FieldDisplay(func(model types.FieldModel) interface{} {
		switch model.Value {
		case "true":
			return "已结束"
		case "false":
			return "执行中"
		default:
			return "未定义"
		}
	})
	info.AddField("银行联系人", "contact_person", db.Varchar)
	info.AddField("联系电话", "contact_tel", db.Varchar).FieldHide()
	// 对应的租赁合同id数组
	// info.AddField("lease_contract_ids", "lease_contract_ids", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp).FieldHide()
	info.AddField("Updated_at", "updated_at", db.Timestamp).FieldHide()
	// info.AddField("lease_contract_ids","lease_contract_ids",db.Varchar)

	info.SetTable("fzzl.bank_loan_contract").SetTitle("BankLoanContract").SetDescription("BankLoanContract")

	formList := bankLoanContract.GetForm()
	formList.AddField("序号", "id", db.Int, form.Default).FieldHide().FieldNotAllowEdit().FieldNotAllowAdd()

	formList.AddField("银行借款合同号", "bank_contract_no", db.Varchar, form.Text)
	formList.AddField("银行借款合同名称", "bank_contract_name", db.Varchar, form.Text)
	formList.AddField("保理户账号", "bank_account", db.Varchar, form.Text)
	formList.AddField("计息方式", "interest_calc_method", db.Varchar, form.Text)
	formList.AddField("银行", "bank_name", db.Varchar, form.Text).FieldHelpMsg("填写银行总行名称，比如工商银行、杭州银行")
	formList.AddField("支行", "bank_branch", db.Varchar, form.Text).FieldHelpMsg("填写支行名称，比如杭州之江支行")
	formList.AddField("借款本金", "loan_principal", db.Int8, form.Text)
	formList.AddField("贷款方式", "loan_method", db.Varchar, form.Text)
	formList.AddField("合同起始日", "contract_start_date", db.Date, form.Date)
	formList.AddField("提款日", "actual_start_date", db.Date, form.Date)
	formList.AddField("合同到期日", "contract_end_date", db.Date, form.Date)

	formList.AddField("定价模式", "is_lpr", db.Bool, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "基于LPR定价", Value: "true"},
			{Text: "基于基准定价", Value: "false"},
		}).FieldDefault("基于LPR定价").
		FieldOnChooseHide("基于基准定价", "current_reprice_day", "current_lpr", "lpr_plus", "next_reprice_day")
	formList.AddField("当前执行利率的重定价日", "current_reprice_day", db.Date, form.Date).FieldPostFilterFn(allowReturnNullString)
	formList.AddField("当前基于的LPR利率", "current_lpr", db.Int, form.Rate).FieldPostFilterFn(allowReturnNullString)
	formList.AddField("LPR加点值", "lpr_plus", db.Int, form.Number).FieldHelpMsg("单位:bp. 请填写整数").FieldPostFilterFn(allowReturnNullString)
	formList.AddField("下一重定价日", "next_reprice_day", db.Date, form.Date).FieldPostFilterFn(allowReturnNullString)

	formList.AddField("当前执行利率", "current_rate", db.Int, form.Rate).
		FieldDisplay(showMoney).FieldPostFilterFn(money2bigint)

	formList.AddField("All_repaid_principal", "all_repaid_principal", db.Int8, form.Text).FieldHide().
		FieldPostFilterFn(func(model types.PostFieldModel) interface{} {
			if len(model.Value) == 0 || model.Value.Value() == "" {
				return "0"
			}
			return model.Value.Value()
		})
	formList.AddField("All_repaid_interest", "all_repaid_interest", db.Int8, form.Text).FieldHide().
		FieldPostFilterFn(func(model types.PostFieldModel) interface{} {
			if len(model.Value) == 0 || model.Value.Value() == "" {
				return "0"
			}
			return model.Value.Value()
		})
	formList.AddField("合同执行", "is_finished", db.Bool, form.Switch).
		FieldOptions(
			types.FieldOptions{
				types.FieldOption{Text: "已结束", Value: "true"},
				types.FieldOption{Text: "执行中", Value: "false"},
			}).FieldDefault("false")
	formList.AddField("联系人", "contact_person", db.Varchar, form.Text)
	formList.AddField("联系电话", "contact_tel", db.Varchar, form.Text)
	formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime).FieldHide().FieldNotAllowEdit().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			if value.Value == nil {
				return time.Now().Format("2006-01-02 15:04:05")
			}
			return value.Value.Value()
		})
	formList.AddField("修改时间", "updated_at", db.Timestamp, form.Datetime).
		FieldHide().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			return time.Now().Format("2006-01-02 15:04:05")
		})

	formList.AddField("对应的项目", "lease_contract_ids", db.Varchar, form.Select).
		FieldOptionsFromTable("lease_contract", "abbreviation", "id").
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			val := value.Value
			// 使用简单的array转string +{} 存储到pgsql数组
			str := strings.Replace(strings.Trim(fmt.Sprint(val), "[]"), " ", ",", -1)
			return "{" + str + "}"
		}).
		FieldDisplay(func(model types.FieldModel) interface{} {
			row := model.Row
			item := row["lease_contract_ids"]
			if item == nil {
				return []string{}
			}
			if a := item.([]uint8); len(a) == 0 {
				return []string{}
			}

			str := fmt.Sprintf("%s", item) // row item type: []unit8
			conn := models.GlobalConn
			queryCol := "abbreviation"
			sel := "SELECT " + queryCol + " FROM lease_contract WHERE id = any($1)"
			if abbr, err := conn.Query(sel, str); err != nil {
				return err
			} else {
				r, _ := mapArr2strArr(abbr, queryCol)
				return r
			}

		})

	formList.SetTable("fzzl.bank_loan_contract").SetTitle("BankLoanContract").SetDescription("BankLoanContract")

	return bankLoanContract
}
