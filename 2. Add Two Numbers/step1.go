/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	curNode := node
	carryover := 0

	for l1 != nil || l2 != nil || carryover > 0 {
		sum := carryover

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carryover = sum / 10

		curNode.Next = &ListNode{
			Val: sum % 10,
		}
		curNode = curNode.Next
	}

	return node.Next
}