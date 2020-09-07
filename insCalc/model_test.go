package insCalc

import (
	"fmt"
	"testing"
)

func TestNewBankRepayPlanCalcModel(t *testing.T) {
	var (
		in       = 3
		expected = 3
	)
	actual, err := NewBankRepayPlanCalcModel(int32(in))
	if err != nil {
		fmt.Println(err)
	}
	ex := int(actual.Bc.ID)
	if ex != expected {
		t.Errorf("sth wrong")
	}
}
