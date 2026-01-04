/**
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。 
 数据范围：二叉树的节点数量满足 ，二叉树节点的值满足 ，树的各节点的值各不相同 示例 1： 
 
 Related Topics 树 递归 dfs 广度优先搜索(BFS) 
示例:
输入:{1,#,2,3}
输出:[1,2,3] 

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

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param root TreeNode类 
 * @return int整型一维数组
*/
func preorderTraversal( root *TreeNode ) []int {
    // write code here
    if root==nil{
        return nil
    }
    ret:=make([]int,0)
    preOrderRecur(root,&ret)
    return ret


}

func preOrderRecur(tree *TreeNode,res *[]int){
    if tree==nil{
        return
    }
    *res = append(*res,tree.Val)
    preOrderRecur(tree.Left,res)
    preOrderRecur(tree.Right,res)

}
//nowcoder submit region end(Prohibit modification and deletion)
