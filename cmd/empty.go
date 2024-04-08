/*
File: empty.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 执行子命令 'empty'
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/trash/cli"
	"github.com/yhyj/trash/general"
)

// emptyCmd represents the empty command
var emptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Empty files in the recycle bin",
	Long:  `Empty files in the recycle bin, number each file, 0 represents all files.`,
	Run: func(cmd *cobra.Command, args []string) {
		general.CheckRecycleBin() // 检查回收站是否存在
		if general.Confirm("Are you sure you want to empty the recycle bin? (yes/No): ", "yes") {
			cli.EmptyTrash() // 清空回收站
		}
	},
}

func init() {
	emptyCmd.Flags().BoolP("help", "h", false, "help for empty command")
	rootCmd.AddCommand(emptyCmd)
}
