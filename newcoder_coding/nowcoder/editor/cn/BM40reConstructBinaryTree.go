/**
给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4
,7,2,1,5,3,8,6}，则重建出如下图所示。 
 
 提示: 1.vin.length == pre.length 2.pre 和 vin 均无重复元素 3.vin出现的元素均出现在 pre里 4.只需要返回根结
点，系统会自动输出整颗树做答案对比 数据范围：，节点的值 要求：空间复杂度 ，时间复杂度 
 Related Topics 树 dfs 数组 
示例:
输入:[1,2,4,7,3,5,6,8],[4,7,2,1,5,3,8,6]
输出:{1,2,3,4,#,5,6,#,7,#,#,8}

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
 * @param preOrder int整型一维数组 
 * @param vinOrder int整型一维数组 
 * @return TreeNode类
*/
func reConstructBinaryTree( preOrder []int ,  vinOrder []int ) *TreeNode {
    // write code here
    if len(preOrder) == 0{
        return nil
    }
    root:=&TreeNode{Val: preOrder[0]}
    //找到当前root所在中序的位置，也就得到他的左子树list和右子树list
    iroot:=0
    for ik,iv:=range vinOrder{
        if iv==root.Val{
            iroot=ik
            break
        }
    }
    lPre:=preOrder[1:1+iroot]
    rPre:=preOrder[1+iroot:]
    root.Left = reConstructBinaryTree(lPre,vinOrder[0:iroot])
    root.Right = reConstructBinaryTree(rPre,vinOrder[iroot+1:])
    return root

}
//func findLeftAndRight(root *TreeNode, preList []int, inOrderList []int){
//    rootIdx:=getRootIdx(root.Val,inOrderList)
//    leftChildren:=inOrderList[0:rootIdx]
//    rightChildren:=inOrderList[rootIdx+1:]
//    leftIdxInPre:=findPreIndexs(leftChildren,preList)
//    if len(leftIdxInPre) >0{
//        idx:=leftIdxInPre[0]
//        root.Left = &TreeNode{Val: preList[idx]}
//    }
//    rightIdxInPre:=findPreIndexs(rightChildren,preList)//元素在前序列表中的索引升序返回
//    if len(rightIdxInPre) > 0{
//        idx:=rightIdxInPre[0]
//        root.Right = &TreeNode{Val: preList[idx]}
//    }
//    lPreList,rPreList:=[]int{},[]int{}
//    if len(leftIdxInPre)>0{
//        lPreEndIdx:=len(leftIdxInPre)-1
//        lPreList = preList[leftIdxInPre[0]:leftIdxInPre[lPreEndIdx]+1]
//    }
//
//    if len(rightIdxInPre) >0{
//        rPreEndIdx:=len(rightIdxInPre)-1
//        rPreList = preList[rightIdxInPre[0]:rightIdxInPre[rPreEndIdx]+1]
//    }
//    if root.Left!=nil{
//        findLeftAndRight(root.Left,lPreList,leftChildren)
//    }
//    if root.Right!=nil{
//        findLeftAndRight(root.Right,rPreList,rightChildren)
//    }
//
//}

//nowcoder submit region end(Prohibit modification and deletion)
