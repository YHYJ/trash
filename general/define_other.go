/*
File: define_other.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 16:28:45

Description: 处理一些杂事
*/

package general

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

// Confirm 风险操作二次确认
//
// 参数：
//   - message: 提示信息
//   - flag: 代表确认（返回 true ）的标识
//
// 返回：
//   - 确认返回 true，否则返回 false
func Confirm(message, flag string) bool {
	color.Print(DangerText(message))

	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))

	return answer == flag
}

// UserFace 输出提示信息，获取用户输入
//
// 参数：
//   - message: 提示信息
//
// 返回：
//   - 用户输入
func UserFace(message string) []int {
	color.Print(DangerText(message))

	// 创建一个 STDIN Scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// 按行获取用户输入
	inputLine := scanner.Text()
	// 将输入行按空格分割
	numbersStr := strings.Fields(inputLine)

	var inputNumbers []int
	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr) // 将字符串转换为整数
		if err != nil {
			color.Error.Println("Please enter as required")
			return inputNumbers
		}
		inputNumbers = append(inputNumbers, num)
	}
	return inputNumbers
}

// StringSliceEqual 比较两个字符串切片是否一样
//
// 参数：
//   - slice1: 切片1
//   - slice2: 切片2
//
// 返回：
//   - 一样返回 true，一样返回 false
func StringSliceEqual(slice1, slice2 []string) bool {
	// 如果切片长度不同，则不相等
	if len(slice1) != len(slice2) {
		return false
	}
	// 逐个比较切片元素
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
