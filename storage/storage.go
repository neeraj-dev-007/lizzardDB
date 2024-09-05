package storage

import "fmt"

type Column struct {
	Name string
	Type string
}

type Table struct {
	Name string
	Columns []Column
	Rows []map[string]interface{} //interface{} because value can be of any type
}

type Database struct {
	Tables map[string]*Table
}

func NewDatabase() *Database {
	return &Database{Tables: make(map[string]*Table)}
}

func (db *Database) Create(name string, columns []Column) {
	db.Tables[name] = &Table{
		Name: name, Columns: columns, Rows: []map[string]interface{}{},
	}
	fmt.Println("Table", name, "created")
}

func (db *Database) Insert(name string, row map[string]interface{}) {
	table, exists := db.Tables[name]
	if (exists) {
		table.Rows = append(table.Rows, row)
		fmt.Println("Inserted row into", name)
	} else {
		fmt.Println("Table", name, "doesn't exist")
	}
}

func (db *Database) Select(name string) {
	table, exists := db.Tables[name]
	if (exists) {
		fmt.Println("Rows in", name)
		for _, row := range table.Rows {
			fmt.Println(row)
		}
	} else {
		fmt.Println("Table", name, "doesn't exist")
	}
}