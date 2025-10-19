/**
已知一个背包最多能容纳体积之和为v的物品
 
 现有 n 个物品，第 i 个物品的体积为 vi , 重量为 wi
 
 求当前背包最多能装多大重量的物品? 
 数据范围： ， ， ， 
 进阶 ： 
 Related Topics 动态规划 
示例:
输入:10,2,[[1,3],[10,4]]
输出:4

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算01背包问题的结果
 * @param V int整型 背包的体积
 * @param n int整型 物品的个数
 * @param vw int整型二维数组 第一维度为n,第二维度为2的二维数组,vw[i][0],vw[i][1]分别描述i+1个物品的vi,wi
 * @return int整型
*/
// dp[i][j], 第一维度为物品纵向，第二维度为背包容量 横向， 在放或不放物品i时 j容量下的最大重量
// dp[i][j]=max(dp[i-1][j], dp[i-1][j-vw[i][0]+vw[i][1]]
//base case: j=0时背包0容量，啥也不能放 dp[0-i][0]=0, i=0时，dp[0][0-j]=0

func Max(x,y int)int{
	if x>y{
		return x
	}
	return y
}

func knapsack( V int ,  n int ,  vw [][]int ) int {
    // write code here
    //初始化二维dp
	dp:=make([][]int,n+1)// dp数组存在边界，0容量或0物品，所以纵向和和横向长度都是n+1 和V+1
	for i:=range dp{
		dp[i] = make([]int,V+1)
	}

	for i:=1;i<=n;i++{ //这里取的是base case 之后的数据，i为第i个物品，所以vw取的时候需要vw[i-1]
		for j:=1;j<=V;j++{
			if j-vw[i-1][0]>=0{ //
				dp[i][j]=Max(dp[i-1][j],dp[i-1][j-vw[i-1][0]]+vw[i-1][1])
			}else{
				dp[i][j]=dp[i-1][j]
			}
		}
	}
	return dp[n][V]

}
//nowcoder submit region end(Prohibit modification and deletion)