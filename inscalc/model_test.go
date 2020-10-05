package inscalc

import (
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/guregu/null"
)

func TestReprice(t *testing.T) {
	lprplus := null.NewInt(4500, true)
	var (
		in = Option{
			Method:  "yearly",
			ExeRate: 47500,
			LprPlus: lprplus,
		}
		expected = Option{
			Method:  "yearly",
			ExeRate: 43000,
			LprPlus: lprplus,
		}
	)
	d := civil.Date{
		Year:  2020,
		Month: time.August,
		Day:   15,
	}
	ex := in.reprice(d)
	if ex.ExeRate != expected.ExeRate {
		t.Errorf("sth wrong")
	}

}

// TODO:pg中将招行、建行三笔合同的current Rate 改回到43000
func TestNewBankRepayPlanCalcModel(t *testing.T) {
	var (
		in       = 12
		expected = in
	)
	actual, err := NewModel(int32(in))
	if err != nil {
		fmt.Println(err)
	}
	p := &actual
	p.Sort("plan_date")
	fmt.Println("origin table")
	fmt.Print(p.Brps.Table())

	// p, err = p.ToICBC(true)
	// p, err = p.ToHZBank(true)
	p, err = p.ToABC(true)
	// p, err = p.ToCMB(true)
	// p, err = p.ToCCB(true)
	if err != nil {
		fmt.Println(err)
	}
	ex := int(p.Bc.ID)
	if ex != expected {
		t.Errorf("sth wrong")
	}

	// tt := time.Now()
	// p.AfterDay(civil.DateOf(tt))
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
