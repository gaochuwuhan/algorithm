/**
给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替） 
 数据范围：,树上每个节点的val满足 
 要求：空间复杂度：，时间复杂度： 例如：
 给定的二叉树是{1,2,3,#,#,4,5}
 
 该二叉树之字形层序遍历的结果是 [ [1], [3,2], [4,5] ] 
 Related Topics 栈 树 队列 
示例:
输入:{1,2,3,#,#,4,5}
输出:[[1],[3,2],[4,5]]

*/
package nowcoder.editor.cn; //根据实际修改



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
 * @param pRoot TreeNode类 
 * @return int整型二维数组
*/
func Print( pRoot *TreeNode ) [][]int {
    // write code here
	if pRoot==nil{
		return nil
	}
	ret:=make([][]int,0)
	queue:=NewQueue()
	queue.Enqueue(pRoot)
	flag:=0
	for !queue.Isempty(){
		levelRes:=make([]int,0)
		enNodes:=make([]*TreeNode,0)
		deNodes:=make([]*TreeNode,0)
		for {
			if queue.Isempty(){
				break
			}
			//继续从队列出
			node:=queue.Dequeue()
			deNodes = append(deNodes,node)
			if node.Left!=nil{
				enNodes = append(enNodes,node.Left)
			}
			if node.Right !=nil{
				enNodes = append(enNodes,node.Right)
			}

		}


		for _,enNode:= range enNodes{
		    queue.Enqueue(enNode)
		}
		if flag%2==1{
			//反转出队列节点顺序
			reverseSlice(&deNodes)
		}
		for _,node:=range deNodes{
			levelRes = append(levelRes,node.Val)
		}
		if len(levelRes)>0{
			ret = append(ret, levelRes)
		}
		flag++
	}

	return ret

}

type Queue struct {
	queue []*TreeNode
}

func NewQueue() *Queue{
	return &Queue{
		queue: make([]*TreeNode, 0),
	}
}

func (q *Queue) Enqueue(tree *TreeNode) {
	q.queue = append(q.queue,tree)
}

func (q *Queue) Dequeue() *TreeNode{
	if !q.Isempty(){
		tree:=q.queue[0]
		l:=len(q.queue)
		cur:=q.queue[1:l]
		q.queue = cur
		return tree
	}
	return nil
}

func (q *Queue) Isempty() bool{
	return len(q.queue)==0
}

func reverseSlice(input *[]*TreeNode){
	l:=0
	r:=len(*input)-1
	for l<r{
		(*input)[l],(*input)[r] = (*input)[r],(*input)[l]
		r--
		l++
	}
}

//nowcoder submit region end(Prohibit modification and deletion)
