/*
File: define_shared.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-30 14:29:16

Description: trash 整体会用到的函数、结构体等
*/

package cli

import (
	"fmt"
	"time"

	"github.com/yhyj/trash/general"
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
	if !general.FileExist(general.TrashFilePath) {
		if err := general.CreateDir(general.TrashFilePath); err != nil {
			fmt.Printf(general.ErrorSuffixFormat, "Error creating trash folder", ": ", err)
		}
	} else if !general.FileExist(general.TrashinfoFilePath) {
		if err := general.CreateDir(general.TrashinfoFilePath); err != nil {
			fmt.Printf(general.ErrorSuffixFormat, "Error creating trash folder", ": ", err)
		}
	}
}
