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
	// 获取所有 trashinfo 文件的绝对路径
	trashinfoFiles, err := filepath.Glob(filepath.Join(general.TrashinfoFilePath, "*"))
	if err != nil {
		fmt.Printf(general.ErrorSuffixFormat, "Error listing trashinfo file", ": ", err)
		return
	}

	for _, trashinfoFile := range trashinfoFiles {
		if general.FileExist(trashinfoFile) {
			originalFilePath := strings.Split(general.ReadFileKey(trashinfoFile, "Path"), "=")[1]
			trashedDateTime := strings.Split(general.ReadFileKey(trashinfoFile, "DeletionDate"), "=")[1]
			trashedDate, trashedTime, err := general.ParseDateTime(general.TrashinfoTimeFormat, trashedDateTime)
			if err != nil {
				fmt.Printf(general.ErrorSuffixFormat, "Error parsing trashinfo file", ": ", err)
				return
			}
			fmt.Printf("%s %s %s\n", trashedDate, trashedTime, originalFilePath)
		}
	}
}
