/**
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。 1.对于该题的最近的公共祖先定义:对于有根树T的两个节点p、q，最近公共祖先LCA(T,p,q)表
示一个节点x，满足x是p和q的祖先且x的深度尽可能大。在这里，一个节点也可以是它自己的祖先. 2.二叉搜索树是若它的左子树不空，则左子树上所有节点的值均小于它的
根节点的值； 若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值 3.所有节点的值都是唯一的。 4.p、q 为不同节点且均存在于给定的二叉搜索树中。
 数据范围: 3<=节点总数<=10000 0<=节点值<=10000 
 如果给定以下搜索二叉树: {7,1,12,0,4,11,14,#,#,3,5}，如下图: 
 
 
 Related Topics 树 递归 
示例:
输入:{7,1,12,0,4,11,14,#,#,3,5},1,12
输出:7 

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
 * @param p int整型 
 * @param q int整型 
 * @return int整型
*/
func lowestCommonAncestor( root *TreeNode ,  p int ,  q int ) int {
    // write code here
    ret:=root.Val
    if p<ret && q<ret{
        return lowestCommonAncestor(root.Left,p,q)
    }
    if p>ret && q>ret{
        return lowestCommonAncestor(root.Right,p,q)
    }
    return ret
}
//nowcoder submit region end(Prohibit modification and deletion)
