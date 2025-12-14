/**
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的
结果输出。 即输出P mod 1000000007 数据范围： ; 对于 的数据, 对于 的数据, 数组中所有数字的值满足 要求：空间复杂度 ，时间复杂度 
 Related Topics 数组 
示例:
输入:[1,2,3,4,5,6,7,0]
输出:7

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param nums int整型一维数组 
 * @return int整型
*/
func InversePairs( nums []int ) int {
    // write code here
    if len(nums) == 0{
        return 0
    }
    //1.时间复杂度o(n2)
    total:=0
    //for i:=0;i<len(nums)-1;i++{
    //
    //    for j:=i+1;j<len(nums);j++{
    //        if nums[i]>nums[j]{
    //            total++
    //        }
    //    }
    //}
    //2. 时间复杂度O(nlogn) 归并排序思想：数组等分，每组排序
    split(nums,0,len(nums)-1,&total)
    return total%1000000007
}

func split(nums []int,l,r int,total *int){
    if l>=r{
        return
    }
    mid:=(l+r)/2
    split(nums,l,mid,total)
    split(nums,mid+1,r,total)
    merge(nums,l,mid,r,total)


}

func merge(nums []int, l,mid,r int,total *int){
    tmp:=make([]int,0)
    curL,curR:=l,mid+1
    for curR<=r && curL<=mid{
        if nums[curL]>nums[curR]{
            //当前左边l-curL的都是小于右边的，所有curL-mid都是大于右边的
            *total+=mid-curL+1

            tmp = append(tmp,nums[curR])

            curR++

        }else{
            tmp = append(tmp,nums[curL])
            curL++
        }
    }
    //
    if curL<=mid{
        for _,v:=range nums[curL:mid+1]{
            tmp = append(tmp,v)
        }
    }
    if curR<=r{
        for _,v:=range nums[curR:r+1]{
            tmp = append(tmp,v)
        }
    }
    copy(nums[l:r+1],tmp)
}
//nowcoder submit region end(Prohibit modification and deletion)
