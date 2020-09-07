package insCalc

import (
	"fmt"
	"testing"
)

func TestNewBankRepayPlanCalcModel(t *testing.T) {
	var (
		in       = 5
		expected = 5
	)
	actual, err := NewModel(int32(in))
	if err != nil {
		fmt.Println(err)
	}
	ex := int(actual.Bc.ID)
	if ex != expected {
		t.Errorf("sth wrong")
	}
	fmt.Println("test table")
	fmt.Print(actual.Brps.Table())

}
