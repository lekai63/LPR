package tables

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/GoAdminGroup/go-admin/template/types"
)

// showMoney 实际将小数点向左移动4位
// 数据库以bigint存储精度至0.01分，可以将该数据转换为“元”；
// 数据库以int存储精度至0.0001%（比如8.68%存储为86800），故也可以将该数据转换为百分数
func showMoney(model types.FieldModel) interface{} {
	switch n := len(model.Value); {
	case n == 0:
		return ""
	case n <= 4:
		if model.Value == "" || model.Value == "0" {
			return model.Value
		} else {
			return "0." + model.Value
		}
	default:
		return model.Value[:len(model.Value)-4] + "." + model.Value[len(model.Value)-4:len(model.Value)-2]
	}
}

// money2BigintString 将“元”字符串转换为bigint精度的字符串
func money2bigint(model types.PostFieldModel) (result interface{}) {
	val := model.Value.Value()
	return floatStr2BigintStr(val)
}

func floatStr2BigintStr(val string) (result string) {
	val = strings.TrimSpace(val)
	if val == "" || val == "0" {
		result = "0"
	} else if strings.Count(val, ".") == 1 {
		// 小数的处理
		digitals := strings.Split(val, ".")
		temp := digitals[1] + strings.Repeat("0", (4-len(digitals[1])))
		if digitals[0] == "0" {
			result = temp
		} else {
			result = digitals[0] + temp
		}

	} else {
		// 整数直接加四个0
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

// id2String convert type int32 to type string
func id2String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

// mapArr2strArr 将数据库查询数据 提取需要的字段值后，重组为[]string
func mapArr2strArr(mapArr []map[string]interface{}, col string) (strArr []string, err error) {

	if len(mapArr) == 0 {
		return
	}
	for _, item := range mapArr {
		if _, ok := item[col]; !ok {
			err = fmt.Errorf("不存在 %s 字段", col)
			break
		}
		if i, ok := item[col].(string); !ok {
			err = fmt.Errorf("map item 无法断言为string")
			break
		} else {
			strArr = append(strArr, i)
		}

	}

	return
}

// print_map 解析 map[string]interface{} 数据格式
func print_map(m map[string]interface{}) string {
	for k, v := range m {
		switch value := v.(type) {
		case nil:
			fmt.Println(k, "is nil", "null")
			return "nilValue"
		case string:
			fmt.Println(k, "is string", value)
			return "string"
		case int:
			fmt.Println(k, "is int", value)
			return "int"
		case float64:
			fmt.Println(k, "is float64", value)
			return "float64"
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range value {
				fmt.Println(i, u)
			}
			return "array"
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_map(value)
		default:
			fmt.Println(k, "is unknown type", fmt.Sprintf("%T", v))
			return "unkown"
		}
	}
	return "unkown"
}
