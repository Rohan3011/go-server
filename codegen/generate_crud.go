package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	// Read the template file
	tmplContent, err := os.ReadFile("codegen/crud_template.go.tmpl")
	if err != nil {
		panic(err)
	}

	// Parse template file
	tmpl, err := template.New("crudTemplate").Funcs(template.FuncMap{
		"camelCase": camelCase,
		"inc":       increment,
	}).Parse(string(tmplContent))

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

// helpers
func camelCase(s string) string {
	// Remove non-alphanumeric characters
	s = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(s, "")
	// Replace underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")
	// Title case
	s = cases.Title(language.AmericanEnglish, cases.NoLower).String(s)
	// Remove spaces
	s = strings.ReplaceAll(s, " ", "")
	// Lowercase the first letter
	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}
	return s
}

func increment(n int) int {
	return n + 1
}
