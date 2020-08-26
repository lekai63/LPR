package lpr

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/modules/db"
	_ "gorm.io/driver/postgres"

	// _ "github.com/lib/pq"
	"github.com/tobgu/qframe"
	qsql "github.com/tobgu/qframe/config/sql"
)

func GetQframe(c db.Connection) {
	// Create a new in-memory SQLite database.
	sqldb := c.GetDB("default")

	// Create a new QFrame to populate our table with.
	/* qf := qframe.New(map[string]interface{}{
		"COL1": []int{1, 2, 3},
		"COL2": []float64{1.1, 2.2, 3.3},
		"COL3": []string{"one", "two", "three"},
		"COL4": []bool{true, true, true},
	})
	fmt.Println("qftest:", qf) */

	// Start a new SQL Transaction.
	tx, _ := sqldb.Begin()

	// Create a new QFrame from SQL.
	newQf := qframe.ReadSQL(tx,
		// A query must return at least one column. In this
		// case it will return all of the columns we created above.
		qsql.Query("SELECT * FROM bank_repay_plan"),
		// SQLite stores boolean values as integers, so we
		// can coerce them back to bools with the CoercePair option.
		// qsql.Coerce(qsql.CoercePair{Column: "ID", Type: qsql.Int64ToBool}),
		qsql.Postgres(),
	)
	fmt.Println("newQF:", newQf)
	//fmt.Println(newQf.Equals(qf))
}
