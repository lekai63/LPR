package inscalc

import (
	"fmt"
	"testing"
)

var (
	in       = 13
	expected = in
)

func TestFilterNilActualRows(t *testing.T) {

	actual, err := NewModel(int32(in))
	if err != nil {
		fmt.Println(err)
	}
	p := &actual
	p.FilterNilActualRows()
	// fmt.Println("coll table")
	// fmt.Print(p.Brps.Table())

}

func TestZZ(t *testing.T) {

	actual, err := NewModel(int32(in))
	if err != nil {
		fmt.Println(err)
	}
	p := &actual
	p.Sort("plan_date")

	// p, err = p.ToICBC(true)
	// p, err = p.ToHZBank(true)
	// p, err = p.ToABC(true)
	// p, err = p.ToCMB(true)
	// p, err = p.ToCCB(true)
	p, err = p.ToSPDB(true)
	if err != nil {
		fmt.Println(err)
	}

	// tt := time.Now()
	// p.AfterDay(civil.DateOf(tt))

	// fmt.Println("ToBank")
	// fmt.Print(p.Brps.Table())

	/* 	err = p.Update()
	   	if err != nil {
	   		fmt.Printf("更新失败:%+v", err)
	   	}

	   	err = p.Insert()
	   	if err != nil {
	   		fmt.Printf("插入失败:%+v", err)
	   	} */

}
