package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

// OK is struct healthcheck
type OK struct {
	Status  int
	Message string
}

func query() {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/Project")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	rows, err := db.Query("SELECT * FROM users")
	columns, err := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	// fmt.Print(values)
	// fmt.Print("rows")
}
func main() {
	e := echo.New()
	e.GET("/", healthcheck)
	listenPort := ":5000"
	e.Logger.Fatal(e.Start(listenPort))
	fmt.Println("listen on : http://localhost" + listenPort)
}

func healthcheck(c echo.Context) error {
	query()

	m := OK{200, "Hello"}
	return c.JSON(http.StatusOK, m)
}
