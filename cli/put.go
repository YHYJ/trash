/*
File: put.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-30 09:44:01

Description:
*/

package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yhyj/trash/general"
)

// PutFile 将文件移动到回收站
//
// 参数：
//   - files: 需要移动的文件列表
func PutFile(files []string) {
	for _, file := range files {
		if general.FileExist(file) {
			absPath := general.GetAbsPath(file)                                      // 待删除文件的绝对路径
			trashedFileName := filepath.Base(file)                                   // 回收站文件名
			trashedFilePath := filepath.Join(general.TrashFilePath, trashedFileName) // 回收站文件的路径

			// 检测回收站中 trashPath 是否已存在，存在则为 trashPath 增加一个累加的后缀
			for num := 1; ; num++ {
				if !general.FileExist(trashedFilePath) {
					break
				}
				trashedFileName = fmt.Sprintf("%s_%d", filepath.Base(file), num)
				trashedFilePath = filepath.Join(general.TrashFilePath, trashedFileName)
			}

			// 将文件移动到回收站
			err := os.Rename(file, trashedFilePath)
			if err != nil {
				fmt.Println("Error moving to trash:", err)
			} else {
				trashinfoCreator(trashedFileName, absPath)
			}
		} else {
			fmt.Printf(general.RegelarFormat, fmt.Sprintf("trash put: cannot remove '%s': No such file or directory", file))
		}
	}
}

// trashinfoCreator 创建已删除文件信息存储文件
//
// 参数：
//   - fileName: 信息存储文件的文件名（不包含后缀名）
//   - originalPath: 已删除文件的原绝对路径
func trashinfoCreator(fileName, originalPath string) {
	// 创建已删除文件信息存储文件
	trashinfoFilePath := filepath.Join(general.TrashinfoFilePath, fmt.Sprintf("%s.trashinfo", filepath.Base(fileName)))
	general.CreateFile(trashinfoFilePath)

	// 写入已删除文件信息
	format := "2006-01-02T15:04:05"
	general.WriteFile(trashinfoFilePath, fmt.Sprintf(general.TrashinfoFileContent, originalPath, general.GetDateTime(format)))
}
