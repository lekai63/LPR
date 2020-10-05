package inscalc

import (
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
