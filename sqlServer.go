package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func connectToSQLServer(serverName string, port string, user string, password string, database string) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", serverName, user, password, port, database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	// Ping the SQL Server to ensure connectivity
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQL Server")
	return db, nil
}

func queryData(db *sql.DB, table string) ([]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columnNames, err := rows.Columns()
	if err != nil {
		// Handle error
	}

	columns := make([]interface{}, len(columnNames))
	columnPointers := make([]interface{}, len(columnNames))

	var data []interface{}

	for i := range columns {
		columnPointers[i] = &columns[i]
	}

	for rows.Next() {
		// Create a slice of interfaces with the length of the number of columns

		// Scan the row into the column pointers

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		data = append(data, columnPointers)
	}

	return data, nil
}
