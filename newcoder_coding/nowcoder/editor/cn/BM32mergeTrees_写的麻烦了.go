/**
已知两颗二叉树，将它们合并成一颗二叉树。合并规则是：都存在的结点，就将结点值加起来，否则空的位置就由另一个树的结点来代替。例如：
 两颗二叉树是:
 Tree 1 
 
 Tree 2
 
 合并后的树为
 数据范围：树上节点数量满足 ，树上节点的值一定在32位整型范围内。 进阶：空间复杂度 ，时间复杂度 
 
 Related Topics 树 
示例:
输入:{1,3,2,5},{2,1,3,#,4,#,7}
输出:{3,4,5,5,4,#,7}

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
 * @param t1 TreeNode类 
 * @param t2 TreeNode类 
 * @return TreeNode类
*/
func mergeTrees( t1 *TreeNode ,  t2 *TreeNode ) *TreeNode {
    // write code here
    if t1==nil{
        return t2
    }
    if t2==nil{
        return t1
    }
    t1.Val+=t2.Val
    t1.Left=mergeTrees(t1.Left,t2.Left)
    t1.Right=mergeTrees(t1.Right,t2.Right)
    return t1

}


//nowcoder submit region end(Prohibit modification and deletion)
