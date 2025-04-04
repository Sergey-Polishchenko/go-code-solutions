package main

type BenchResult struct {
	Name     string
	Runs     string
	NsPerOp  string
	MemAlloc string
	MemOps   string
}

type BenchData struct {
	Task         string
	LastRun      string
	GoVersion    string
	CPU          string
	BenchResults map[string][]BenchResult
}

func (d BenchData) AlgorithmsCount() int {
	algorithms := make(map[string]struct{})
	for _, results := range d.BenchResults {
		for _, res := range results {
			algorithms[res.Name] = struct{}{}
		}
	}
	return len(algorithms)
}
