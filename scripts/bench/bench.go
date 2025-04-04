package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/Sergey-Polishchenko/go-code-solutions/internal/utils"
)

func main() {
	problemDirs, err := findProblemDirs("leetcode", "codewars")
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range problemDirs {
		if err := generateBenchmarks(dir); err != nil {
			log.Fatal(err)
		}
	}
}

func generateBenchmarks(dir string) error {
	benchDir := fmt.Sprintf("./%s/iterations", dir)

	cmd := exec.Command(
		"go",
		"test",
		"-bench=.",
		"-benchmem",
		"-run=^$",
		benchDir,
	)
	out, _ := cmd.CombinedOutput()
	fmt.Print(string(out))

	data := parseBenchOutput(string(out), dir)
	if len(data.BenchResults) == 0 {
		return fmt.Errorf("empty bench results")
	}

	outputPath := filepath.Join(benchDir, "bench.md")
	data.GoVersion = runtime.Version()
	data.LastRun = time.Now().Format("2006-01-02 15:04:05")

	templateName := "bench.md.tmpl"
	if data.AlgorithmsCount() == 1 {
		templateName = "benchDS.md.tmpl"
	}

	if err := utils.RenderTemplate(outputPath, templateName, data); err != nil {
		return fmt.Errorf("render error: %v", err)
	}

	fmt.Printf("Generated: %s\n\n", outputPath)
	return nil
}

func parseBenchOutput(output, dir string) BenchData {
	data := BenchData{
		Task:         strings.Split(filepath.Base(dir), "-")[1],
		BenchResults: make(map[string][]BenchResult),
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "cpu: ") {
			data.CPU = strings.TrimSpace(strings.TrimPrefix(line, "cpu: "))
		}

		if strings.HasPrefix(line, "Benchmark") {
			parts := strings.Fields(line)
			if len(parts) < 7 {
				continue
			}

			fullName := strings.TrimPrefix(parts[0], "Benchmark")
			nameParts := strings.Split(fullName, "/")

			size, name := "Medium", "Unknown"
			if len(nameParts) >= 3 {
				size = nameParts[1]
				name = strings.Join(nameParts[2:], " ")
			} else if len(nameParts) == 2 {
				name = nameParts[1]
			}

			result := BenchResult{
				Name:     strings.Split(name, "-")[0],
				Runs:     parts[1],
				NsPerOp:  parts[2],
				MemAlloc: parts[4],
				MemOps:   parts[6],
			}

			data.BenchResults[size] = append(data.BenchResults[size], result)
		}
	}
	return data
}

func findProblemDirs(platforms ...string) ([]string, error) {
	var dirs []string
	for _, platform := range platforms {
		if _, err := os.Stat(platform); os.IsNotExist(err) {
			fmt.Printf("Platform directory %s does not exist, skipping\n", platform)
			continue
		}
		err := filepath.Walk(
			platform,
			func(path string, info os.FileInfo, err error) error {
				if info.IsDir() && hasSolutionFiles(path) {
					dirs = append(dirs, path)
				}
				return nil
			},
		)
		if err != nil {
			return nil, fmt.Errorf("can't walk on platforms: %v", err)
		}
	}
	return dirs, nil
}

func hasSolutionFiles(path string) bool {
	for _, f := range []string{"solution.go", "problem.md"} {
		if _, err := os.Stat(filepath.Join(path, f)); os.IsNotExist(err) {
			return false
		}
	}
	return true
}
