package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := new(ListNode)
	left, right, cur := l1, l2, dummyNode
	carry := 0
	for left != nil || right != nil {
		leftVal, rightVal := 0, 0
		if left != nil {
			leftVal = left.Val
		}
		if right != nil {
			rightVal = right.Val
		}
		sum := leftVal + rightVal + carry
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
		if left != nil {
			left = left.Next
		}
		if right != nil {
			right = right.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return dummyNode.Next
}
