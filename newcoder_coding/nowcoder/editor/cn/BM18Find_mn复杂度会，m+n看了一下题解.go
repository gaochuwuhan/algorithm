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
//方法一：时间复杂度较高O(MN)
//func Find( target int ,  array [][]int ) bool {
//    // write code here
//    if len(array)==0{
//        return false
//    }
//    //高
//    m:=len(array)
//    if m==0{
//        return false
//    }
//    //长
//    n:=len(array[0])
//
//
//    for i:=0;i<m;i++{
//        for j:=0;j<n;j++{
//
//            if array[i][j] == target{
//                return true
//            }
//        }
//
//
//    }
//    return false
//}
//方案二，使用横纵两个指针，不断根据数组横向纵向递增特性向右向上移动
func Find( target int ,  array [][]int ) bool {
    // write code here
    if len(array)==0{
        return false
    }
    //高
    m:=len(array)
    if m==0{
        return false
    }
    //长
    n:=len(array[0])

    left:=0
    down:=m-1
    for left<n&&down>=0{
        data:=array[down][left]
        if data==target{
            return true
        }

        if data < target{ //最下面的小于target就不能纵向看了 横向移动
            left++
        }else{
            down-- // 大于target继续向上看，横向可以不动
        }
    }
    return false
}
//nowcoder submit region end(Prohibit modification and deletion)
