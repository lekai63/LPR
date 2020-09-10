package inscalc

import (
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
	p.ToICBC()

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
