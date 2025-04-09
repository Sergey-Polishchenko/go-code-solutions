package testdata

var Tests = []struct {
	Name string
	Arr  []string
	Want []string
}{
	{
		Name: "example 1: all opposites in sequence",
		Arr:  []string{"NORTH", "SOUTH", "EAST", "WEST"},
		Want: []string{},
	},
	{
		Name: "example 2: complex reduction",
		Arr:  []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"},
		Want: []string{"WEST"},
	},
	{
		Name: "example 3: same as example 2",
		Arr:  []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"},
		Want: []string{"WEST"},
	},
	{
		Name: "example 4: non-reducible directions",
		Arr:  []string{"NORTH", "WEST", "SOUTH", "EAST"},
		Want: []string{"NORTH", "WEST", "SOUTH", "EAST"},
	},
	{
		Name: "example 5: multiple reductions lead to single direction",
		Arr:  []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"},
		Want: []string{"NORTH"},
	},
	{
		Name: "two opposites cancel",
		Arr:  []string{"NORTH", "SOUTH"},
		Want: []string{},
	},
	{
		Name: "two same directions",
		Arr:  []string{"NORTH", "NORTH"},
		Want: []string{"NORTH", "NORTH"},
	},
	{
		Name: "single direction",
		Arr:  []string{"EAST"},
		Want: []string{"EAST"},
	},
	{
		Name: "opposites not adjacent",
		Arr:  []string{"NORTH", "EAST", "SOUTH"},
		Want: []string{"NORTH", "EAST", "SOUTH"},
	},
	{
		Name: "three elements, last two cancel",
		Arr:  []string{"EAST", "EAST", "WEST"},
		Want: []string{"EAST"},
	},
	{
		Name: "multiple pairs cancelling",
		Arr:  []string{"NORTH", "SOUTH", "WEST", "EAST"},
		Want: []string{},
	},
	{
		Name: "interleaved opposites",
		Arr:  []string{"NORTH", "SOUTH", "NORTH", "SOUTH"},
		Want: []string{},
	},
	{
		Name: "no adjacent opposites",
		Arr:  []string{"NORTH", "EAST", "SOUTH", "WEST"},
		Want: []string{"NORTH", "EAST", "SOUTH", "WEST"},
	},
	{
		Name: "sequential reductions leading to single element",
		Arr:  []string{"NORTH", "SOUTH", "EAST", "WEST", "NORTH"},
		Want: []string{"NORTH"},
	},
	{
		Name: "cancelling pairs after earlier reduction",
		Arr:  []string{"WEST", "EAST", "EAST", "WEST"},
		Want: []string{},
	},
	{
		Name: "nested reductions",
		Arr:  []string{"NORTH", "EAST", "WEST", "SOUTH"},
		Want: []string{},
	},
}
