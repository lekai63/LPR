package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetLesseeInfoTable(ctx *context.Context) table.Table {

	cc := selfTableConfig{table.DefaultConfigWithDriver(db.DriverPostgresql)}

	lesseeInfo := table.NewDefaultTable(cc.setPrimaryKey("customer_id"))

	info := lesseeInfo.GetInfo().HideFilterArea()

	info.AddField("Customer_id", "customer_id", db.Int4)
	info.AddField("Business_license", "business_license", db.Varchar)
	info.AddField("Lessee", "lessee", db.Varchar)
	info.AddField("Short_name", "short_name", db.Varchar)
	info.AddField("Email", "email", db.Varchar)
	info.AddField("Contact_person", "contact_person", db.Varchar)
	info.AddField("Contact_tel", "contact_tel", db.Varchar)
	info.AddField("Customer_manager", "customer_manager", db.Varchar)
	info.AddField("Risk_manager", "risk_manager", db.Varchar)
	info.AddField("Create_time", "create_time", db.Timestamp)
	info.AddField("Modify_time", "modify_time", db.Timestamp)

	info.SetTable("fzzl.lessee_info").SetTitle("LesseeInfo").SetDescription("LesseeInfo")

	formList := lesseeInfo.GetForm()
	formList.AddField("Customer_id", "customer_id", db.Int4, form.Number)
	formList.AddField("Business_license", "business_license", db.Varchar, form.Text)
	formList.AddField("Lessee", "lessee", db.Varchar, form.Text)
	formList.AddField("Short_name", "short_name", db.Varchar, form.Text)
	formList.AddField("Email", "email", db.Varchar, form.Email)
	formList.AddField("Contact_person", "contact_person", db.Varchar, form.Text)
	formList.AddField("Contact_tel", "contact_tel", db.Varchar, form.Text)
	formList.AddField("Customer_manager", "customer_manager", db.Varchar, form.Text)
	formList.AddField("Risk_manager", "risk_manager", db.Varchar, form.Text)
	formList.AddField("Create_time", "create_time", db.Timestamp, form.Datetime)
	formList.AddField("Modify_time", "modify_time", db.Timestamp, form.Datetime)

	formList.SetTable("fzzl.lessee_info").SetTitle("LesseeInfo").SetDescription("LesseeInfo")

	return lesseeInfo
}
