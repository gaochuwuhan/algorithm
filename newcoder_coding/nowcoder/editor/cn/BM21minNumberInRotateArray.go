/**
有一个长度为 n 的非降序数组，比如[1,2,3,4,5]，将它进行旋转，即把一个数组最开始的若干个元素搬到数组的末尾，变成一个旋转数组，比如变成了[3,4,5
,1,2]，或者[4,5,1,2,3]这样的。请问，给定这样一个旋转数组，求数组中的最小值。
 
 数据范围：，数组中任意元素的值: 要求：空间复杂度： ，时间复杂度： 
 Related Topics 二分 
示例:
输入:[3,4,5,1,2]
输出:1

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
func minNumberInRotateArray( nums []int ) int {
    // write code here
    //[5,0,0,2,2]
    //[2,3,4,5,1] //从旋转的数量超过一半和未超过一般来试探，发现num[mid]>num[r]时最小值肯定在mid右边，将l移动到mid+1
    //当num[mid]<=num[r]时,不能确定，需要让mid往左移动，也就是让r往左移，r--
    l:=0

    r:=len(nums)-1
    for l<r{
        mid:=(l+r)/2
        if nums[mid] > nums[r]{
            l=mid+1
            continue
        }
        if nums[mid] <= nums[r]{
            r--
        }
    }
    return nums[l]
}

//nowcoder submit region end(Prohibit modification and deletion)
