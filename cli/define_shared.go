/*
File: define_shared.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-30 14:29:16

Description: trash 整体会用到的函数
*/

package cli

import (
	"fmt"

	"github.com/yhyj/trash/general"
)

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
