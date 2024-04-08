/*
File: define_math.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 10:49:25

Description: 处理数学计算
*/

package general

import "math"

// CountDigits 计算数字的位数
//
// 参数：
//   - number: 数字
//
// 返回：
//   - 位数
func CountDigits(number int) int {
	if number == 0 {
		return 1
	}

	// 使用log10函数计算位数
	digits := int(math.Log10(math.Abs(float64(number)))) + 1

	return digits
}
