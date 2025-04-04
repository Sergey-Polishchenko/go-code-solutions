# Go Code Solutions

![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue)
![Benchmarks](https://img.shields.io/badge/Benchmarks-Auto_updated-yellow)
![License](https://img.shields.io/badge/License-MIT-green)

## Description

### What is this?

This repository contains my solutions to algorithmic problems from platforms like LeetCode and Codewars. 

It serves as:
- Personal learning journal
- Code practice space
- Benchmarking playground
- Community knowledge sharing

### Key Features

- **Structured Solutions**: Problem descriptions + multiple solution approaches
- **Performance Tracking**: Automated benchmark statistics
- **Task Automation**: Simplified workflow with Taskfile
- **Testing**: Comprehensive test cases for each problem
- **CI/CD**: Auto-updated benchmarks on push

---

## Getting Started

### Prerequisites

- Go 1.24+
- [Task](https://taskfile.dev/) (optional but recommended)

### Usage

```bash
# Install Task
brew install go-task/tap/go-task  # macOS
choco install go-task             # Windows
sudo snap install task --classic  # Linux

# Common commands:
task new leetcode "15. 3Sum" medium  # Create new problem template
task gen-bench                       # Generates local benchmarks results
task test-all                        # Run all tests
task bench-all                       # Run all benchmarks
task test ./leetcode/1-Two_Sum       # Run tests for specific problem
task bench ./leetcode/1-Two_Sum      # Run benchmarks for specific problem iterations
task lint                            # Run linter (golangci-lint required)
task list                            # Show all solutions

# Manual commands:
go test -v ./leetcode/...  # Run tests
go run ./scripts/bench     # Generate benchmarks
```

---

## Problems

### Litcode

| № | Problem | Difficulty | Tags | Solution | Benchmark |
|---|:-----:|:--------:|:--:|------|:------:|
| 1 | [Two Sum](./leetcode/1-Two_Sum) | Easy | `Arrays`, `Hash Tables` | [Code](./leetcode/1-Two_Sum/solution.go) | [Benchmark](./leetcode/1-Two_Sum/iterations/bench.md) |
| 2 | [Add Two Numbers](./leetcode/2-Add_Two_Numbers/) | Medium | `Linked List` | [Code](./leetcode/2-Add_Two_Numbers/solution.go) | [Benchmark](./leetcode/2-Add_Two_Numbers/iterations/bench.md) |
| 3 | [Longest Substring Without Repeating Characters](./leetcode/3-Longest_Substring_Without_Repeating_Characters/) | Medium | `Strings` | [Code](./leetcode/3-Longest_Substring_Without_Repeating_Characters/solution.go) | [Benchmark](./leetcode/3-Longest_Substring_Without_Repeating_Characters/iterations/bench.md) |

### Codewars

| № | Problem | Difficulty | Tags | Solution | Benchmark |
|---|:-----:|:--------:|:--:|------|:------:|
| 1 | [Maximum subarray sum](./codewars/5kyu-Maximum_subarray_sum/) | 5 kyu | `Arrays` | [Code](./codewars/5kyu-Maximum_subarray_sum/solution.go) | [Benchmark](./codewars/5kyu-Maximum_subarray_sum/iterations/bench.md) |

### Tags

```text
Arrays, Hash Tables, Graphs, Linked Lists, Strings, Sort
```

---

## Problem Structure

```text
/platform/N-Problem_Name
├── problem.md              # Problem description
├── solution.go             # Best solution
├── solution_test.go        # Tests
├── testdata/               # Problem test data package
│   └── testcases.go        # Problem test cases
└── iterations/             # Solution iterations
    ├── v1.go               # 1 Version 
    ├── vN.go               # N Versions
    ├── iterations_test.go  # Iterations test and benchmarks
    └── bench.md            # Iterations benchmark report
```

---

## TODO:

- [ ] Add leetcode/codewars scraper for automatic creation directory for solution by URL.
- [ ] Codecov configuration for test coverage monitoring.
- [ ] Add tests for scripts.
- [ ] Make better templates.
- [ ] Readme generator for new problems.

---

## License

MIT License - see [LICENSE](./LICENSE) for details.
