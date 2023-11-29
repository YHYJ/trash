/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 程序子命令'restore'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore files from file recycle bin",
	Long:  `Restore files from file recycle bin, number each file, 0 represents all files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restore called")
	},
}

func init() {
	restoreCmd.Flags().BoolP("help", "h", false, "help for list command")
	rootCmd.AddCommand(restoreCmd)
}
