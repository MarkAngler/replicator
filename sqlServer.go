package main

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// ConnectToDatabase establishes a new database connection
func connectToDatabase(serverName, port, username, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=your_db", username, password, serverName, port)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	return db, err
}

// ReadFromTable reads all rows from a table and populates the rows slice with them.
func readFromTable(db *gorm.DB, tableName string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	result := db.Find(&tableName).Scan(&results)
	if result.Error != nil {
		return nil, result.Error
	}
	return results, nil
}

func getAllTables(db *gorm.DB) ([]map[string]interface{}, error) {
	result, err := readFromTable(db, "INFORMATION_SCHEMA.TABLES")
	return result, err
}

// InsertToTable inserts a new row into a table. Generic function, less type safe.
func insertToTable(db *gorm.DB, tableName string, values interface{}) error {
	result := db.Table(tableName).Create(values)
	return result.Error
}
