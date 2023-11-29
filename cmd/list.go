/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 程序子命令'list'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files in the file recycle bin",
	Long:  `List all files in the file recycle bin.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	listCmd.Flags().BoolP("help", "h", false, "help for list command")
	rootCmd.AddCommand(listCmd)
}
