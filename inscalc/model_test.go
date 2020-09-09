package inscalc

import (
	"context"
	"fmt"
	"testing"
)

func TestNewBankRepayPlanCalcModel(t *testing.T) {
	var (
		in       = 1
		expected = 1
	)
	actual, err := NewModel(int32(in))

	p := &actual
	p.Sort("plan_date")
	fmt.Println("origin Table")
	fmt.Print(p.Brps.Table())
	// 注意使用括号决定计算优先级，不要直接链式调用
	(p.FillPlanDateMonthly()).CalcAccruedPrincipal(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
	ex := int(actual.Bc.ID)
	if ex != expected {
		t.Errorf("sth wrong")
	}
	fmt.Println("test table")
	fmt.Print(p.Brps.Table())

}
