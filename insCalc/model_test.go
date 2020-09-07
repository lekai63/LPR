package insCalc

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

	// fmt.Print(actual.Brps.Table())
	// actual.fillInsPlanDateICBC().CalcAccruedPrincipal(context.TODO())
	// actual.Brps.Sort(context.TODO(), []dataframe.SortKey{
	// 	{Key: "plan_date", Desc: false},
	// })
	// fmt.Print(actual.Brps.Table())
	p := &actual
	p.fillInsPlanDateICBC().CalcAccruedPrincipal(context.TODO())

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
