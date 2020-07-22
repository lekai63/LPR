package tables

import (
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	"github.com/GoAdminGroup/go-admin/template/types"
)

func GetLesseeInfoTable(ctx *context.Context) table.Table {

	cc := selfTableConfig{table.DefaultConfigWithDriver(db.DriverPostgresql)}

	lesseeInfo := table.NewDefaultTable(cc.setPrimaryKey("customer_id"))

	info := lesseeInfo.GetInfo().HideFilterArea()

	info.AddField("序号", "customer_id", db.Int).FieldSortable()
	info.AddField("营业执照", "business_license", db.Varchar)
	info.AddField("承租人全称", "lessee", db.Varchar).FieldHide()
	info.AddField("承租人", "short_name", db.Varchar)
	info.AddField("开票邮箱", "email", db.Varchar)
	info.AddField("Contact_person", "contact_person", db.Varchar).FieldHide()
	info.AddField("Contact_tel", "contact_tel", db.Varchar).FieldHide()
	info.AddField("项目经理", "customer_manager", db.Varchar).FieldSortable()
	info.AddField("贷后经理", "risk_manager", db.Varchar).FieldSortable()
	info.AddField("Create_time", "create_time", db.Timestamp).FieldHide()
	info.AddField("Modify_time", "modify_time", db.Timestamp).FieldHide()

	info.SetTable("fzzl.lessee_info").SetTitle("LesseeInfo").SetDescription("LesseeInfo")

	formList := lesseeInfo.GetForm()
	formList.AddField("序号", "customer_id", db.Int, form.Text).FieldNotAllowEdit()
	formList.AddField("营业执照", "business_license", db.Varchar, form.Text)
	formList.AddField("承租人全称", "lessee", db.Varchar, form.Text)
	formList.AddField("承租人简称", "short_name", db.Varchar, form.Text)
	formList.AddField("开票邮箱", "email", db.Varchar, form.Email)
	formList.AddField("联系人", "contact_person", db.Varchar, form.Text)
	formList.AddField("联系电话", "contact_tel", db.Varchar, form.Text)
	formList.AddField("项目经理", "customer_manager", db.Varchar, form.Text)
	formList.AddField("贷后经理", "risk_manager", db.Varchar, form.Text)
	formList.AddField("创建时间", "create_time", db.Timestamp, form.Datetime).FieldHide().FieldNotAllowEdit()

	// t := time.Now().Format("2006-01-02 15:04:05")
	formList.AddField("修改时间", "modify_time", db.Timestamp, form.Datetime).
		FieldHide().
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			return time.Now().Format("2006-01-02 15:04:05.000")
		})

	formList.SetTable("fzzl.lessee_info").SetTitle("LesseeInfo").SetDescription("LesseeInfo")

	return lesseeInfo
}
