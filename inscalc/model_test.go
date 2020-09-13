package inscalc

import (
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/civil"
)

func TestNewBankRepayPlanCalcModel(t *testing.T) {
	var (
		in       = 1
		expected = 1
	)
	actual, err := NewModel(int32(in))
	if err != nil {
		fmt.Println(err)
	}
	p := &actual
	p.Sort("plan_date")
	fmt.Println("origin table")
	fmt.Print(p.Brps.Table())

	p, err = p.ToICBC(false)
	if err != nil {
		fmt.Println(err)
	}
	ex := int(p.Bc.ID)
	if ex != expected {
		t.Errorf("sth wrong")
	}

	tt := time.Now()
	p.AfterDay(civil.DateOf(tt))
	fmt.Println("test table")
	fmt.Print(p.Brps.Table())

	/* 	err = p.Update()
	   	if err != nil {
	   		fmt.Printf("更新失败:%+v", err)
	   	}

	   	err = p.Insert()
	   	if err != nil {
	   		fmt.Printf("插入失败:%+v", err)
	   	} */

}
