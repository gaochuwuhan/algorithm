/**
给定一个长度为 n 的数组 num 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。 
 例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2
,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6
],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。 
 窗口大于数组长度或窗口长度为0的时候，返回空。
 
 数据范围： ，，数组中每个元素的值满足 要求：空间复杂度 ，时间复杂度 
 
 
 Related Topics 堆 双指针 队列 
示例:
输入:[2,3,4,2,6,2,5,1],3
输出:[4,4,6,6,6,5]

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param num int整型一维数组 
 * @param size int整型 
 * @return int整型一维数组
*/
func maxInWindows( num []int ,  size int ) []int {
    // write code here
    if size==0 || size>len(num){
        return nil
    }

    //使用递增队列，存储从出口到入口为递增的数列
    res:=make([]int,0)
    iq:=&IncQueue{data: make([]int, 0)}

    //初始化把第一个窗口最大值放进递增队列里
    for i:=0;i<len(num);i++{
        iq.Push(num[i])
        if i>=size-1{
            if i>=size{
                popData:=num[i-size]
                iq.Pop(popData)
            }
            maxVal:=iq.GetMax()
            res = append(res,maxVal)
        }

    }

    return res

}

type IncQueue struct {
    data []int

}

func (i *IncQueue) Push(node int) {
    //入队列先和头部比较，若>=头部直接替换整个队列，只放这个最大的；
    //若<头部则 持续和尾部最后元素比较，如果大于直接替换，直到替换到头部
    if len(i.data) == 0{
        i.data = append(i.data, node)

        return
    }
    if node >= i.data[0]{
        i.data = []int{node}
        return
    }
    // 找到比node大的idx,从尾部开始比较，右比node小于等于的就去掉
    for i.data[len(i.data)-1] < node{
        i.data = i.data[0:len(i.data)-1]
    }
    i.data = append(i.data, node)

}

func (i *IncQueue) Pop(val int) {
    //从头部pop数据
    if len(i.data)>0{
        if val==i.data[0]{
            i.data = i.data[1:]
        }
    }
}

func (i *IncQueue) GetMax() int{
    if len(i.data) >0{
        return i.data[0]
    }
    return 1
}
//nowcoder submit region end(Prohibit modification and deletion)
