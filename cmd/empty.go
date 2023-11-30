/*
File: empty.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 程序子命令'empty'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// emptyCmd represents the empty command
var emptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Empty files in the recycle bin",
	Long:  `Empty files in the recycle bin, number each file, 0 represents all files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("empty called")
	},
}

func init() {
	emptyCmd.Flags().BoolP("help", "h", false, "help for empty command")
	rootCmd.AddCommand(emptyCmd)
}
