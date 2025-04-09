package directions_reduction_iterations

var opposite = map[string]string{
	"SOUTH": "NORTH",
	"NORTH": "SOUTH",
	"EAST":  "WEST",
	"WEST":  "EAST",
}

// Simple solution for Directions Reduction
// Dificulty: O(n)
func SolutionV1(arr []string) []string {
	var res []string
	for _, dir := range arr {
		if len(res) > 0 && opposite[dir] == res[len(res)-1] {
			res = res[:len(res)-1]
			continue
		}
		res = append(res, dir)
	}
	return res
}
