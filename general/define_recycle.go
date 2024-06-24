/*
File: define_recycle.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-08 12:22:27

Description: 操作回收站
*/

package general

import (
	"path/filepath"
	"time"

	"github.com/gookit/color"
)

// FileEntry 存储回收站文件信息
type FileEntry struct {
	Index        int       // 文件索引
	Time         time.Time // 文件删除时间
	OriginalPath string    // 文件原绝对路径（即删除前的路径）
	Path         string    // 文件绝对路径（即删除后的路径）
}

var (
	TrashPath       = filepath.Join(UserInfo.HomeDir, "/.local/share/Trash") // 回收站路径
	TrashFilePath   = filepath.Join(TrashPath, "files")                      // 回收站文件存储路径
	TrashInfoPath   = filepath.Join(TrashPath, "info")                       // 已删除文件的 trashinfo 文件路径
	DeviceTrashPath = func() string {                                        // 可移动设备回收站名
		otherTrash := filepath.Join(".Trash", UserInfo.Uid)
		if FileExist(otherTrash) {
			return otherTrash
		}
		return color.Sprintf(".Trash-%s", UserInfo.Uid)
	}()
	TrashInfoFileContent    = "[Trash Info]\nPath=%s\nDeletionDate=%s\n" // 已删除文件的 trashinfo 文件内容
	TrashInfoFileTimeFormat = "2006-01-02T15:04:05"                      // 记录文件删除时间的字符串格式
)

type StringSlice []string

var fsTypes = StringSlice{ // 在此切片中的文件系统类型视为物理设备（来自 https://github.com/andreafrancia/trash-cli）
	"btrfs",
	"fuse", // https://github.com/andreafrancia/trash-cli/issues/250
	"fuse.glusterfs",
	"fuse.mergerfs", // https://github.com/andreafrancia/trash-cli/issues/255
	"nfs",
	"nfs4",
	"p9", // file system used in WSL 2 (Windows Subsystem for Linux)
}

// Have 检查字符串切片中是否存在指定的字符串
//
// 参数：
//   - target: 目标字符串
//
// 返回：
//   - 存在返回 true，否则返回 false
func (s StringSlice) Have(target string) bool {
	// 遍历切片中的每个元素
	for _, item := range s {
		// 如果找到目标字符串，则返回 true
		if item == target {
			return true
		}
	}
	// 如果切片中不存在目标字符串，则返回 false
	return false
}

// CheckRecycleBin 检查回收站是否存在
func CheckRecycleBin() {
	if !FileExist(TrashFilePath) {
		if err := CreateDir(TrashFilePath); err != nil {
			fileName, lineNo := GetCallerInfo()
			color.Printf("%s %s -> Unable to create trash folder: %s\n", DangerText("Error:"), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		}
	} else if !FileExist(TrashInfoPath) {
		if err := CreateDir(TrashInfoPath); err != nil {
			fileName, lineNo := GetCallerInfo()
			color.Printf("%s %s -> Unable to create trash folder: %s\n", DangerText("Error:"), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		}
	}
}
