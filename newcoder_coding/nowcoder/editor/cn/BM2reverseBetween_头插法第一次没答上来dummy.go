/**
将一个节点数为 size 链表 m ;位置到 ;n 位置之间的区间反转，要求时间复杂度 ，空间复杂度 。 例如： 给出的链表为 , , 返回 . 数据范围： ;
链表长度 ，，链表中每个节点的值满足 要求：时间复杂度 ; ，空间复杂度 进阶：时间复杂度 ，空间复杂度
 Related Topics 链表
示例:
输入:{1,2,3,4,5},2,4
输出:{1,4,3,2,5}

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"
import . "nc_tools"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
*/
func reverseBetween( head *ListNode ,  m int ,  n int ) *ListNode {
    // write code here
	//解法1： 找到要反转的区间两端及两端的前一个和后一个，找的时候反转中间的
	if m==n{
		return head
	}
	//
	//var startPre,start,end,endAfter *ListNode
	//cur:=head
	//var pre *ListNode
	//for i:=0;i<=n;i++{
	//	if cur==nil{
	//		break
	//	}
	//	if i==m-1{
	//		startPre = pre
	//		start = cur
	//	}
	//	next:=cur.Next
	//	if i<n&&i>m-1{
	//		cur.Next = pre
	//	}
	//	pre = cur
	//	cur = next
	//	if i==n-1{
	//		//找到了第n个
	//		end = pre
	//		endAfter = cur
	//		break
	//	}
	//
	//}
	//if startPre!=nil{
	//	startPre.Next=end
	//	start.Next = endAfter
	//
	//}else{
	//	//说明从第一个开始就要反转了，startPre为空，那么head指向end
	//	head = end
	//	start.Next=endAfter
	//}
	//return head

	//解法2：使用虚拟头节点（头插法）
	dummy:=&ListNode{
		Val:  0,
		Next: head,
	}
	preStart:=dummy //关键点
	for i:=1;i<m;i++{
		preStart = preStart.Next //i=1时prestart就是head 符合
	}
	// 区间反转
	var pre *ListNode
	cur:=preStart.Next
	for i:=m;i<=n;i++{
		next:=cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	start:=preStart.Next
	preStart.Next = pre
	start.Next = cur
	return dummy.Next


}
//nowcoder submit region end(Prohibit modification and deletion)
