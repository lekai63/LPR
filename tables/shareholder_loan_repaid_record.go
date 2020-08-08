package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	"github.com/lekai63/lpr/models"
)

func GetShareholderLoanRepaidRecordTable(ctx *context.Context) table.Table {

	shareholderLoanRepaidRecord := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := shareholderLoanRepaidRecord.GetInfo().HideFilterArea()

	info.AddField("序号", "id", db.Int).FieldFilterable()
	info.AddField("还款日期", "repaid_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("还款总额", "repaid_amount", db.Int8)
	info.AddField("还款本金", "repaid_principal", db.Int8)
	info.AddField("还款利息", "repaid_interest", db.Int8)
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

	dbGorm := models.Gormv2
	var slc models.ShareholderLoanContract
	formList.AddField("项目简称", "shareholder_loan_contract_id", db.Int, form.Text).FieldNotAllowAdd().FieldNotAllowEdit().
		FieldDisplay(func(model types.FieldModel) interface{} {
			// 获取第一条匹配的记录
			id := model.Value
			dbGorm.Where("id = ?", id).First(&slc)
			a := slc.Abbreviation
			if a.Valid {
				return a.String
			} else {
				return ""
			}
		})
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).FieldHide().FieldNotAllowEdit()
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime).FieldHide()

	formList.SetTable("fzzl.shareholder_loan_repaid_record").SetTitle("ShareholderLoanRepaidRecord").SetDescription("ShareholderLoanRepaidRecord")

	return shareholderLoanRepaidRecord
}
