package tables

import (
	"fmt"
	// "strconv"
	// "time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

	// "github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	// "github.com/lekai63/lpr/models"

	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func GetLeaseRepayPlanImportTable(ctx *context.Context) table.Table {

	leaseRepayPlan := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	// info := leaseRepayPlan.GetInfo()

	// dbGorm := models.Gorm
	// var leaseContractGorm models.LeaseContract
	formList := leaseRepayPlan.GetForm()

	formList.AddField("导入租金台账", "custom", db.Varchar, form.File).FieldNotAllowEdit()

	// formList.SetPostHook()
	// formList.SetUpdateFn()
	// formList.SetInsertFn()

	var f *excelize.File
	formList.SetPostValidator(func(values form2.Values) (err error) {
		if values["custom"] == nil {
			err = fmt.Errorf("未上传文件")
		} else {
			fileName := "./uploads/" + values["custom"][0]
			f, err = excelize.OpenFile(fileName)

		}
		return
	})

	formList.SetPostValidator(func(values form2.Values) error {
		//   if values.Get("actual_amount") != (values.Get("actual_principal")+values.Get("actual_interest")) {
		// 	  err = fmt.Errorf("实际还款金额≠实际还款本金+实际还款利息")
		// 	  return
		//   }
		fmt.Printf("%s", values.Get("actual_amount"))
		return nil

	})

	formList.SetTable("fzzl.lease_repay_plan").SetTitle("LeaseRepayPlan").SetDescription("LeaseRepayPlan")

	return leaseRepayPlan
}
