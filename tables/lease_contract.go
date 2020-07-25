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

	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"

	// import GORM
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lekai63/lpr/models"
)

func GetLeaseContractTable(ctx *context.Context) table.Table {

	dbGorm := models.Gorm
	var lesseeInfoGorm models.LesseeInfo

	leaseContract := table.NewDefaultTable(table.DefaultConfigWithDriver(db.DriverPostgresql))

	info := leaseContract.GetInfo().HideFilterArea()

	info.AddField("序号", "id", db.Int)
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
	info.AddField("Irr", "irr", db.Int).FieldDisplay(showMoney)
	info.AddField("Is_lpr", "is_lpr", db.Bool).FieldHide()
	info.AddField("Current_reprice_day", "current_reprice_day", db.Date).FieldHide()
	info.AddField("Current_LPR", "current_LPR", db.Int).FieldHide()
	info.AddField("Lpr_plus", "lpr_plus", db.Int).FieldHide()
	info.AddField("当前租息率", "current_rate", db.Int).FieldDisplay(showMoney)
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
	formList.AddField("序号", "id", db.Int, form.Number).FieldHide()
	formList.AddField("合同号", "contract_no", db.Varchar, form.Text)
	formList.AddField("承租人全称", "lessee", db.Varchar, form.Text)
	formList.AddField("项目简称", "abbreviation", db.Varchar, form.Text)
	formList.AddField("起始日", "start_date", db.Date, form.Date)
	formList.AddField("到期日", "end_date", db.Date, form.Date)

	//todo: 转换金额
	formList.AddField("手续费", "fee", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(money2bigint)
	formList.AddField("保证金", "margin", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(money2bigint)
	formList.AddField("合同本金", "contract_principal", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(money2bigint)
	formList.AddField("实际投放", "actual_principal", db.Int8, form.Text).
		FieldDisplay(showMoney).
		FieldHelpMsg("单位:元").
		FieldPostFilterFn(money2bigint)
	//todo：数据校验
	formList.AddField("期限", "term_month", db.Int2, form.Number)

	formList.AddField("标的物", "subject_matter", db.Varchar, form.Text)

	//todo: 转换百分比
	formList.AddField("Irr", "irr", db.Int, form.Rate).
		FieldDisplay(showMoney)

	// 选中基准定价后，显示4个LPR相关表单。
	// 此处有一点小bug，不影响使用，暂不修改。
	formList.AddField("定价模式", "is_lpr", db.Bool, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "基于LPR定价", Value: "true"},
			{Text: "基于基准定价", Value: "false"},
		}).FieldDefault("基于LPR定价").
		FieldOnChooseHide("基于基准定价", "current_reprice_day", "current_LPR", "lpr_plus", "next_reprice_day")

		//默认隐藏LPR表单项
	formList.AddField("当前执行利率的重定价日", "current_reprice_day", db.Date, form.Date)
	formList.AddField("当前基于的LPR利率", "current_LPR", db.Int, form.Rate)
	formList.AddField("LPR加点值", "lpr_plus", db.Int, form.Number).FieldHelpMsg("单位:bp. 请填写整数")
	formList.AddField("下一重定价日", "next_reprice_day", db.Date, form.Date)

	formList.AddField("当前租息率", "current_rate", db.Int, form.Rate).
		FieldDisplay(showMoney)

	formList.AddField("Received_principal", "received_principal", db.Int8, form.Text).FieldHide()
	formList.AddField("Received_interest", "received_interest", db.Int8, form.Text).FieldHide()

	formList.AddField("合同执行", "is_finished", db.Bool, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: "已结束", Value: "true"},
			{Text: "执行中", Value: "false"},
		}).FieldDefault("false")

	// 从lessee_info中获取承租人id，并写入本表对应字段
	var lesseeInfoId int32
	formList.AddField("承租人ID", "lessee_info_id", db.Int4, form.Number).FieldHide().FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
		return lesseeInfoId
	})

	formList.AddField("创建时间", "created_time", db.Timestamp, form.Datetime).FieldHide().FieldNotAllowEdit().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			if value.Value == nil {
				return time.Now().Format("2006-01-02 15:04:05")
			}
			return value.Value.Value()
		})
	formList.AddField("修改时间", "updated_time", db.Timestamp, form.Datetime).
		FieldHide().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			return time.Now().Format("2006-01-02 15:04:05")
		})

	formList.SetTable("fzzl.lease_contract").SetTitle("LeaseContract").SetDescription("LeaseContract")

	//数据校验
	// formList.SetPostValidator(func(values form2.Values) error {
	// 	if values.Get("sex") != "women" && values.Get("sex") != "men" {
	// 		return fmt.Errorf("error info")
	// 	}
	// 	return nil
	// })

	// 数据校验
	formList.SetPostValidator(func(values form2.Values) error {
		query := dbGorm.First(&lesseeInfoGorm, "lessee = ?", values.Get("lessee"))

		if query == nil {
			return fmt.Errorf("承租人全称不匹配，可能未在《承租人信息表》中新增它")
		}
		lesseeInfoId = query.ID
		return nil

	})

	return leaseContract
}