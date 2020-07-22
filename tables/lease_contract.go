package tables

import (
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetLeaseContractTable(ctx *context.Context) table.Table {
	cc := selfTableConfig{table.DefaultConfigWithDriver(db.DriverPostgresql)}
	leaseContract := table.NewDefaultTable(cc.setPrimaryKey("cid"))

	info := leaseContract.GetInfo().HideFilterArea()

	info.AddField("序号", "cid", db.Int)
	info.AddField("合同号", "contract_no", db.Varchar)
	info.AddField("Lessee", "lessee", db.Varchar).FieldHide()
	// 项目简称 模糊查询
	info.AddField("项目简称", "abbreviation", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldFilterProcess(func(s string) string {
			// 即使前端错误输入带空格，在这里可以过滤空格进行sql查询
			return strings.TrimSpace(s)
		})
	info.AddField("起始日", "start_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("到期日", "end_date", db.Date).
		FieldSubstr(0, 10).
		FieldSortable().
		FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("手续费", "fee", db.Int8).
		FieldDisplay(showMoney)
	info.AddField("保证金", "margin", db.Int8).
		FieldDisplay(showMoney)
	info.AddField("合同本金", "contract_principal", db.Int8).
		FieldDisplay(showMoney)
	info.AddField("实际投放", "actual_principal", db.Int8).
		FieldDisplay(showMoney)
	info.AddField("期限", "term_month", db.Int2).FieldDisplay(func(model types.FieldModel) interface{} {
		return model.Value + "月"
	})
	info.AddField("标的物", "subject_matter", db.Varchar).FieldHide()
	info.AddField("Irr", "irr", db.Int).FieldDisplay(showPercent)
	info.AddField("Is_lpr", "is_lpr", db.Bool).FieldHide()
	info.AddField("Current_reprice_day", "current_reprice_day", db.Date).FieldHide()
	info.AddField("Current_LPR", "current_LPR", db.Int).FieldHide()
	info.AddField("Lpr_plus", "lpr_plus", db.Int).FieldHide()
	info.AddField("当前租息率", "current_rate", db.Int).FieldDisplay(showPercent)
	info.AddField("Next_reprice_day", "next_reprice_day", db.Date).FieldHide()
	info.AddField("已收本金", "received_principal", db.Int8).
		FieldDisplay(showMoney)
	info.AddField("已收利息", "received_interest", db.Int8).
		FieldDisplay(showMoney)
	info.AddField("合同执行", "is_finished", db.Bool).FieldBool("已结束", "")
	info.AddField("Customer_id", "customer_id", db.Int4).FieldHide()
	info.AddField("Create_time", "create_time", db.Timestamp).FieldHide()
	info.AddField("Modify_time", "modify_time", db.Timestamp).FieldHide()

	info.SetTable("fzzl.lease_contract").SetTitle("LeaseContract").SetDescription("LeaseContract")

	formList := leaseContract.GetForm()
	formList.AddField("Cid", "cid", db.Int4, form.Number)
	formList.AddField("Contract_no", "contract_no", db.Varchar, form.Text)
	formList.AddField("Lessee", "lessee", db.Varchar, form.Text)
	formList.AddField("Abbreviation", "abbreviation", db.Varchar, form.Text)
	formList.AddField("Start_date", "start_date", db.Date, form.Datetime)
	formList.AddField("End_date", "end_date", db.Date, form.Datetime)
	formList.AddField("Fee", "fee", db.Int8, form.Text)
	formList.AddField("Margin", "margin", db.Int8, form.Text)
	formList.AddField("Contract_principal", "contract_principal", db.Int8, form.Text)
	formList.AddField("Actual_principal", "actual_principal", db.Int8, form.Text)
	formList.AddField("Term_month", "term_month", db.Int2, form.Text)
	formList.AddField("Subject_matter", "subject_matter", db.Varchar, form.Text)
	formList.AddField("Irr", "irr", db.Int4, form.Number)
	formList.AddField("Is_lpr", "is_lpr", db.Bool, form.Text)
	formList.AddField("Current_reprice_day", "current_reprice_day", db.Date, form.Datetime)
	formList.AddField("Current_LPR", "current_LPR", db.Int4, form.Number)
	formList.AddField("Lpr_plus", "lpr_plus", db.Int4, form.Number)
	formList.AddField("Current_rate", "current_rate", db.Int4, form.Number)
	formList.AddField("Next_reprice_day", "next_reprice_day", db.Date, form.Datetime)
	formList.AddField("Received_principal", "received_principal", db.Int8, form.Text)
	formList.AddField("Received_interest", "received_interest", db.Int8, form.Text)
	formList.AddField("Is_finished", "is_finished", db.Bool, form.Text)
	formList.AddField("Customer_id", "customer_id", db.Int4, form.Number)
	formList.AddField("Create_time", "create_time", db.Timestamp, form.Datetime)
	formList.AddField("Modify_time", "modify_time", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.lease_contract").SetTitle("LeaseContract").SetDescription("LeaseContract")

	return leaseContract
}
