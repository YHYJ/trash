/*
File: put.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 09:44:01

Description: 子命令 'put' 的实现
*/

package cli

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/yhyj/trash/general"
)

// PutFiles 将文件移动到回收站
//
//  1. [文件是否存在]                   -- 检测目标文件列表中的文件是否存在，存在的文件继续处理，不存在的文件报错文件不存在
//  2. [文件和用户回收站是否在同一分区] -- 判断目标文件和用户回收站是否在同一分区，是则将文件移动到回收站，不是则继续
//  3. [文件是否在可移动设备上]         -- 判断目标文件所在设备是否可移动，可移动则将文件移动到设备回收站，不可移动则报错跨文件系统操作
//
// 参数：
//   - files: 需要移动的文件列表
func PutFiles(files []string) {
	for _, file := range files {
		// 判断文件是否存在
		if general.FileExist(file) {
			// 待删除文件
			absPath := general.GetAbsPath(file)       // 待删除文件的绝对路径
			filename := general.GetFilePureName(file) // 待删除文件名
			// 回收站文件
			fileNameInTrash := filename // 文件在回收站的名字

			// 获取分区信息
			partitionInfo, err := general.GetPartitionInfo()
			if err != nil {
				color.Info.Tips(err.Error())
			}

			// 判断文件和用户回收站是否在同一分区
			if strings.HasPrefix(absPath, "/home") {
				filePathInTrash := filepath.Join(general.TrashFilePath, fileNameInTrash) // 文件在回收站的路径
				// 创建回收站
				if err := general.CreateDir(general.TrashFilePath); err != nil {
					fileName, lineNo := general.GetCallerInfo()
					color.Printf("%s %s -> Unable to create trash folder: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				}
				// 若回收站中存在同名文件，则为待删除文件名字后增加一个累加的数字后缀
				for num := 1; ; num++ {
					if !general.FileExist(filePathInTrash) {
						break
					}
					fileNameInTrash = color.Sprintf("%s_%d", filename, num)
					filePathInTrash = filepath.Join(general.TrashFilePath, fileNameInTrash)
				}
				// 将文件移动到回收站
				if err = os.Rename(file, filePathInTrash); err != nil {
					fileName, lineNo := general.GetCallerInfo()
					color.Printf("%s %s -> Unable to move to trash: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
				} else {
					trashinfoCreator(general.TrashInfoPath, fileNameInTrash, absPath)
				}
			} else {
				flag := false
				for _, partition := range partitionInfo {
					// 判断文件是否在可移动设备上
					if partition.Removable && strings.HasPrefix(absPath, partition.Mount) {
						// 构建设备回收站参数
						trashFilePath := filepath.Join(partition.Mount, general.DeviceTrashPath, "files") // 回收站文件存储路径
						trashinfoPath := filepath.Join(partition.Mount, general.DeviceTrashPath, "info")  // 已删除文件的 trashinfo 文件路径
						filePathInTrash := filepath.Join(trashFilePath, fileNameInTrash)                  // 文件在回收站的路径
						// 创建回收站
						if err := general.CreateDir(trashFilePath); err != nil {
							fileName, lineNo := general.GetCallerInfo()
							color.Printf("%s %s -> Unable to create trash folder: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
						}
						// 若回收站中存在同名文件，则为待删除文件名字后增加一个累加的数字后缀
						for num := 1; ; num++ {
							if !general.FileExist(filePathInTrash) {
								break
							}
							fileNameInTrash = color.Sprintf("%s_%d", filename, num)
							filePathInTrash = filepath.Join(trashFilePath, fileNameInTrash)
						}
						// 将文件移动到回收站
						if err = os.Rename(file, filePathInTrash); err != nil {
							fileName, lineNo := general.GetCallerInfo()
							color.Printf("%s %s -> Unable to move to trash: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
						} else {
							trashinfoCreator(trashinfoPath, fileNameInTrash, filename) // 可移动设备"已删除文件的原路径"不能用绝对路径而应用纯粹的文件名
						}
						flag = true
						break
					}
				}
				if !flag {
					fileName, lineNo := general.GetCallerInfo()
					color.Printf("%s %s -> Cross-file system operations: move '%s' to '%s'\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), file, general.TrashPath)
				}
			}
		} else {
			fileName, lineNo := general.GetCallerInfo()
			color.Printf("%s %s -> Unable to remove '%s': No such file or directory\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), file)
			break
		}
	}
}

// trashinfoCreator 创建已删除文件的 trashinfo 文件
//
// 参数：
//   - trashPath: trashinfo 文件的存储路径
//   - fileName: trashinfo 文件的文件名（不包含后缀名）
//   - originalPath: 已删除文件的原路径
func trashinfoCreator(trashPath, fileName, originalPath string) {
	var writeMode = "t" // 写入模式

	// 创建已删除文件的 trashinfo 文件
	trashinfoFilePath := filepath.Join(trashPath, color.Sprintf("%s.trashinfo", filepath.Base(fileName)))
	if err := general.CreateFile(trashinfoFilePath); err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s -> Unable to create trashinfo file: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
		return
	}

	// 写入已删除文件信息
	trashinfoFileContent := color.Sprintf(general.TrashInfoFileContent, originalPath, general.GetDateTime(general.TrashInfoFileTimeFormat))
	if err := general.WriteFile(trashinfoFilePath, trashinfoFileContent, writeMode); err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s -> Unable to write trashinfo file: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}
}
