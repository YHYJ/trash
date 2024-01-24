/*
File: put.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 09:44:01

Description: 子命令`put`的实现
*/

package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yhyj/trash/general"
)

// PutFiles 将文件移动到回收站
//
// 参数：
//   - files: 需要移动的文件列表
func PutFiles(files []string) {
	for _, file := range files {
		if general.FileExist(file) {
			// 待删除文件
			absPath := general.GetAbsPath(file)   // 待删除文件的绝对路径
			filename := filepath.Base(file)       // 待删除文件名（清除路径等因素，仅保留纯粹的文件名）
			deviceRoot := file                    // 初始化文件系统根路径为文件路径
			fsID, err := general.GetFsID(absPath) // 获取文件系统设备 ID
			if err != nil {
				fmt.Printf(general.RegelarFormat, err)
			}
			// 回收站文件
			trashedFileName := filename                                               // 回收站文件名
			trashedFilePath := filepath.Join(general.TrashFilesPath, trashedFileName) // 回收站文件的路径

			// 若回收站中存在待删除文件的同名文件则为待删除文件增加一个累加的数字后缀作为其在回收站中的文件名
			for num := 1; ; num++ {
				if !general.FileExist(trashedFilePath) {
					break
				}
				trashedFileName = fmt.Sprintf("%s_%d", filename, num)
				trashedFilePath = filepath.Join(general.TrashFilesPath, trashedFileName)
			}
			// 将文件移动到回收站
			err = os.Rename(file, trashedFilePath)
			if err != nil {
				// 跨文件系统移动文件时，重设回收站参数
				if linkErr, ok := err.(*os.LinkError); ok && linkErr.Op == "rename" {
					// 逐级向上查找文件系统的根路径
					for {
						// 获取父目录路径，如果父目录路径和当前路径相同，表示已到达根目录
						parent := filepath.Dir(deviceRoot)
						if parent == deviceRoot {
							break
						}

						parentFsID, err := general.GetFsID(parent) // 获取父文件系统设备 ID
						if err != nil {
							fmt.Printf(general.RegelarFormat, err)
							break
						}

						// 如果父目录所在文件系统的设备 ID 不同于原始文件的设备 ID，表示已经跨越文件系统边界
						if parentFsID != fsID {
							break
						}

						// 更新根路径为父目录路径，继续向上查找
						deviceRoot = parent
					}
					trashFlePath := filepath.Join(deviceRoot, general.CrossTrashPath, "files")
					trashinfoFlePath := filepath.Join(deviceRoot, general.CrossTrashPath, "info")
					err := general.CreateDir(trashFlePath)
					if err != nil {
						fmt.Printf(general.ErrorSuffixFormat, "Error creating trash folder", ": ", err)
					}
					trashedFilePath = filepath.Join(trashFlePath, trashedFileName)
					// 检测回收站中 trashPath 是否已存在，存在则为 trashPath 增加一个累加的后缀
					for num := 1; ; num++ {
						if !general.FileExist(trashedFilePath) {
							break
						}
						trashedFileName = fmt.Sprintf("%s_%d", filename, num)
						trashedFilePath = filepath.Join(trashFlePath, trashedFileName)
					}
					// 将文件移动到回收站
					err = os.Rename(file, trashedFilePath)
					if err != nil {
						fmt.Printf(general.ErrorSuffixFormat, "Error moving to trash", ": ", err)
					} else {
						trashinfoCreator(trashinfoFlePath, trashedFileName, filename)
					}
				} else {
					fmt.Printf(general.ErrorSuffixFormat, fmt.Sprintf("Cannot remove '%s'", file), ": ", err)
				}
			} else {
				trashinfoCreator(general.TrashInfoPath, trashedFileName, absPath)
			}
		} else {
			fmt.Printf(general.ErrorSuffixFormat, fmt.Sprintf("Cannot remove '%s'", file), ": ", "No such file or directory")
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
	// 创建已删除文件的 trashinfo 文件
	trashinfoFilePath := filepath.Join(trashPath, fmt.Sprintf("%s.trashinfo", filepath.Base(fileName)))
	if err := general.CreateFile(trashinfoFilePath); err != nil {
		fmt.Printf(general.ErrorSuffixFormat, "Error creating trashinfo file", ": ", err)
		return
	}

	// 写入已删除文件信息
	trashinfoFileContent := fmt.Sprintf(general.TrashInfoFileContent, originalPath, general.GetDateTime(general.TrashInfoFileTimeFormat))
	if err := general.WriteFile(trashinfoFilePath, trashinfoFileContent); err != nil {
		fmt.Printf(general.ErrorSuffixFormat, "Error writing trashinfo file", ": ", err)
	}
}
