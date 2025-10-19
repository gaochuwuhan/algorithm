/**
一个机器人在m×n大小的地图的左上角（起点）。 机器人每次可以向下或向右移动。机器人要到达地图的右下角（终点）。 可以有多少种不同的路径从起点走到终点？ 
 备注：m和n小于等于100,并保证计算结果在int范围内 
 数据范围：，保证计算结果在32位整型范围内 要求：空间复杂度 ，时间复杂度 进阶：空间复杂度 ，时间复杂度 
 Related Topics 动态规划 
示例:
输入:2,1
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
 * @param m int整型 
 * @param n int整型 
 * @return int整型
*/

// 将走到某一格子设为dp[m][n],从题干的图可以知道第一行和第一列的格子每个到达只有一种走法, 其他格子走法为上面格子的走法数+左边格子的走法数
//即dp[m][n]=dp[m][n-1]+dp[m-1][n]
//base case: dp[1][0]=1 dp[2][0]=1 dp[0][1]=1 dp[0][2]=1 ... dp[0][1~n-1]=1 dp[1~m-1][0]=1

func uniquePaths( m int ,  n int ) int {
    // write code here
	if m<1 || n<1 {
		return 0
	}
	dp:=make([][]int,n)
	for x:=range dp{
		dp[x]= make([]int, m)
	}
	for i:=0;i<m;i++{
		dp[0][i]=1
	}
	for i:=0;i<n;i++{
		dp[i][0]=1
	}
	for i:=1;i<n;i++{
		for j:=1;j<m;j++{
			dp[i][j]=dp[i-1][j]+dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}
//nowcoder submit region end(Prohibit modification and deletion)
