package tables

import (
	"fmt"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/spf13/cast"
)

func GetShareholderLoanRepaidRecordTable(ctx *context.Context) table.Table {

	shareholderLoanRepaidRecord := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := shareholderLoanRepaidRecord.GetInfo().HideFilterArea()

	info.AddField("序号", "id", db.Int).FieldFilterable()
	info.AddField("还款日期", "repaid_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("还款总额", "repaid_amount", db.Int8).FieldDisplay(showMoney)
	info.AddField("还款本金", "repaid_principal", db.Int8).FieldDisplay(showMoney)
	info.AddField("还款利息", "repaid_interest", db.Int8).FieldDisplay(showMoney)
	info.AddField("项目简称", "abbreviation", db.Varchar).FieldJoin(types.Join{
		Table:     "shareholder_loan_contract",    // 连表的表名，对应的外面的表的表名
		Field:     "shareholder_loan_contract_id", // 要连表的字段,即本表的字段
		JoinField: "id",                           // 连表的表的字段，对应的外面的表字段
	})
	info.AddField("Created_at", "created_at", db.Timestamp).FieldHide()
	info.AddField("Updated_at", "updated_at", db.Timestamp).FieldHide()

	info.SetTable("fzzl.shareholder_loan_repaid_record").SetTitle("ShareholderLoanRepaidRecord").SetDescription("ShareholderLoanRepaidRecord")

	formList := shareholderLoanRepaidRecord.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).FieldHide().FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("还款日期", "repaid_date", db.Date, form.Date)
	formList.AddField("还款总额", "repaid_amount", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元")
	formList.AddField("还款本金", "repaid_principal", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元")
	formList.AddField("还款利息", "repaid_interest", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元")

	formList.AddField("项目简称", "shareholder_loan_contract_id", db.Int, form.SelectSingle).
		FieldOptionsFromTable("shareholder_loan_contract", "abbreviation", "id")

	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).FieldHide().FieldNotAllowEdit().FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
		if value.Value == nil {
			return time.Now().Format("2006-01-02 15:04:05")
		}
		return value.Value.Value()
	})
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).FieldHide().FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
		return time.Now().Format("2006-01-02 15:04:05")
	})

	formList.SetTable("fzzl.shareholder_loan_repaid_record").SetTitle("ShareholderLoanRepaidRecord").SetDescription("ShareholderLoanRepaidRecord")

	formList.SetPostValidator(func(values form2.Values) error {

		// 校验还款额
		amount := cast.ToInt64(floatStr2BigintStr(values.Get("repaid_amount")))
		principal := cast.ToInt64(floatStr2BigintStr(values.Get("repaid_principal")))
		interest := cast.ToInt64(floatStr2BigintStr(values.Get("repaid_interest")))
		if amount != principal+interest {
			return fmt.Errorf("实际还款金额≠实际还款本金+实际还款利息")
		}
		return nil
	})

	return shareholderLoanRepaidRecord
}
