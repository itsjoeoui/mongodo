/*
Copyright © 2022 Joey Yu <joey@itsjoeoui.com>
*/
package cmd

import (
	"log"
	"mongodo/db"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a given task",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("delete called")
		db.DeleteTask(strings.Join(args, ""))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
