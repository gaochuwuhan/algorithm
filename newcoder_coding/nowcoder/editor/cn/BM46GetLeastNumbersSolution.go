/**
给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3
,4(任意顺序皆可)。 数据范围：，数组中每个数的大小 要求：空间复杂度 ，时间复杂度 
 Related Topics 堆 排序 分治 
示例:
输入:[4,5,1,6,2,7,3,8],4 
输出:[1,2,3,4]

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param input int整型一维数组 
 * @param k int整型 
 * @return int整型一维数组
*/
func GetLeastNumbers_Solution( input []int ,  k int ) []int {
    // write code here
	if k==0||len(input)==0{
		return nil
	}
	splitNum(input,0,len(input)-1)
	return input[0:k]
}

func splitNum(num []int, l,r int) {
	//mid:=(l+r)/2
	//if mid<=l || mid>=r{
	//	return
	//} 这个判断不对，如果是0，1 mid结果还是0，但合法，唯一不行的就是只有一个元素的时候
	if l==r{
		return
	}
	mid:=(l+r)/2
	//对左边继续分割
	splitNum(num,l,mid)
	splitNum(num,mid+1,r)
	mergeTwoSorted(num,l,mid,r)

}
//[4,5,6] [1,3,5]
func mergeTwoSorted(num []int, l,mid,r int){
	//比较两个排序的数组，l为起点，mid为两个数组的重点，r为右端点
	//使用一个新数组保存排序后的，最后复制到num对应位置
	sortedRes:=make([]int,0)
	li:=l
	ri:=mid+1
	for li<=mid && ri<=r{
		if num[li]<=num[ri]{
			sortedRes = append(sortedRes,num[li])
			li++
			continue
		}
		sortedRes = append(sortedRes,num[ri])
		ri++
	}
	//检查剩余未比较元素
	if li<=mid{
		for i:=li;i<=mid;i++{
			sortedRes = append(sortedRes,num[i])
		}
	}
	if ri<=r{
		for i:=ri;i<=r;i++{
			sortedRes = append(sortedRes,num[i])
		}
	}
	//复制排序后的到原数组
	copy(num[l:r+1],sortedRes)
}
//nowcoder submit region end(Prohibit modification and deletion)
