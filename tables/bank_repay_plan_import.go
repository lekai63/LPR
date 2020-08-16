package tables

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/lekai63/lpr/models"
	"github.com/spf13/cast"
)

func GetBankRepayPlanImportTable(ctx *context.Context) table.Table {

	bankRepayPlan := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))
	formList := bankRepayPlan.GetForm()

	// 文件模板待完善
	formList.AddField("导入银行还款台账", "file", db.Varchar, form.File).FieldNotAllowEdit().FieldHelpMsg(`<a href="../../../uploads/csv/bank_repay_plan_model.csv" download >右击另存模板</a>日期格式需为yyyy-mm-dd,存为UTF8的csv后上传`)

	formList.SetTable("fzzl.bank_repay_plan").SetTitle("BankRepayPlan").SetDescription("BankRepayPlan")

	formList.SetInsertFn(func(values form2.Values) (err error) {
		var fileName string

		if a := values["file"]; a[0] == "" {
			err = fmt.Errorf("未上传文件")
			return
		} else {
			fileName = "./uploads/" + a[0]
			// f, err = ioutil.ReadFile(fileName)
		}

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
		plans := make([]models.BankRepayPlan, len(content)-1)
		// blc 存储项目简称，避免多次查询数据库
		var blc struct {
			ID   int32
			Abbr string
		}

		dbGorm := models.Gormv2
		for i := 1; i < len(content); i++ {
			row := content[i]
			// fmt.Println(row)
			// 若项目简称相同，则不额外查询BankLoanContract ID
			if row[0] != blc.Abbr {
				blc.Abbr = row[0]
				result := dbGorm.Debug().Table("bank_loan_contract").Select("ID").Where("abbreviation = ?", blc.Abbr).Scan(&blc)
				// fmt.Printf("result struct is : %+v", result)
				if result.RowsAffected == 0 {
					err = fmt.Errorf("未在《银行融资合同》中查到对应项目简称，请先在《银行融资合同》中新增相关合同。详细错误信息：%s", result.Error)
					break
				}

			}
			var p models.BankRepayPlan
			// p := plans[i-1] 方式 实际是无法修改plans的，因为是值传递
			// 参考 https://cloudsjhan.github.io/2018/10/27/技术周刊之golang中修改struct的slice的值/

			p.BankLoanContractID = blc.ID
			p.PlanDate, e = cast.StringToDate(row[1])
			if e != nil {
				return e
			}
			p.PlanPrincipal = cast.ToInt64(floatStr2BigintStr(row[2]))

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

	return bankRepayPlan

}
