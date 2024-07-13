package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

type TemplateData struct {
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

func generateCRUD(data TemplateData) {
	tmpl, err := template.ParseFiles("codegen/crud_template.go.tmpl")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	output := buf.String()
	fileName := fmt.Sprintf("%s_crud.go", data.StructName)
	err = os.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("Generated %s successfully.\n", fileName)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run generate_code.go crud <schema_package> <schema_name>")
		return
	}

	command := os.Args[1]
	schemaPackage := os.Args[2]
	schemaName := os.Args[3]

	if command != "crud" {
		fmt.Printf("Unknown command: %s\n", command)
		return
	}

	schemaFile := fmt.Sprintf("%s/schema.go", schemaPackage)
	cmd := exec.Command("go", "run", schemaFile, schemaName)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running schema file: %v\n", err)
		return
	}

	var templateData TemplateData
	err = json.Unmarshal(output, &templateData)
	if err != nil {
		fmt.Printf("Error parsing schema output: %v\n", err)
		return
	}

	generateCRUD(templateData)
}
