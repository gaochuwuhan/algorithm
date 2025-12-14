/**
在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维
数组和一个整数，判断数组中是否含有该整数。 [ [1,2,8,9], [2,4,9,12], [4,7,10,13], [6,8,11,15] ] 给定 
target ;= 7，返回 ;true。 给定 ;target ;= ;3，返回 ;false。 数据范围：矩阵的长宽满足 ， 矩阵中的值满足 进阶：空间复杂度 ，
时间复杂度 
 Related Topics 数组 
示例:
输入:7,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
输出:true

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main
//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param target int整型 
 * @param array int整型二维数组 
 * @return bool布尔型
*/
func Find( target int ,  array [][]int ) bool {
    // write code here
    if len(array)==0 || len(array[0])==0{
        return false
    }
    downI:=len(array)-1
    leftI:=0
    for downI>=0 && leftI < len(array[0]){
        if array[downI][leftI] == target{
            return true
        }
        if array[downI][leftI] < target{
            leftI++
            continue
        }
        if array[downI][leftI]>leftI{
            downI--
        }
    }

    return false

}
//nowcoder submit region end(Prohibit modification and deletion)
