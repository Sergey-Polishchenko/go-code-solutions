package testdata

var Tests = []struct {
	Name string
	L1   *ListNode
	L2   *ListNode
	Want *ListNode
}{
	{
		Name: "test case 1",
		L1:   NewListNode(2, 4, 3),
		L2:   NewListNode(5, 6, 4),
		Want: NewListNode(7, 0, 8),
	},
	{
		Name: "test case 2",
		L1:   NewListNode(0),
		L2:   NewListNode(0),
		Want: NewListNode(0),
	},
	{
		Name: "test case 3",
		L1:   NewListNode(9, 9, 9, 9, 9, 9, 9),
		L2:   NewListNode(9, 9, 9, 9),
		Want: NewListNode(8, 9, 9, 9, 0, 0, 0, 1),
	},
}
