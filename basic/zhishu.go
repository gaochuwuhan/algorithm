package basic

import "fmt"

// 找出0-100之间的质数
func FindZhiShu() {
	zs := []int{2}
	for i := 3; i <= 10000; i++ {
		zl := len(zs)
		for j := 0; j < zl; j++ {
			if i%zs[j] == 0 {
				break //这个数能整除质数说明不是质数
			}
			if zl-1 == j {
				zs = append(zs, i) //排除到最后一个代表不能整除所有当前质数，作为结果
			}
		}
	}
	fmt.Println(zs)
}
