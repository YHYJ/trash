/*
File: restore.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 程序子命令'restore'时执行
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/trash/cli"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore files from file recycle bin",
	Long:  `Restore files from file recycle bin, number each file, 0 represents all files.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.CheckRecycleBin()  // 检查回收站是否存在
		cli.RestoreFromTrash() // 恢复回收站中的文件
	},
}

func init() {
	restoreCmd.Flags().BoolP("help", "h", false, "help for restore command")
	rootCmd.AddCommand(restoreCmd)
}
