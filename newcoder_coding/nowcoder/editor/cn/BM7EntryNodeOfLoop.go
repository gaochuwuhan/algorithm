/**
给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。 
 数据范围： ， 要求：空间复杂度 ，时间复杂度 
 例如，输入{1,2},{3,4,5}时，对应的环形链表如下图所示： 
 可以看到环的入口结点的结点值为3，所以返回结点值为3的结点。
 
 Related Topics 链表 哈希 双指针 
示例:
输入:{1,2},{3,4,5}
输出:3

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main

func EntryNodeOfLoop(pHead *ListNode) *ListNode{
    if pHead==nil||pHead.Next==nil{
        return nil
    }
    slow,fast,thirdNode:=pHead,pHead,pHead
    for slow!=nil && fast!=nil && fast.Next!=nil{
        slow = slow.Next
        fast = fast.Next.Next
        if fast==slow{
            break
        }
    }
    if fast!=slow{
        return nil
    }
    cycNode:=fast
    cycCount:=0

    //二次转圈，用thirdNode从头开始，fast始终在环转圈
    for {
        if cycNode==thirdNode{
            return thirdNode
        }
        cycNode = cycNode.Next
        cycCount++
        if cycCount>0 && cycNode==fast{
            thirdNode = thirdNode.Next
            cycCount=0
        }


    }

}
//nowcoder submit region end(Prohibit modification and deletion)
