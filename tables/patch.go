package tables

import (
	"database/sql"
	"strings"

	"github.com/GoAdminGroup/go-admin/template/types"
)

// showMoney 实际将小数点向左移动4位
// 数据库以bigint存储精度至0.01分，可以将该数据转换为“元”；
// 数据库以int存储精度至0.0001%（比如8.68%存储为86800），故也可以将该数据转换为百分数
func showMoney(model types.FieldModel) (result interface{}) {
	if model.Value == "" || model.Value == "0" {
		result = model.Value
	} else if len(model.Value) <= 4 {
		result = "0." + model.Value
	} else {
		result = model.Value[:len(model.Value)-4] + "." + model.Value[len(model.Value)-4:len(model.Value)-2]
	}
	return
}

// money2BigintString 将“元”字符串转换为bigint精度的字符串
func money2bigint(model types.PostFieldModel) (result interface{}) {
	val := model.Value.Value()
	if val == "" || val == "0" {
		result = "0"
	} else if strings.Count(val, ".") == 1 {
		// 小数的处理
		digitals := strings.Split(val, ".")
		temp := digitals[1] + strings.Repeat("0", (4-len(digitals[1])))
		result = digitals[0] + temp
	} else {
		result = val + "0000"
	}
	return
}

// 允许null返回
func allowReturnNullString(model types.PostFieldModel) interface{} {
	if len(model.Value) == 0 || model.Value.Value() == "" {
		return sql.NullString{}
	}
	return model.Value.Value()
}
