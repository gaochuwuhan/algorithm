/**
请根据二叉树的前序遍历，中序遍历恢复二叉树，并打印出二叉树的右视图 
 数据范围： 
 要求： 空间复杂度 ，时间复杂度 
 
 如输入[1,2,4,5,3],[4,2,5,1,3]时，通过前序遍历的结果[1,2,4,5,3]和中序遍历的结果[4,2,5,1,3]可重建出以下二叉树： 
 所以对应的输出为[1,3,5]。 
 Related Topics 树 
示例:
输入:[1,2,4,5,3],[4,2,5,1,3]
输出:[1,3,5] 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 求二叉树的右视图
 * @param preOrder int整型一维数组 先序遍历
 * @param inOrder int整型一维数组 中序遍历
 * @return int整型一维数组
*/
func solve( preOrder []int ,  inOrder []int ) []int {
    // write code here
    tree:=recoverBT(preOrder,inOrder)
    //使用层序遍历，最后一个出队列的就是最右侧的树
    q:=make([]*TreeNode,0)
    res:=make([]int,0)
    q=append(q,tree)
    tempQ:=make([]*TreeNode,0)
    for len(q) != 0{
        deTree:=dequeueQ(&q)
        if deTree.Left!=nil{
            tempQ = append(tempQ,deTree.Left)
        }
        if deTree.Right!=nil{
            tempQ = append(tempQ,deTree.Right)
        }
        if len(q)==0{
            res = append(res,deTree.Val)
            //入队列
            for _,tree1:=range tempQ{
                q = append(q,tree1)
            }
            tempQ = []*TreeNode{}
        }

    }
    return res

}

func dequeueQ(queue *[]*TreeNode) *TreeNode{
    if len(*queue)>0{
        data:=(*queue)[0]
        *queue = (*queue)[1:len(*queue)]
        return data
    }
    return nil
}

func recoverBT(preOrder []int, inOrder []int) *TreeNode{
    if len(preOrder) == 0{
        return nil
    }
    root:=&TreeNode{Val: preOrder[0]}
    //通过中序找到左子树和右子树
    var midI int
    for k,v:=range inOrder{
        if v==preOrder[0]{
            midI=k
            break
        }
    }

    leftPre:=preOrder[1:1+midI]
    leftTree:=recoverBT(leftPre,inOrder[0:midI])
    root.Left=leftTree
    rightPre:=preOrder[1+midI:]
    rightTree:=recoverBT(rightPre,inOrder[midI+1:])
    root.Right=rightTree

    return root


}

//nowcoder submit region end(Prohibit modification and deletion)
