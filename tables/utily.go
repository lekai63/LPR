package tables

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type selfTableConfig struct {
	table.Config
}

func (cc selfTableConfig) setPrimaryKey(name string, typ ...string) (result table.Config) {
	cc.PrimaryKey.Name = name
	if len(typ) > 0 {
		m := typ[0]
		result = cc.SetPrimaryKeyType(m)
	} else {
		// 为了返回 table.Config 类型
		result = cc.SetPrimaryKeyType("INT")
	}
	return
}

// showMoney 数据表格显示函数
func showMoney(model types.FieldModel) (result interface{}) {

	if len(model.Value) <= 4 {
		result = "0." + model.Value[:]
	} else {
		result = model.Value[:len(model.Value)-4] + "." + model.Value[len(model.Value)-4:len(model.Value)-2]
	}

	return
}
