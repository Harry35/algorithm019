package Week_03

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reversePrint(head *ListNode) []int {
	var cur *ListNode
	prv := head
	for prv != nil {
		tmp := prv.Next
		prv.Next = cur
		cur = prv
		prv = tmp
	}
	res := []int{}
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}
	