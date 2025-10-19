/**
输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组，子数组最小长度为1。求所有子数组的和的最大值。 数据范围: 
 
 
 要求:时间复杂度为 ，空间复杂度为 进阶:时间复杂度为 ，空间复杂度为 
 
 Related Topics 动态规划 贪心 
示例:
输入:[1,-2,3,10,-4,7,2,-5]
输出:18

*/
package nowcoder




.editor.cn; //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main


//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param array int整型一维数组 
 * @return int整型
*/
// 连续子数组的和与他的子数组的和有关，举例：索引3结尾的和=索引2结尾的和+索引3元素的值，sum(3)=sum(2)+arr[3]，sum(2)=sum(1)+arr[2]
//由此可以推导公式为 dp[i]=max(dp[i-1]+arr[i],arr[i]) dp[i-1]如果为负数，还不如不加，所以和当前元素去max值
//base case: 当i=0时，和为数组索引元素dp[0]=arr[0]
func Max(x,y int)int{
    if x>y{
        return x
    }
    return y
}

func FindGreatestSumOfSubArray( array []int ) int {

    // write code here
    dp:=make([]int,len(array))
    dp[0]=array[0]
    maxSum:=dp[0]

    for i:=1;i<len(array);i++{
        dp[i]=Max(dp[i-1]+array[i],array[i])
        maxSum = Max(maxSum,dp[i])
    }

    return maxSum
}
//nowcoder submit region end(Prohibit modification and deletion)
