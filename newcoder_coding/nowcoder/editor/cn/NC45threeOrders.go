/**
给定一棵二叉树，分别按照二叉树先序，中序和后序打印所有的节点。 
 数据范围：，树上每个节点的val值满足 要求：空间复杂度 ，时间复杂度 样例解释： 如图二叉树结构
 
 Related Topics 栈 树 哈希 
示例:
输入:{1,2,3}
输出:[[1,2,3],[2,1,3],[2,3,1]]

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"
import . "nc_tools"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */
 //type TreeNode struct {
 //    Val int
 //    Left *TreeNode
 //    Right *TreeNode
 //    }
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
*/

func preOrder(tree *TreeNode,res *[]int) {
    if tree == nil{
        return
    }
    *res = append(*res, tree.Val)
    preOrder(tree.Left,res)
    preOrder(tree.Right,res)
}

func midOrder(tree *TreeNode,res *[]int) {
    if tree == nil{
        return
    }
    midOrder(tree.Left,res)
    *res = append(*res, tree.Val)
    midOrder(tree.Right,res)
}
func backOrder(tree *TreeNode,res *[]int) {
    if tree == nil{
        return
    }
    backOrder(tree.Left,res)
    backOrder(tree.Right,res)
    *res = append(*res, tree.Val)
}

func threeOrders( root *TreeNode ) [][]int {
    // write code here
    preRes:=make([]int,0)
    preOrder(root,&preRes)
    midRes:=make([]int,0)
    midOrder(root,&midRes)
    backRes:=make([]int,0)
    backOrder(root,&backRes)
    return [][]int{
        preRes,midRes,backRes,
    }
}
//nowcoder submit region end(Prohibit modification and deletion)
