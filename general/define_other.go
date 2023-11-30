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
	"fmt"
	"os"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))
	return answer == flag
}
