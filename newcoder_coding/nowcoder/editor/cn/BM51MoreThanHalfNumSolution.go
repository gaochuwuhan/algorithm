/**
给一个长度为 n 的数组，数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。 例如输入一个长度为9的数组[1,2,3,2,2,2,5,4,2]。由于
数字2在数组中出现了5次，超过数组长度的一半，因此输出2。 
 数据范围：，数组中元素的值 要求：空间复杂度：，时间复杂度 
 Related Topics 哈希 数组 
示例:
输入:[1,2,3,2,2,2,5,4,2]
输出:2

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param numbers int整型一维数组 
 * @return int整型
*/
func MoreThanHalfNum_Solution( numbers []int ) int {
    // write code here
    maxCount:=0
    storeInfo:=map[int]int{}
    for _,v:=range numbers{
        val,ok:=storeInfo[v]
        if !ok{
            storeInfo[v] = 1
        }else{
            storeInfo[v] = val+1
        }
        maxCount = Max(maxCount,storeInfo[v])
        if maxCount>len(numbers)/2{
            return v
        }
    }
    return 0

}
func Max(x,y int)int{
    if x>y{
        return x
    }
    return y
}
//nowcoder submit region end(Prohibit modification and deletion)
