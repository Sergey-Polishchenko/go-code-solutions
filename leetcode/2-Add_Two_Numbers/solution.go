// 2. Add Two Numbers
package add_two_numbers

import "github.com/Sergey-Polishchenko/go-code-solutions/leetcode/2-Add_Two_Numbers/testdata"

// Solution for Add Two Sum
// Difficulty: O(n)
func Solution(l1, l2 *testdata.ListNode) *testdata.ListNode {
	sumNode := &testdata.ListNode{}
	current := sumNode
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		newCarry := sum + carry
		carry = newCarry / 10
		current.Next = &testdata.ListNode{Val: newCarry % 10}
		current = current.Next
	}
	return sumNode.Next
}
