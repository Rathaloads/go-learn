package main

import (
	"fmt"
	"math/rand"
)

// 数组的一些操作

// 随机访问数组
func randomAccess(ary []int) int {
	idx := rand.Intn(len(ary))
	return ary[idx]
}

// 数组插入(这里实现的是切插入)
// 将数组指定为止之后的元素往后移动，然后再将数值插入到指定为止
// 注意：如果数组是固定长度的，最后一位会丢失
func sliceInsert(ary []int, ele int, index int) []int {
	if index < 0 || index > len(ary) {
		return nil
	}
	return append(ary[:index], append([]int{ele}, ary[index:]...)...)
}

// 切片删除，删除切片指定位置的元素
func sliceRemove(ary []int, start int, count int) []int {
	if start < 0 || start > len(ary) {
		return nil
	}
	if count < 0 {
		return nil
	}
	if start+count > len(ary) {
		return nil
	}
	return append(ary[:start], ary[start+count:]...)
}

// 数组
func main() {
	var num1 [5]int //创建固定长度的数组
	num1[0] = 1
	num1[1] = 2
	num1[2] = 3
	num1[3] = 4
	num1[4] = 5
	// num1[5] = 6 // 超出数组长度，数组溢出
	fmt.Printf("num1=%v \n", num1)

	// 当不指定数组的长度时候，创建出来的就是切片
	num2 := []int{0, 1, 2, 3, 4}
	// num2[6] = 5 // 超出切片长度，溢出，切片新增不能直接添加需要用到append()
	// num2 = append(num2, 5)
	num2 = sliceRemove(num2, 2, 1)
	fmt.Printf("num2=%v \n", num2)

	fmt.Printf("randomAccess=%v", randomAccess(num2))
}
