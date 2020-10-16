package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func toLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	curr := head
	for _, num := range nums {
		curr.Next = &ListNode{Val: num, Next: nil}
		curr = curr.Next
	}
	return head.Next
}

func stringifyList(list *ListNode) string {
	curr := list
	str := ""
	for curr != nil {
		str += strconv.Itoa(curr.Val)
		curr = curr.Next
	}
	return str
}

func TestWithTwoPlusTwo(t *testing.T) {
	sumList := addTwoNumbers(
		toLinkedList([]int{2}),
		toLinkedList([]int{2}))

	assert.Equal(t, stringifyList(sumList), "4")
}

func TestWhenOneListIsEmpty(t *testing.T) {
	sumListA := addTwoNumbers(
		toLinkedList([]int{}),
		toLinkedList([]int{2}))

	assert.Equal(t, stringifyList(sumListA), "2")

	sumListB := addTwoNumbers(
		toLinkedList([]int{2}),
		toLinkedList([]int{}))

	assert.Equal(t, stringifyList(sumListB), "2")
}

func TestWhenBothListsAreEmpty(t *testing.T) {
	result := addTwoNumbers(
		toLinkedList([]int{}),
		toLinkedList([]int{}))

	assert.Nil(t, result)
}

// i.e. 342 + 465 = 807
func TestTwoMultiDigitNumbersSameLength(t *testing.T) {
	sumList := addTwoNumbers(
		toLinkedList([]int{2, 4, 3}),
		toLinkedList([]int{5, 6, 4}))

	assert.Equal(t, stringifyList(sumList), "708")
}

func TestZeroPlusZero(t *testing.T) {
	sumList := addTwoNumbers(
		toLinkedList([]int{0}),
		toLinkedList([]int{0}))

	assert.Equal(t, stringifyList(sumList), "0")
}

func TestMultiDigitNumbersDifferentLengths(t *testing.T) {
	sumList := addTwoNumbers(
		toLinkedList([]int{9, 9, 9, 9, 9, 9, 9}),
		toLinkedList([]int{9, 9, 9, 9}))

	assert.Equal(t, stringifyList(sumList), "89990001")
}
