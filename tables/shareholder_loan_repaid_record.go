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
	"github.com/lekai63/lpr/models"
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

	var repaidInfo struct {
		amount    int64
		principal int64
		interest  int64
	}

	formList.AddField("还款总额", "repaid_amount", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			// 在SetPostValidator之后执行FieldPostFilterFn，故repaidInfo已完成数据格式转换
			// fmt.Printf("post repaidInfo is:%+v", repaidInfo)
			return repaidInfo.amount
		})
	formList.AddField("还款本金", "repaid_principal", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			// 在SetPostValidator之后执行
			return repaidInfo.principal
		})
	formList.AddField("还款利息", "repaid_interest", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			return repaidInfo.interest
		})

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
		repaidInfo.amount = cast.ToInt64(floatStr2BigintStr(values.Get("repaid_amount")))
		repaidInfo.principal = cast.ToInt64(floatStr2BigintStr(values.Get("repaid_principal")))
		repaidInfo.interest = cast.ToInt64(floatStr2BigintStr(values.Get("repaid_interest")))
		if repaidInfo.amount != repaidInfo.principal+repaidInfo.interest {
			return fmt.Errorf("实际还款金额≠实际还款本金+实际还款利息")
		}
		return nil
	})

	formList.SetPostHook(func(values form2.Values) error {
		dbGorm := models.Gormv2
		// var slc models.ShareholderLoanContract

		var r struct {
			Slcid          int32
			Totalprincipal int64
			Totalinterest  int64
		}

		// 查询并记录 更新的这笔 内部借款还款记录 所对应的内部借款合同，并计算已还本金、已还利息
		dbGorm.Table("shareholder_loan_repaid_record").
			Select("shareholder_loan_contract_id as slcid,sum(repaid_principal) as totalprincipal,sum(repaid_interest) as totalinterest").
			Where("shareholder_loan_contract_id = ?", values.Get("shareholder_loan_contract_id")).Group("shareholder_loan_contract_id").Scan(&r)

		fmt.Printf("%+v will be updated to slc", r)

		var slc models.ShareholderLoanContract
		slc.ID = r.Slcid
		slc.AllRepaidInterest.SetValid(r.Totalinterest)
		slc.AllRepaidPrincipal.SetValid(r.Totalprincipal)

		// 更新 内部借款合同 中的记录
		if e := dbGorm.Model(&slc).
			Select("AllRepaidPrincipal", "AllRepaidInterest", "UpdatedAt").
			Updates(slc).Error; e != nil {
			return e
		}

		return nil
	})
	return shareholderLoanRepaidRecord
}
