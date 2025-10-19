/**
你有一个背包，最多能容纳的体积是V。
 现在有n种物品，每种物品有任意多个，第i种物品的体积为 ,价值为。
 （1）求这个背包至多能装多大价值的物品？ （2）若背包恰好装满，求至多能装多大价值的物品？
 数据范围：

 Related Topics 动态规划
示例:
输入:6,2,[[5,10],[3,1]]
输出:[10,2]

*/
//nowcoder submit region begin(Prohibit modification and deletion)
package main

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param v int整型
 * @param n int整型
 * @param nums int整型二维数组
 * @return int整型一维数组
 */

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func knapsack(v int, n int, nums [][]int) []int {
	// write code here
	//dp为求放前i个物品，容量达到j时的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, v+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= v; j++ {
			dp[i][j] = dp[i-1][j] // 不取当前物品
			if j >= nums[i-1][0] {
				dp[i][j] = Max(dp[i][j], dp[i][j-nums[i-1][0]]+nums[i-1][1])
			}
		}
	}
	maxVal := dp[n][v]

	//重新初始dp，记录装满的情况，未装满则默认置为-1
	dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, v+1)
		for j := range dp[i] {
			if i == 0 && j == 0 { //todo 不太懂这里为什么只有ij都是0而不是或条件
				dp[i][j] = 0
			} else {
				dp[i][j] = -1
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= v; j++ {
			dp[i][j] = dp[i-1][j] // 不取当前物品
			if j >= nums[i-1][0] && dp[i][j-nums[i-1][0]] != -1 {
				dp[i][j] = Max(dp[i][j], dp[i][j-nums[i-1][0]]+nums[i-1][1])
			}
		}
	}
	fullVal := dp[n][v]
	fullVal = Max(fullVal, 0)

	return []int{maxVal, fullVal}
} //nowcoder submit region end(Prohibit modification and deletion)
