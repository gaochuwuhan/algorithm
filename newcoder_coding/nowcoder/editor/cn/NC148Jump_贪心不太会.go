/**
给你一个长度为 n 的数组 a。 ai 表示从 i 这个位置开始最多能往后跳多少格。 求从 1 开始最少需要跳几次就能到达第 n 个格子。 
 数据范围： ， 
 进阶： 空间复杂度 ， 时间复杂度 
 
 Related Topics 动态规划 
示例:
输入:2,[1,2]
输出:1 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 最少需要跳跃几次能跳到末尾
 * @param n int整型 数组A的长度
 * @param A int整型一维数组 数组A
 * @return int整型
*/
// dp[i] 为跳到第i个格子时用的最多少次数
// base case :dp[1]=0,dp[2]=1，dp[3]=min(dp[1]+1 (a[0]>=3-0-1),dp[2]+1 (a[1]>=3-2))
//dp[4]=dp[到大dp[n]的跳跃次数为到达1-n之间每种可能的次数
// dp[n]=

func Jump(n int, A []int) int {

    // write code here
    if n == 1 {
        return 0
    }
    longest := 0 //下一步最远到哪
    end:=0
    times := 0
    for i := 0; i < n-1; i++ {
        if i+A[i] > longest {
            longest = i + A[i]
        }

        if i==end{
            times++
            end=longest
            if end>n-1{
                break
            }
        }


    }

    return times

}
//nowcoder submit region end(Prohibit modification and deletion)
