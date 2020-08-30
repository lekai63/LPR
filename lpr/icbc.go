package lpr

import (
	"context"
	"fmt"
)

func icbc(model BankRepayPlanCalcModel) {
	m, err := NewCalcModel(model)
	if err != nil {
		panic(err)
	}
	m, _ = m.AddAccruedPrincipalSeries(context.TODO())
	df := m.Brps
	fmt.Print(df.Table())

}
