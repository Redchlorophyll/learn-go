package learn_golang_db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_db")

	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlInsertCustomer := "INSERT INTO customer(id, name) VALUES('DAHS', 'DAHS')"
	_, err := db.ExecContext(ctx, sqlInsertCustomer)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}
