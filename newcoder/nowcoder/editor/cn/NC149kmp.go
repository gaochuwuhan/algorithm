/**
给你一个文本串 T ，一个非空模板串 S ，问 S 在 T 中出现了多少次
 
 数据范围： 要求：空间复杂度 ，时间复杂度 
 
 Related Topics 字符串 
示例:
输入:"ababab","abababab"
输出:2 

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算模板串S在文本串T中出现了多少次
 * @param S string字符串 模板串
 * @param T string字符串 文本串
 * @return int整型
*/
func kmp1( S string ,  T string ) int {
	//暴力穷举 时间复杂度较高
	if len(T)<len(S){
		return 0
	}
	sl:=len(S)
	count:=0
	for i:=range T{
		if S[0]!=T[i]{ //o(n)
			continue
		}
		if len(T)-i>=sl{
			subT:=T[i:i+sl]
			if subT==S{
				count++
			}
		}
	}
	return count
}
func kmp( S string ,  T string ) int {
    // write code here
	if len(T)<len(S){
		return 0
	}
	times:=0
	j:=0
	matchedLen:=0
	for i:=0;i<len(T);i++{
		if j>len(S)-1 || T[i]!=S[j]{
			break
		}
		if j==len(S)-1{
			j=0
			matchedLen=0
		}

		if S[j] != T[i]{
			j=0
			matchedLen=0
			continue
		}else{
			matchedLen++
			if matchedLen==len(S){
				times++
			}
		}
	}

	return times
}
//nowcoder submit region end(Prohibit modification and deletion)
