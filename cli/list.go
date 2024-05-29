/*
File: list.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-30 14:40:28

Description: 子命令 'list' 的实现
*/

package cli

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/gookit/color"
	"github.com/yhyj/trash/general"
)

func ListFiles() {
	// 获取所有 trashinfo 文件的绝对路径
	trashinfoFiles, err := filepath.Glob(filepath.Join(general.TrashInfoPath, "*"))
	if err != nil {
		color.Danger.Printf("Error listing trashinfo file: %s\n", err)
		return
	}

	// 将数据解析为 FileEntry 类型
	var fileEntries []general.FileEntry
	for index, trashinfoFile := range trashinfoFiles {
		originalFilePath := strings.Split(general.ReadFileKey(trashinfoFile, "Path"), "=")[1]           // 文件的原绝对路径
		deletionDate := strings.Split(general.ReadFileKey(trashinfoFile, "DeletionDate"), "=")[1]       // 文件的删除日期时间（未解析）
		parsedDeletionDate, err := general.ParseDateTime(general.TrashInfoFileTimeFormat, deletionDate) // 文件的删除日期时间（已解析）
		if err != nil {
			color.Danger.Printf("Error parsing trashinfo file: %s\n", err)
			break
		}

		entry := general.FileEntry{
			Index:        index,
			Time:         parsedDeletionDate,
			OriginalPath: originalFilePath,
		}
		fileEntries = append(fileEntries, entry)

		// 按时间排序
		sort.SliceStable(fileEntries, func(i, j int) bool {
			return fileEntries[i].Time.Before(fileEntries[j].Time)
		})
		// 更新排序后的 Index
		for i := range fileEntries {
			fileEntries[i].Index = i
		}
	}

	// 输出文件列表
	for _, entry := range fileEntries {
		color.Printf("%s %s %s\n", general.FgCyanText(entry.Time.Format("2006-01-02")), general.FgMagentaText(entry.Time.Format("15:04:05")), entry.OriginalPath)
	}

	// 输出文件总数
	color.Printf("%s %s\n", general.SecondaryText("Total:"), general.SecondaryText(len(fileEntries)))
}
