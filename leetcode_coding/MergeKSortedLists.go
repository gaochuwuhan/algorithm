
//IMPORTANT!! Submit Code Region Begin(Do not remove this line)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var newListNode, l1curNode, l2curNode *ListNode
	l1Head := list1
	l2Head := list2
	if l1Head.Val < l2Head.Val {
		newListNode = l1Head
		l1curNode = l1Head.Next
		l2curNode = l2Head
	} else {
		newListNode = l2Head
		l2curNode = l2Head.Next
		l1curNode = l1Head
	}

	for {
		if l1curNode == nil {
			newListNode.Next = l2curNode
			break
		}
		if l2curNode == nil {
			newListNode.Next = l1curNode
			break
		}
		if l1curNode.Val < l2curNode.Val {
			newListNode.Next = l1curNode
		} else {
			newListNode.Next = l2curNode
		}
	}
	return newListNode
}
//IMPORTANT!! Submit Code Region End(Do not remove this line)