//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}
	var newRes=&ListNode{}
	tmpNode:=newRes
	for list1!=nil && list2!=nil{
		if list1.Val<list2.Val{
			tmpNode.Next=list1
			list1=list1.Next
		}else{
			tmpNode.Next=list2
			list2=list2.Next
		}
		tmpNode=tmpNode.Next
	}
	if list1 == nil{
		tmpNode.Next=list2
	}else{
		tmpNode.Next=list1
	}
	return newRes.Next
}
//leetcode submit region end(Prohibit modification and deletion)
