package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Sergey-Polishchenko/go-code-solutions/internal/utils"
)

type ProblemData struct {
	PackageName        string
	PackageDir         string
	ProblemTitle       string
	Difficulty         string
	ProblemDescription string
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf(
			"Usage: %s <platform> <problem-name> [difficulty]\nExample: %s leetcode '1. Two Sum' easy",
			os.Args[0],
			os.Args[0],
		)
	}

	platform := os.Args[1]
	problemName := os.Args[2]
	difficulty := "unknown"
	if len(os.Args) > 3 {
		difficulty = os.Args[3]
	}

	platforms := []string{"leetcode", "codewars"}
	if !slices.Contains(platforms, platform) {
		log.Fatal("Problem platform must be 'leetcode' or 'codewars'")
	}

	if !strings.Contains(problemName, ". ") {
		log.Fatal("Problem name must be in format 'X. Problem Name' (e.g. '1. Two Sum')")
	}

	data := ProblemData{
		PackageName:        formatName(problemName),
		PackageDir:         formatDirName(platform, problemName),
		ProblemTitle:       problemName,
		Difficulty:         difficulty,
		ProblemDescription: "TODO: Add problem description",
	}

	dirPath := data.PackageDir
	if err := os.MkdirAll(filepath.Join(dirPath, "iterations"), 0750); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dirPath, "testdata"), 0750); err != nil {
		log.Fatal(err)
	}

	templates := map[string]string{
		"solution.go.tmpl":        filepath.Join(dirPath, "solution.go"),
		"solution_test.go.tmpl":   filepath.Join(dirPath, "solution_test.go"),
		"problem.md.tmpl":         filepath.Join(dirPath, "problem.md"),
		"testcases.go.tmpl":       filepath.Join(dirPath, "testdata", "testcases.go"),
		"iterations_test.go.tmpl": filepath.Join(dirPath, "iterations", "iterations_test.go"),
		"iteration.go.tmpl":       filepath.Join(dirPath, "iterations", "v1.go"),
	}

	for tmpl, output := range templates {
		if err := utils.RenderTemplate(output, tmpl, data); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Created problem: %s\n", dirPath)
}

func formatName(name string) string {
	s := strings.Split(name, ". ")
	return strings.ToLower(strings.ReplaceAll(s[1], " ", "_"))
}

func formatDirName(platform, name string) string {
	s := strings.Split(name, ". ")
	dirName := fmt.Sprintf("%s-%s", s[0], strings.ReplaceAll(s[1], " ", "_"))
	return filepath.Join(platform, dirName)
}
