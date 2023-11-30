/*
File: list.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-30 14:40:28

Description: 子命令`list`的实现
*/

package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/yhyj/trash/general"
)

func ListFiles() {
	files, err := filepath.Glob(filepath.Join(general.TrashFilePath, "*"))
	if err != nil {
		fmt.Printf(general.ErrorSuffixFormat, "Error listing trash", ": ", err)
		return
	}

	for _, file := range files {
		// 获取基础文件名
		fileName := filepath.Base(file)
		// 获取对应的 trashinfo 文件路径
		trashinfoFilePath := filepath.Join(general.TrashinfoFilePath, fmt.Sprintf("%s.trashinfo", fileName))

		if general.FileExist(trashinfoFilePath) {
			originalFilePath := strings.Split(general.ReadFileKey(trashinfoFilePath, "Path"), "=")[1]
			trashedDateTime := strings.Split(general.ReadFileKey(trashinfoFilePath, "DeletionDate"), "=")[1]
			trashedDate, trashedTime, err := general.ParseDateTime(general.TrashinfoTimeFormat, trashedDateTime)
			if err != nil {
				fmt.Printf(general.ErrorSuffixFormat, "Error parsing trashinfo file", ": ", err)
				return
			}
			fmt.Printf("%s %s %s\n", trashedDate, trashedTime, originalFilePath)
		}
	}
}
