package lpr

import (
	"fmt"
)

func Icbc() {
	fmt.Println("成功调用icbc函数")
	res := GetOneContractRepayPlan(3)
	fmt.Printf("res:%+v", res)

}
