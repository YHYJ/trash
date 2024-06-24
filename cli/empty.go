/*
File: empty.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 15:49:08

Description: 子命令 'empty' 的实现
*/

package cli

import (
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/yhyj/trash/general"
)

// EmptyTrash 清空回收站
func EmptyTrash() {
	trashFilesName, err := os.ReadDir(general.TrashFilePath)
	if err != nil {
		fileName, lineNo := general.GetCallerInfo()
		color.Printf("%s %s -> Unable to read trash folder: %s\n", general.DangerText("Error:"), general.SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}

	for _, trashFileName := range trashFilesName {
		trashFile := filepath.Join(general.TrashFilePath, trashFileName.Name())
		trashinfoFile := filepath.Join(general.TrashInfoPath, color.Sprintf("%s.trashinfo", trashFileName.Name()))
		os.RemoveAll(trashFile)     // 删除回收站中的文件
		os.RemoveAll(trashinfoFile) // 删除 trashinfo 文件
	}

	// 删除 directorysizes 文件
	directorysizesFile := filepath.Join(general.TrashPath, "directorysizes")
	if general.FileExist(directorysizesFile) {
		os.RemoveAll(directorysizesFile)
	}
}
