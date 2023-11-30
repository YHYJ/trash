/*
File: list.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 程序子命令'list'时执行
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/trash/cli"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files in the file recycle bin",
	Long:  `List all files in the file recycle bin.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.CheckRecycleBin() // 检查回收站是否存在
		cli.ListFiles()       // 列出回收站中的文件
	},
}

func init() {
	listCmd.Flags().BoolP("help", "h", false, "help for list command")
	rootCmd.AddCommand(listCmd)
}
