/*
File: define_disk.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-01-30 11:35:38

Description: 磁盘、挂载点相关函数
*/

package general

import (
	"github.com/jaypipes/ghw"
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

// InThisMountpoint 判断文件是否在指定的挂载点中
//
// 参数：
//   - filePath: 文件路径
//   - mountpoints: 挂载点切片
//
// 返回：
//   - 文件在指定的挂载点中返回 true，否则返回 false
func InThisMountpoint(filepath string, mountpoints []string) bool {
	for _, mountpoint := range mountpoints {
		if mountpoint == "/home" {
			return true
		}
	}
	return false
}

// PartitionInfo 结构体存储分区信息
type PartitionInfo struct {
	Device    string // 设备名
	Mount     string // 挂载点
	Removable bool   // 设备是否可移除
}

// GetPartitionInfo 获取分区信息
//
// - 包括设备名、挂载点和设备是否可移除
//
// 返回：
//   - 分区信息和错误信息
func GetPartitionInfo() ([]PartitionInfo, error) {
	// 获取所有分区信息

	// 创建一个切片以存储分区信息
	var partitionInfo []PartitionInfo

	// 使用 ghw 获取设备名和可移除状态
	block, err := ghw.Block()
	if err != nil {
		return nil, err
	}
	for _, disk := range block.Disks {
		for _, partition := range disk.Partitions {
			partitionInfo = append(partitionInfo, PartitionInfo{
				Device:    disk.Name,
				Mount:     partition.MountPoint,
				Removable: disk.IsRemovable,
			})
		}
	}

	return partitionInfo, nil
}
