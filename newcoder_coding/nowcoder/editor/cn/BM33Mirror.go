/**
操作给定的二叉树，将其变换为源二叉树的镜像。 数据范围：二叉树的节点数 ， 二叉树每个节点的值 要求： 空间复杂度 。本题也有原地操作，即空间复杂度 的解法，时
间复杂度 
 比如： 源二叉树
 镜像二叉树 
 
 Related Topics 树 
示例:
输入:{8,6,10,5,7,9,11}
输出:{8,10,6,11,9,7,5}

*/
package nowcoder.editor.cn;  //根据实际修改
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
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param pRoot TreeNode类 
 * @return TreeNode类
*/
func Mirror( pRoot *TreeNode ) *TreeNode {
    // write code here
    if pRoot==nil{
        return nil
    }
    pRoot.Left,pRoot.Right=pRoot.Right,pRoot.Left
    Mirror(pRoot.Left)
    Mirror(pRoot.Right)
    // 下面写法麻烦
    //exchange(pRoot)
    return pRoot
}

func exchange(root *TreeNode) {
    if root.Left!=nil && root.Right!=nil{
        root.Right,root.Left = root.Left,root.Right
        exchange(root.Left)
        exchange(root.Right)
        return
    }
    if root.Left==nil && root.Right!=nil{
        root.Right,root.Left = root.Left,root.Right
        exchange(root.Left)
        return
    }
    if root.Left!=nil && root.Right==nil{
        root.Right,root.Left = root.Left,root.Right
        exchange(root.Right)
        return
    }
    //{8,#,7,#,6,#,5,#,4}
}
//nowcoder submit region end(Prohibit modification and deletion)
