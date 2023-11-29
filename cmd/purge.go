/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 程序子命令'purge'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge files from the recycle bin",
	Long:  `Purge files from the recycle bin, number each file, 0 represents all files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("purge called")
	},
}

func init() {
	purgeCmd.Flags().BoolP("help", "h", false, "help for purge command")
	rootCmd.AddCommand(purgeCmd)
}
