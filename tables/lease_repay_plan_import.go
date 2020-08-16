package tables

import (
	"fmt"
	"os"

	"encoding/csv"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/lekai63/lpr/models"
	"github.com/spf13/cast"

	"github.com/GoAdminGroup/go-admin/template/types/form"

	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"

	_ "gorm.io/driver/postgres"
)

func GetLeaseRepayPlanImportTable(ctx *context.Context) table.Table {

	leaseRepayPlan := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	// info := leaseRepayPlan.GetInfo()

	formList := leaseRepayPlan.GetForm()

	formList.AddField("导入租金台账", "file", db.Varchar, form.File).FieldNotAllowEdit()

	// formList.SetPostHook()
	// formList.SetUpdateFn()
	// formList.SetInsertFn()

	formList.SetInsertFn(func(values form2.Values) (err error) {

		var fileName string

		if a := values["file"]; a[0] == "" {
			err = fmt.Errorf("未上传文件")
			return
		} else {
			fileName = "./uploads/" + a[0]
			// f, err = ioutil.ReadFile(fileName)
		}

		// 对于小文件，一次性读取所有行
		fs1, e := os.Open(fileName)
		if e != nil {
			err = e
			return
		}
		r := csv.NewReader(fs1)
		content, e := r.ReadAll()
		if e != nil {
			err = fmt.Errorf("can not readall,err is %+v", e)
			return
		}

		// plans 定义要插入的模型数组,要传入指针，否则plans不会被更新
		plans := make([]models.LeaseRepayPlan, len(content)-1)
		// lc 存储项目简称，避免多次查询数据库
		var lc struct {
			ID   int32
			Abbr string
		}

		dbGorm := models.Gormv2
		for i := 1; i < len(content); i++ {
			row := content[i]
			// fmt.Println(row)
			// 若项目简称相同，则不额外查询LeaseContract ID
			if row[0] != lc.Abbr {
				lc.Abbr = row[0]
				result := dbGorm.Table("lease_contract").Select("ID").Where("Abbreviation = ?", lc.Abbr).Scan(&lc)
				// fmt.Printf("lc struct is : %+v", lc)
				if result.RowsAffected == 0 {
					err = fmt.Errorf("未在《租赁合同》中查到对应项目简称，请先在《租赁合同》中新增相关合同。详细错误信息：%s", result.Error)
					break
				}

			}
			var p models.LeaseRepayPlan
			// p := plans[i-1] 方式 实际是无法修改plans的，因为是值传递
			// 参考 https://cloudsjhan.github.io/2018/10/27/技术周刊之golang中修改struct的slice的值/

			p.LeaseContractID = lc.ID
			p.Period.SetValid(cast.ToInt64(row[1]))
			p.PlanDate, e = cast.StringToDate(row[2])
			if e != nil {
				return e
			}
			p.PlanAmount = cast.ToInt64(floatStr2BigintStr(row[3]))
			p.PlanPrincipal = cast.ToInt64(floatStr2BigintStr(row[4]))
			p.PlanInterest = cast.ToInt64(floatStr2BigintStr(row[5]))
			// 修改plans的值
			plans[i-1] = p

		}
		if err != nil {
			return
		}
		dbGorm.Create(&plans)

		return
	})

	formList.SetUpdateFn(func(values form2.Values) error {
		return nil
	})

	formList.SetTable("fzzl.lease_repay_plan").SetTitle("LeaseRepayPlanImport").SetDescription("LeaseRepayPlanImport")

	return leaseRepayPlan
}
