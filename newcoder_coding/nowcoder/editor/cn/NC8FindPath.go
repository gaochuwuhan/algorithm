/**
输入一颗二叉树的根节点root和一个整数expectNumber，找出二叉树中结点值的和为expectNumber的所有路径。 1.该题路径定义为从树的根结点开
始往下一直到叶子结点所经过的结点 2.叶子节点是指没有子节点的节点 3.路径只能从父节点到子节点，不能从子节点到父节点 4.总节点数目为n 
 如二叉树root为{10,5,12,4,7},expectNumber为22 则合法路径有[[10,5,7],[10,12]]
 
 数据范围: 树中节点总数在范围 [0, 5000] 内 -1000 <= 节点值 <= 1000 -1000 <= expectNumber <= 1000 

 Related Topics 树 
示例:
输入:{10,5,12,4,7},22
输出:[[10,5,7],[10,12]] 

*/
package nowcoder

import "fmt"

.editor.cn; //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
import "fmt"
import . "nc_tools"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

//type TreeNode struct {
//   Val int
//   Left *TreeNode
//   Right *TreeNode
//}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param root TreeNode类 
 * @param target int整型 
 * @return int整型二维数组
*/

func walk(tree *TreeNode, res *[][]int,curSum int,curSlice []int, target int){
    if curSlice == nil{
        curSlice = make([]int,0)
    }
    fmt.Println(curSlice)
    if tree.Left == nil && tree.Right == nil{
        if curSum+tree.Val == target{
            curSlice = append(curSlice,tree.Val)
            *res = append(*res,curSlice)
            return
        }
    }
    curSlice = append(curSlice,tree.Val)
    curSum = curSum+tree.Val
    if tree.Left!=nil{
        inputS:=make([]int,len(curSlice))
        copy(inputS,curSlice)
        walk(tree.Left,res,curSum,inputS,target)
    }
    if tree.Right !=nil{
        inputS:=make([]int,len(curSlice))
        copy(inputS,curSlice)
        walk(tree.Right, res, curSum,inputS,target)
    }
}

func FindPath( root *TreeNode ,  target int ) [][]int {
    // write code here
    if root == nil{
        return nil
    }
    res:=make([][]int,0)
    walk(root,&res,0,nil,target)
    return res
}
//nowcoder submit region end(Prohibit modification and deletion)
