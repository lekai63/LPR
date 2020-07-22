package tables

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type selfTableConfig struct {
	table.Config
}

// setPrimaryKey 自定义连接相关字段的主键
func (config selfTableConfig) setPrimaryKey(name string, typ ...string) (result table.Config) {
	config.PrimaryKey.Name = name

	if len(typ) > 0 {
		m := typ[0]
		result = config.SetPrimaryKeyType(m)
	} else {
		// 为了返回 table.Config 类型
		result = config.SetPrimaryKeyType("INT")
	}
	return
}

// showMoney 数据表格显示金额，数据库以bigint存储精度至0.01分，将该数据转换为“元”
func showMoney(model types.FieldModel) (result interface{}) {
	if len(model.Value) <= 4 {
		result = "0." + model.Value[:]
	} else {
		result = model.Value[:len(model.Value)-4] + "." + model.Value[len(model.Value)-4:len(model.Value)-2]
	}
	return
}

// showPercent 数据表格显示百分比，数据库以int存储精度至0.0001%（比如8.68%存储为86800），将该数据转换为百分数
func showPercent(model types.FieldModel) (result interface{}) {
	result = model.Value[:len(model.Value)-4] + "." + model.Value[len(model.Value)-4:len(model.Value)-2] + "%"
	return
}
