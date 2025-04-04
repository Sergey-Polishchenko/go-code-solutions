package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func RenderTemplate(outputPath, tmplName string, data any) error {
	tmplPath := filepath.Join("templates", tmplName)
	tmpl, err := template.New(tmplName).
		Funcs(
			template.FuncMap{
				"formatNumber": formatNumber,
			},
		).
		ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", tmplName, err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outputPath, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file %s: %v", outputPath, err)
		}
	}()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("template execution failed: %w", err)
	}
	return nil
}

func formatNumber(n string) string {
	if len(n) > 3 {
		return n[:len(n)-3] + "," + n[len(n)-3:]
	}
	return n
}
