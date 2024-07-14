//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Schema struct {
	PackageName string
	StructName  string
	EntityName  string
	TableName   string
	Fields      []Field
}

type Field struct {
	Name string
	Type string
}

var todoSchema = Schema{
	PackageName: "todo",
	StructName:  "TodoStore",
	EntityName:  "Todo",
	TableName:   "todos",
	Fields: []Field{
		{Name: "ID", Type: "int"},
		{Name: "Title", Type: "string"},
		{Name: "Completed", Type: "bool"},
		{Name: "CreatedAt", Type: "time.Time"},
	},
}

var productSchema = Schema{
	PackageName: "product",
	StructName:  "ProductStore",
	EntityName:  "Product",
	TableName:   "products",
	Fields: []Field{
		{Name: "Id", Type: "int"},
		{Name: "Name", Type: "string"},
		{Name: "Price", Type: "float64"},
		{Name: "Stock", Type: "int"},
		{Name: "CreatedAt", Type: "time.Time"},
	},
}

func main() {
	schemas := map[string]Schema{
		"todo":    todoSchema,
		"product": productSchema,
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run schema.go <schema_name>")
		return
	}

	schemaName := os.Args[1]
	schema, found := schemas[schemaName]
	if !found {
		fmt.Printf("Schema not found: %s\n", schemaName)
		return
	}

	schemaJSON, err := json.Marshal(schema)
	if err != nil {
		fmt.Printf("Error marshaling schema: %v\n", err)
		return
	}

	fmt.Println(string(schemaJSON))
}
