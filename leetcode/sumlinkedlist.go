package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
ATTEMPT ONE
The numbers are in reverse order.
342 + 465 = 807		|	 [2, 4, 3] + [5, 6, 4] => [7, 0, 8]
- There may be different lengths
  - Handle this should handle nil input
- We need to maintain a carry
- We need to create the sum node
  - We'll add nodes to it, but we need to return the head
- We need to handle case where two lists are exhausted, but carry is still 1
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{Val: -1, Next: nil}
	sum := dummyHead
	carry := 0

	// 1) orchestrate looping the two lists in sync
	for l1 != nil || l2 != nil || carry > 0 {

		// 2) Sum two nodes, create new node for sum + carry; handle nil nodes
		sum, carry = add(sum, l1, l2, carry)

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	return dummyHead.Next
}

// Adds the values l1 and l2 together into a new node, and attaches it to the sum's next.
// If the new value is greater than 10, the value will be the 10's digit (e.g. val % 10,
// and the carry value will be all other digits.
//
// Returns the new node and the carry value.
func add(sum *ListNode, l1 *ListNode, l2 *ListNode, carry int) (*ListNode, int) {
	newVal := carry

	if l1 != nil {
		newVal += l1.Val
	}

	if l2 != nil {
		newVal += l2.Val
	}

	nextNode := &ListNode{Val: newVal % 10, Next: nil}
	sum.Next = nextNode

	return nextNode, newVal / 10
}
