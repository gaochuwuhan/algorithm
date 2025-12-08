/**
合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。 
 数据范围：节点总数 ，每个节点的val满足 要求：时间复杂度 
 Related Topics 堆 链表 分治 
示例:
输入:[{1,2,3},{4,5,6,7}]
输出:{1,2,3,4,5,6,7}

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
 * @param lists ListNode类一维数组 
 * @return ListNode类
*/
func mergeKLists( lists []*ListNode ) *ListNode {
    // write code here
    if  len(lists) == 0{
        return nil
    }
    if len(lists)==1 {
        return lists[0]
    }
    head:=lists[0]
    var res *ListNode
    for i:=1;i<len(lists);i++{
        res = rerank(head,lists[i])
        head = res
    }
    return head
}

//使用两两比较的方法，复用两个链表顺序合并
func rerank(l1 *ListNode,l2 *ListNode) *ListNode{
    if l1 == nil{
        return l2
    }
    if l2==nil{
        return l1
    }
    newHead:=&ListNode{}
    cur:=newHead
    cur1:=l1
    cur2:=l2
    for {
        if cur1 == nil || cur2 == nil{
            break
        }
        if cur1.Val<cur2.Val{
            cur.Next = cur1
            cur1 = cur1.Next
        }else{
            cur.Next = cur2
            cur2 = cur2.Next
        }
        cur = cur.Next
    }
    if cur1 != nil{
        cur.Next = cur1
    }
    if cur2 != nil{
        cur.Next = cur2
    }
    return newHead.Next
}
//nowcoder submit region end(Prohibit modification and deletion)
