/*
File: define_datetime.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-30 10:49:31

Description: 处理日期/时间
*/

package general

import (
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
//   - time.Parse 解析结果
//   - 错误信息
func ParseDateTime(format, datetimeStr string) (time.Time, error) {
	// 解析时间字符串
	parsedTime, err := time.Parse(format, datetimeStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
