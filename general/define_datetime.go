/*
File: define_datetime.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-30 10:49:31

Description: 日期/时间相关
*/

package general

import "time"

// GetDateTime 按照给定格式返回当前日期和时间
//
// 参数：
//   - format: 期望的时间日期返回格式
//
// 返回：
//   - 符合格式要求的当前日期和时间
func GetDateTime(format string) string {
	return time.Now().Format(format)
}
