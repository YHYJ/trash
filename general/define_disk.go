/*
File: define_disk.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-01-30 11:35:38

Description: 磁盘、挂载点相关函数
*/

package general

import (
	"path/filepath"
	"syscall"

	"github.com/shirou/gopsutil/v3/disk"
)

// GetMountpoints 获取系统中的挂载点
//
//   - 包括 '/tmp'
//
// 返回：
//   - 挂载点的字符串切片
func GetMountpoints() ([]string, error) {
	// 将物理设备的文件系统类型添加到预定义切片
	physicalPartitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	for _, partition := range physicalPartitions {
		fsTypes = append(fsTypes, partition.Fstype)
	}

	// 获取符合要求的挂载点，要求为：
	// - partition.Device == "tmpfs" && partition.Mountpoint == "/tmp" && partition.Fstype == "tmpfs"
	// - partition.Fstype 在 fsTypes 中
	allPartitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	mountpoints := make([]string, 0)
	for _, partition := range allPartitions {
		if StringSliceEqual([]string{partition.Device, partition.Mountpoint, partition.Fstype}, []string{"tmpfs", "/tmp", "tmpfs"}) {
			mountpoints = append(mountpoints, partition.Mountpoint)
		}
		if fsTypes.Have(partition.Fstype) {
			mountpoints = append(mountpoints, partition.Mountpoint)
		}
	}
	return mountpoints, nil
}

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
