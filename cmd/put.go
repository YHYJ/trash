/*
File: put.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 执行子命令 'put'
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/trash/cli"
	"github.com/yhyj/trash/general"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Put files into the recycle bin",
	Long:  `Put files into the recycle bin instead of deleting them completely.`,
	Run: func(cmd *cobra.Command, args []string) {
		general.CheckRecycleBin() // 检查回收站是否存在
		cli.PutFiles(args)        // 将文件放入回收站
	},
}

func init() {
	putCmd.Flags().BoolP("help", "h", false, "help for put command")
	rootCmd.AddCommand(putCmd)
}
