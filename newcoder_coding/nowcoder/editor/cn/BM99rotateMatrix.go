/**
有一个NxN整数矩阵，请编写一个算法，将矩阵顺时针旋转90度。 给定一个NxN的矩阵，和矩阵的阶数N,请返回旋转后的NxN矩阵。 
 数据范围：，矩阵中的值满足 
 要求：空间复杂度 ，时间复杂度 进阶：空间复杂度 ，时间复杂度 
 Related Topics 数组 基础数学 
示例:
输入:[[1,2,3],[4,5,6],[7,8,9]],3
输出:[[7,4,1],[8,5,2],[9,6,3]]

*/
package nowcoder.editor.cn;  //根据实际修改
//nowcoder submit region begin(Prohibit modification and deletion)
package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param mat int整型二维数组 
 * @param n int整型 
 * @return int整型二维数组
*/
func rotateMatrix( mat [][]int ,  n int ) [][]int {
    // write code here
    //定义新的matrix
    output:=make([][]int,n)
    for i:=0;i<n;i++{
        output[i] = make([]int,n)
    }
    for i:=0;i<n;i++{
        rotateBoarder(mat,i,n,&output)
    }

    return output
}

func rotateBoarder(oriMatrix [][]int, x int, n int,newMatrix *[][]int) {
    //顶边
    fromX:=x
    endX:=n-1-x
    if endX-fromX<0{
        return
    }
    if endX==fromX{
        (*newMatrix)[fromX][endX] = oriMatrix[fromX][endX]
        return
    }
    topLine:=oriMatrix[x][fromX:endX+1]
    for i:=fromX;i<=endX;i++{        //顶边旋转
        (*newMatrix)[i][endX]= topLine[i-fromX]
    }
    // 右边旋转
    rightLine:=make([]int,endX-fromX+1)
    for i:=fromX;i<=endX;i++{
        j:=fromX
        border:=oriMatrix[i][endX]
        rightLine[i-j] = border
    }



    for i,_:=range rightLine{
        (*newMatrix)[endX][i] = rightLine[endX-i]
    }
    //底部旋转
    bottomLine:=oriMatrix[endX][fromX:endX+1]
    for i:=fromX;i<=endX;i++{
        (*newMatrix)[i][fromX]= bottomLine[i-fromX]

    }

    //左边旋转
    leftLine:=make([]int,endX-fromX+1)
    for i:=fromX;i<=endX;i++{
        leftLine[i-fromX] = oriMatrix[i][fromX]
    }
    for i:=range leftLine{
        (*newMatrix)[fromX][i]=leftLine[endX-i]
    }

}
//nowcoder submit region end(Prohibit modification and deletion)
