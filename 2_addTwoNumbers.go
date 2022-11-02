package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	flag := false
	var res *ListNode
	var headRes *ListNode
	for {
		if l1 == nil && l2 == nil && !flag {
			return headRes
		}
		val1 := getValue(l1)
		val2 := getValue(l2)
		valRes := val1 + val2
		if flag {
			valRes++
		}
		if valRes >= 10 {
			flag = true
			valRes = valRes - 10
		} else {
			flag = false
		}
		if res == nil {
			res = &ListNode{
				Val:  valRes,
				Next: nil,
			}
			headRes = res
		} else {
			res.Next = &ListNode{
				Val:  valRes,
				Next: nil,
			}
			res = res.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
}

func getValue(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}
