package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type TemplateData struct {
	PackageName string
	StructName  string
	EntityName  string
}

func generateCRUD(structName, entityName, packageName string) {
	data := TemplateData{
		PackageName: packageName,
		StructName:  structName,
		EntityName:  entityName,
	}

	tmpl, err := template.ParseFiles("crud_template.go.tmpl")
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
	fileName := fmt.Sprintf("%s_crud.go", structName)
	err = os.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("Generated %s successfully.\n", fileName)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run generate_code.go <struct_name> <entity_name> <package_name>")
		return
	}

	structName := os.Args[1]
	entityName := os.Args[2]
	packageName := os.Args[3]

	generateCRUD(structName, entityName, packageName)
}
