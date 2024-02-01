/*
File: restore.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 16:55:14

Description: 子命令`restore`的实现
*/

package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/yhyj/trash/general"
)

// RestoreFromTrash 恢复回收站中的文件
func RestoreFromTrash() {
	// 获取所有 trashinfo 文件的绝对路径
	trashinfoFiles, err := filepath.Glob(filepath.Join(general.TrashInfoPath, "*"))
	if err != nil {
		fmt.Printf(general.ErrorSuffixFormat, "Error listing trashinfo file", ": ", err)
		return
	}

	// 将数据解析为 FileEntry 类型
	var fileEntries []FileEntry
	for index, trashinfoFile := range trashinfoFiles {
		originalFilePath := strings.Split(general.ReadFileKey(trashinfoFile, "Path"), "=")[1]           // 文件的原绝对路径
		deletionDate := strings.Split(general.ReadFileKey(trashinfoFile, "DeletionDate"), "=")[1]       // 文件的删除日期时间（未解析）
		parsedDeletionDate, err := general.ParseDateTime(general.TrashInfoFileTimeFormat, deletionDate) // 文件的删除日期时间（已解析）
		if err != nil {
			fmt.Printf(general.ErrorSuffixFormat, "Error parsing trashinfo file", ": ", err)
			break
		}

		entry := FileEntry{
			Index:        index,
			Time:         parsedDeletionDate,
			OriginalPath: originalFilePath,
			Path:         filepath.Join(general.TrashFilePath, strings.TrimSuffix(filepath.Base(trashinfoFile), filepath.Ext(trashinfoFile))),
		}
		fileEntries = append(fileEntries, entry)

		// 按时间排序
		sort.SliceStable(fileEntries, func(i, j int) bool {
			return fileEntries[i].Time.Before(fileEntries[j].Time)
		})
		// 更新排序后的 Index
		for i := range fileEntries {
			fileEntries[i].Index = i + 1 // 显示编号从 1 开始
		}
	}

	// 检测回收站中的文件数
	fileEntriesLen := len(fileEntries)
	if fileEntriesLen != 0 {
		// 输出排序后的数据
		digits := general.CountDigits(fileEntriesLen) // 文本对齐位数
		for _, entry := range fileEntries {
			fmt.Printf("%*d %s %s %s\n", digits+2, entry.Index, entry.Time.Format("2006-01-02"), entry.Time.Format("15:04:05"), entry.OriginalPath)
		}

		// 交互获取要恢复的文件切片
		userIndexs := general.UserFace(fmt.Sprintf("What file to restore (Example: 0 or 1 ... %d, 0 restore all): ", fileEntriesLen))
		if len(userIndexs) == 0 {
			fmt.Println("Exiting")
			return
		}

		var restoreThisFiles []FileEntry // 待恢复文件信息
		// 切片只有一个元素且其值为 0，表示全部恢复
		if len(userIndexs) == 1 && userIndexs[0] == 0 {
			restoreThisFiles = fileEntries
		} else {
			// 对 userIndexs 进行去重得到 uniqueUserIndexs
			uniqueUserIndexs := make(map[int]bool)
			for _, idx := range userIndexs {
				uniqueUserIndexs[idx] = true
			}
			// 将 uniqueUserIndexs 转换为切片得到 sortedUserIndexs
			var sortedUserIndexs []int
			for uniqueUserIndex := range uniqueUserIndexs {
				sortedUserIndexs = append(sortedUserIndexs, uniqueUserIndex)
			}
			// 按升序对 sortedUserIndexs 排序
			slices.Sort(sortedUserIndexs)
			// 遍历 sortedUserIndexs 并在 fileEntries 中使用二分搜索
			for _, sortedUserIndex := range sortedUserIndexs {
				index := sort.Search(fileEntriesLen, func(i int) bool {
					return fileEntries[i].Index >= sortedUserIndex
				})

				if sortedUserIndex == 0 {
					fmt.Printf(general.ErrorBaseFormat, "0 can only be used alone")
					return
				} else if index < fileEntriesLen && sortedUserIndex == fileEntries[index].Index {
					entry := FileEntry{
						Index:        fileEntries[index].Index,
						Time:         fileEntries[index].Time,
						OriginalPath: fileEntries[index].OriginalPath,
						Path:         fileEntries[index].Path,
					}
					restoreThisFiles = append(restoreThisFiles, entry)
				} else {
					fmt.Printf("没有编号为 %d 的文件\n", sortedUserIndex)
					return
				}
			}
		}

		// 开始恢复
		for _, restoreThisFile := range restoreThisFiles {
			if general.FileExist(restoreThisFile.OriginalPath) && !general.FileEmpty(restoreThisFile.OriginalPath) {
				fmt.Printf(general.ErrorSuffixFormat, "Error restoring files", ": ", fmt.Sprintf("%s file already exists and is not empty", restoreThisFile.OriginalPath))
				break
			}
			// 将回收站文件恢复到原路径
			err := os.Rename(restoreThisFile.Path, restoreThisFile.OriginalPath)
			if err != nil {
				fmt.Printf(general.ErrorSuffixFormat, "Error restoring files", ": ", err)
			}
			// 删除其对应的 trashinfo 文件
			general.DeleteFile(filepath.Join(general.TrashInfoPath, fmt.Sprintf("%s.trashinfo", filepath.Base(restoreThisFile.Path))))
		}
	} else {
		fmt.Printf(general.RegelarFormat, "No files in trash")
		return
	}
}
