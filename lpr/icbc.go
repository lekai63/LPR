package lpr

import (
	"fmt"
)

func Icbc() {
	fmt.Println("成功调用icbc函数")
	res,err := GetOneContractRepayPlan(3)
	if err !=nil {
		fmt.Errorf("Get合同失败:%s",err)
	}
	fmt.Printf("\n res:%+v", res)

}