/**
给出一个整型数组 numbers 和一个目标值 target，请在数组中找出两个加起来等于目标值的数的下标，返回的下标按升序排列。 （注：返回的数组下标从1开始
算起，保证target一定可以由数组里面2个数字相加得到） 
 数据范围：，， 要求：空间复杂度 ，时间复杂度 
 Related Topics 数组 哈希 
示例:
输入:[3,2,4],6
输出:[2,3]

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
 * @param target int整型 
 * @return int整型一维数组
*/
func twoSum( numbers []int ,  target int ) []int {
    // write code here
    tgtMap:=map[int]int{}
    for i:=0;i<len(numbers);i++{
       subVal:=target-numbers[i]
       idx,ok:=tgtMap[subVal]
       if ok{
           return []int{idx,i+1}
       }
       tgtMap[numbers[i]] = i+1
    }
    return nil
}
//nowcoder submit region end(Prohibit modification and deletion)

