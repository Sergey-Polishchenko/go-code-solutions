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
| 4 | [Median of Two Sorted Arrays](./leetcode/4-Median_of_Two_Sorted_Arrays/) | Hard | `Arrays`, `Bin Search` | [Code](./leetcode/4-Median_of_Two_Sorted_Arrays/solution.go) | [Benchmark](./leetcode/4-Median_of_Two_Sorted_Arrays/iterations/bench.md) |
| 5 | [Longest Palindromic Substring](./leetcode/5-Longest_Palindromic_Substring/) | Medium | `Strings` | [Code](./leetcode/5-Longest_Palindromic_Substring/solution.go) | [Benchmark](./leetcode/5-Longest_Palindromic_Substring/iterations/bench.md) |

### Codewars

| № | Problem | Difficulty | Tags | Solution | Benchmark |
|---|:-----:|:--------:|:--:|------|:------:|
| 1 | [Maximum subarray sum](./codewars/5kyu-Maximum_subarray_sum/) | 5 kyu | `Arrays` | [Code](./codewars/5kyu-Maximum_subarray_sum/solution.go) | [Benchmark](./codewars/5kyu-Maximum_subarray_sum/iterations/bench.md) |
| 2 | [Directions Reduction](./codewars/5kyu-Directions_Reduction/) | 5 kyu | `Arrays`, `Strings` | [Code](./codewars/5kyu-Directions_Reduction/solution.go) | [Benchmark](./codewars/5kyu-Directions_Reduction/iterations/bench.md) |

### Tags

```text
Arrays, Hash Tables, Graphs, Linked Lists, Strings, Sort, Bin Search
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

### 1. Automation of LeetCode/Codewars Problem Parsing

- [ ] Implement a basic parser to:
    - Extract problem number/title from the URL
    - Generate directory structure based on a template

- [ ] Integrate with LeetCode API to:
    - Auto-fill problem.md (description, examples, constraints)
    - Generate basic test cases in testdata/

- [ ] Support for Codewars:
    - Parse kata description
    - Automatically classify by difficulty (kyu level)

### 2. Test Coverage Monitoring

- [ ] Set up Codecov:
    - Add -coverprofile to test commands
    - Integrate with GitHub Actions
    - Badge in README + reports for pull requests

- [ ] Add minimum coverage requirements (85% for new solutions)

### 3. Script Testing

- [ ] Write tests for:
    - New problem generator (scripts/new)
    - Benchmark scripts (scripts/bench)
    - Template validation

- [ ] Create test scenarios:
    - File creation correctness
    - Input error handling
    - Benchmark report formatting

### 4. Template Improvements

- [ ] Universal templates for:
    - Graph problems (Node structure)
    - Dynamic programming (DP table)
    - Tree-based problems

- [ ] Context-aware comments:
    - Auto-generate complexity descriptions
    - Example usage in godoc

- [ ] Support for multi-file solutions:
    - Helper function packages
    - Alternative implementations

### 5. README Generator

- [ ] Script for:
    - Auto-updating solution table
    - Parsing metadata from problem.md
    - Sorting by number/difficulty

- [ ] Support for:
    - Filter tags
    - Difficulty statistics
    - Links to solution iterations

### 6. Solution Documentation

- [ ] For each solution, include:
    - Detailed algorithm explanation
    - Time and space complexity
    - ASCII diagram of the process
    - Links to related algorithms

- [ ] Unified comment style:
    ```text
    // Solution for Two Sum problem
    // Approach: Hash Map One-pass
    // Time: O(n), Space: O(n)
    // Optimization: Early exit on duplicate pairs
    ```

---

## License

MIT License - see [LICENSE](./LICENSE) for details.
