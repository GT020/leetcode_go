package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	var sum int = 0
	var carry int = 0

	for l2 != nil || l1 != nil || carry > 0 {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		newNode := ListNode{Val: sum % 10,
			Next: nil}
		temp.Next = &newNode
		temp = temp.Next

	}

	return result.Next
}
