package hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

// IsValid 20 有效的括号数
func isValid(s string) bool {
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		switch s[i : i+1] {
		case "{", "(", "[":
			stack = append(stack, s[i:i+1])
		case "}", ")", "]":
			if len(stack) == 0 {
				return false
			}
			peek := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if s[i:i+1] == "}" && peek != "{" {
				return false
			} else if s[i:i+1] == ")" && peek != "(" {
				return false
			} else if s[i:i+1] == "]" && peek != "[" {
				return false
			}
		default:
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	for list1 != nil {
		cur.Next = list1
		list1 = list1.Next
		cur = cur.Next
	}
	for list2 != nil {
		cur.Next = list2
		list2 = list2.Next
		cur = cur.Next
	}

	return dummy.Next
}
