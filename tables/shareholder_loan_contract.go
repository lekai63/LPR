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

	// use pg for array https://www.opsdash.com/blog/postgres-arrays-golang.html
	"github.com/lekai63/lpr/models"
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
	formList.AddField("序号", "id", db.Int4, form.Default).FieldHide().FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("出借人", "creditor", db.Varchar, form.Text)
	formList.AddField("项目简称", "abbreviation", db.Varchar, form.Text)
	formList.AddField("借款本金", "loan_principal", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(money2bigint)
	formList.AddField("借款利率", "loan_rate", db.Int, form.Rate).FieldDisplay(showMoney).FieldPostFilterFn(money2bigint)
	formList.AddField("借款合同号", "loan_contract_no", db.Varchar, form.Text)
	formList.AddField("起息日", "loan_start_date", db.Date, form.Date)
	formList.AddField("到期日", "loan_end_date", db.Date, form.Date)
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
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).FieldHide().FieldNotAllowEdit().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			if value.Value == nil {
				return time.Now().Format("2006-01-02 15:04:05")
			}
			return value.Value.Value()
		})
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).FieldHide().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			return time.Now().Format("2006-01-02 15:04:05")
		})

	//	formList.AddField("对应的lease_contract_ids", "lease_contract_ids", db.Varchar, form.Array)
	// form.Array 还不成熟,尽量以 form.Select 替代
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

	formList.SetTable("fzzl.shareholder_loan_contract").SetTitle("ShareholderLoanContract").SetDescription("ShareholderLoanContract")

	return shareholderLoanContract
}

/*
// https://www.opsdash.com/blog/postgres-arrays-golang.html
func getArrbMaybeNull(db *sql.DB, str string) (tags []string) {
	sel := "SELECT abbreviation FROM lease_contract WHERE id = any($1)"
	if err := db.QueryRow(sel, str).Scan(pq.Array(&tags)); err != nil {
		log.Fatal(err)
	}
	return
}

*/
