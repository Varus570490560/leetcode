package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else {
		var pHead, curNode *ListNode
		if list1.Val < list2.Val {
			pHead = list1
			list1 = list1.Next
		} else {
			pHead = list2
			list2 = list2.Next
		}
		curNode = pHead
		for list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				curNode.Next = list1
				list1 = list1.Next
			} else {
				curNode.Next = list2
				list2 = list2.Next
			}
			curNode = curNode.Next
		}
		if list1 == nil {
			curNode.Next = list2
		} else {
			curNode.Next = list1
		}
		return pHead
	}
}
