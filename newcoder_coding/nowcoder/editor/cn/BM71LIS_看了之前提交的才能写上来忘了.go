/**
给定一个长度为 n 的数组 arr，求它的最长严格上升子序列的长度。 所谓子序列，指一个数组删掉一些数（也可以不删）之后，形成的新数组。例如 [1,5,3,7,
3] 数组，其子序列有：[1,3,3]、[7] 等。但 [1,6]、[1,3,5] 则不是它的子序列。 我们定义一个序列是 严格上升 的，当且仅当该序列不存在两
个下标 和 满足 且 。
 数据范围： 要求：时间复杂度 ， 空间复杂度

 Related Topics 动态规划
示例:
输入:[6,3,1,5,2,3,7]
输出:4

*/
package nowcoder.editor.cn; //根据实际修改


import "sort"

//nowcoder submit region begin(Prohibit modification and deletion)
package main
import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 给定数组的最长严格上升子序列的长度。
 * @param arr int整型一维数组 给定的数组
 * @return int整型
*/
func LIS( arr []int ) int {
    // write code here
    if len(arr)<=1{
        return len(arr)
    }
    //[6,3,1,5,2,3,7]  从后面开始找base case，dp[6]=1 dp[5]=dp[6]+1 (如果arr[5]<arr[6])  否则 dp[5]=1
    //设定dp保存从下标到最后的最大上升子序列数
    dp:=make([]int,len(arr))

    //base case
    dp[len(arr)-1] = 1

    //从倒数第二个数向后找，就是找比自己大的数的dp值，就可以+1了
    for i:=len(arr)-2;i>=0;i--{
        //由于自己本身算一个元素所以初始化1
        dp[i] = 1
        for j:=i+1;j<len(arr);j++{
            //先后找每发现一个比arr[i]大的数，就看看他的dp值，和之前比过的是不是大，大了就更新到dp中每次保存最大的
            if arr[j]>arr[i]{
                dp[i] = Max(dp[i],dp[j]+1)
            }
        }

    }
    //把找到的所有集合排序
    sort.Ints(dp)
    return dp[len(dp)-1]

}

func Max(x,y int)int{
    if x>y{
        return x
    }
    return y
}
//nowcoder submit region end(Prohibit modification and deletion)
