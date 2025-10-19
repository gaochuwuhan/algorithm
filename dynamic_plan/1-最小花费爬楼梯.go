package dynamic_plan

/*
描述
给定一个整数数组
c
o
s
t

cost  ，其中
c
o
s
t
[
i
]

cost[i]  是从楼梯第
i

i 个台阶向上爬需要支付的费用，下标从0开始。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

数据范围：数组长度满足
1
≤
n
≤
1
0
5

1≤n≤10
5
   ，数组中的值满足
1
≤
c
o
s
t
i
≤
1
0
4

1≤cost
i
​
 ≤10
4

示例1
输入：
[2,5,20]
复制
返回值：
5
复制
说明：
你将从下标为1的台阶开始，支付5 ，向上爬两个台阶，到达楼梯顶部。总花费为5
示例2
输入：
[1,100,1,1,1,90,1,1,80,1]
复制
返回值：
6
复制
说明：
你将从下标为 0 的台阶开始。
1.支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
2.支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
3.支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
4.支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
5.支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
6.支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6 。
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param cost int整型一维数组
 * @return int整型
 */
func minCostClimbingStairs(cost []int) int {
	// write code here
	// 递推公式

	// n=len(cost)-1
	// dp[0]=0
	// dp[1]=0
	// dp[2]=dp[0]+cost[n-1] / dp[1]
	//dp[3] = dp[2]+cost[2] / dp[1] + cost[3]

	sets := make([]int, len(cost)+1, len(cost)+1) // 并非走到数组最后索引是楼顶，还得加1，就是求sets[len(cost)]对应的值 花费
	sets[0] = 0                                   //走到第0和1索引的台阶不需要花费
	sets[1] = 0
	for i := 2; i <= len(cost); i++ {
		sets[i] = min(sets[i-1]+cost[i-1], sets[i-2]+cost[i-2]) //两种走法，看上一个楼梯总数的花费的两钟走法哪个小就行
	}
	return sets[len(cost)]

}

func min(x, y int) int {
	if x >= y {
		return y
	}

	return x
}

func Jump(n int, A []int) int {

	// write code here
	if n == 1 {
		return 0
	}
	longest := 0 //下一步最远到哪
	end := 0
	times := 0
	for i := 0; i < n-1; i++ {
		if i+A[i] > longest {
			longest = i + A[i]
		}

		if i == end {
			times++
			end = longest
		}
		if longest > n-1 {
			break
		}

	}

	return times

}
