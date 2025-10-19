//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。 
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。 
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。 
//
// 
//
// 示例 1： 
// 
// 
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
// 
//
// 示例 2： 
//
// 
//输入：l1 = [0], l2 = [0]
//输出：[0]
// 
//
// 示例 3： 
//
// 
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
// 
//
// 
//
// 提示： 
//
// 
// 每个链表中的节点数在范围 [1, 100] 内 
// 0 <= Node.val <= 9 
// 题目数据保证列表表示的数字不含前导零 
// 
//
// Related Topics 递归 链表 数学 👍 11514 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	var preHead = &ListNode{}
	curNode:=preHead
	nextValPlus:=false
	for l1!=nil || l2!=nil{
		
	}
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode
	//注意题干，逆序指的是个位从指针的head开始
	arr1:=make([]int,0)
	arr2:=make([]int,0)
	putToArr(l1,&arr1)
	putToArr(l2,&arr2)
	var preHead = &ListNode{}
	curNode:=preHead
	arrLen := max(len(arr1),len(arr2))
	nextValPlus:=false
	for i:=0;i<=arrLen-1;i++{
		var plusVal int
		plusVal=popData(arr1,i)+popData(arr2,i)
		previousPlus := nextValPlus
		if previousPlus{
			plusVal=plusVal+1
		}
		var nodeVal int
		if plusVal>=10{
			nextValPlus = true
			nodeVal = plusVal-10
		}else{
			nextValPlus = false
			nodeVal = plusVal
		}

		curNode.Next = &ListNode{Val:nodeVal}
		curNode = curNode.Next
	}
	if nextValPlus{
	curNode.Next = &ListNode{Val:1}
	}
	return preHead.Next

}
func putToArr(node *ListNode, arr *[]int){
	var curNode *ListNode
	curNode = node
	for curNode!=nil {
		*arr = append(*arr, curNode.Val)
		curNode = curNode.Next
	}
	return
}

func popData(arr []int,idx int) int{
	if len(arr)-1<idx{
		return 0
	}
	return arr[idx]
}
//leetcode submit region end(Prohibit modification and deletion)
