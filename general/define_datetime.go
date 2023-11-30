/*
File: define_datetime.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-30 10:49:31

Description: 日期/时间相关
*/

package general

import (
	"fmt"
	"time"
)

// GetDateTime 按照指定格式返回当前日期和时间
//
// 参数：
//   - format: 期望的时间日期返回格式
//
// 返回：
//   - 符合格式要求的当前日期和时间的字符串
func GetDateTime(format string) string {
	return time.Now().Format(format)
}

// ParseDateTime 按照指定格式将字符串解析为日期/时间
//
// 参数：
//   - format: 解析时间日期的格式
//   - datetimeStr: 待解析的日期时间字符串
//
// 返回：
//   - 解析后的日期字符串
//   - 解析后的时间字符串
//   - 错误信息
func ParseDateTime(format, datetimeStr string) (string, string, error) {
	// 解析时间字符串
	parsedTime, err := time.Parse(format, datetimeStr)
	if err != nil {
		fmt.Println("解析时间错误:", err)
		return "", "", err
	}

	// 格式化日期和时间部分
	datePart := parsedTime.Format("2006-01-02")
	timePart := parsedTime.Format("15:04:05")

	return datePart, timePart, nil
}
