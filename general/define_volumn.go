/*
File: define_volumn.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-01-30 11:35:38

Description: 磁盘、挂载点相关函数
*/

package general

import (
	"path/filepath"
	"syscall"
)

// IsMountPoint 测试路径是否为挂载点
//
// 参数：
//   - path: 待测试路径
//
// 返回：
//   - 是挂载点返回 true，否则返回 false 和错误信息
func IsMountPoint(path string) (bool, error) {
	// 根目录单独处理
	if path == "/" {
		return isRootMount(path)
	}

	// 获取给定路径的状态信息
	var stat syscall.Stat_t
	if err := syscall.Stat(path, &stat); err != nil {
		return false, err
	}

	// 获取父目录的路径
	parent := filepath.Dir(path)
	// 获取父目录的状态信息
	var parentStat syscall.Stat_t
	if err := syscall.Stat(parent, &parentStat); err != nil {
		return false, err
	}

	// 如果给定路径的设备号不等于父目录的设备号，则认为是挂载点
	return stat.Dev != parentStat.Dev, nil
}

// isRootMount 检查给定路径是否为根挂载点
//
// 参数：
//   - path: 待检查路径
//
// 返回：
//   - 是根挂载点返回 true，否则返回 false 和错误信息
func isRootMount(path string) (bool, error) {
	// 获取给定路径的设备号
	var stat syscall.Stat_t
	if err := syscall.Stat(path, &stat); err != nil {
		return false, err
	}

	// 获取根目录的设备号
	var rootStat syscall.Stat_t
	if err := syscall.Stat("/", &rootStat); err != nil {
		return false, err
	}

	// 如果给定路径的设备号等于根目录的设备号，则认为是根挂载点
	return stat.Dev == rootStat.Dev, nil
}
