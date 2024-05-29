/*
File: define_recycle.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-08 12:22:27

Description: 操作回收站
*/

package general

import (
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

// CheckRecycleBin 检查回收站是否存在
func CheckRecycleBin() {
	if !FileExist(TrashFilePath) {
		if err := CreateDir(TrashFilePath); err != nil {
			color.Danger.Printf("Error creating trash folder: %s\n", err)
		}
	} else if !FileExist(TrashInfoPath) {
		if err := CreateDir(TrashInfoPath); err != nil {
			color.Danger.Printf("Error creating trash folder: %s\n", err)
		}
	}
}
