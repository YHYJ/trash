/*
File: pacnew.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-11-26 11:05:39

Description: 程序子命令'put'时执行
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Put files into the recycle bin",
	Long:  `Put files into the recycle bin instead of deleting them completely.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("put called")
	},
}

func init() {
	putCmd.Flags().BoolP("help", "h", false, "help for put command")
	rootCmd.AddCommand(putCmd)
}
