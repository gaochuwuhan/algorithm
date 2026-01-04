/**
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。如下图所示 
 
 数据范围：输入二叉树的节点数 ，二叉树中每个节点的值 
 要求：空间复杂度（即在原树上操作），时间复杂度 
 
 注意: 1.要求不能创建任何新的结点，只能调整树中结点指针的指向。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继
 2.返回链表中的第一个节点的指针
 3.函数返回的TreeNode，有左右指针，其实可以看成一个双向链表的数据结构 4.你不用输出双向链表，程序会根据你的返回值自动打印输出 
 Related Topics 分治 
示例:
输入:{10,6,14,4,8,12,16}
输出:From left to right are:4,6,8,10,12,14,16;From right to left are:16,14,12,10,8
,6,4;

*/
package nowcoder.editor.cn; //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
import . "nc_tools"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 
 * @param pRootOfTree TreeNode类 
 * @return TreeNode类
*/
func Convert( pRootOfTree *TreeNode ) *TreeNode {
    // write code here
    if pRootOfTree==nil{
        return nil
    }
    midOrderListTree:=make([]*TreeNode,0)
    midOrder1(pRootOfTree,&midOrderListTree)
    //使用中序遍历存节点的方式，存完遍历依次连接
    for i:=0;i<len(midOrderListTree)-1;i++{
        midOrderListTree[i].Right=midOrderListTree[i+1]
        midOrderListTree[i+1].Left=midOrderListTree[i]
    }

    return midOrderListTree[0]


}

func midOrder1(tree *TreeNode,res *[]*TreeNode) {
    if tree==nil{
        return
    }
    midOrder1(tree.Left,res)
    *res = append(*res,tree)

    midOrder1(tree.Right,res)
}



//nowcoder submit region end(Prohibit modification and deletion)

