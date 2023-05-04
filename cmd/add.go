/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/koooyooo/cdd/model"
	"github.com/koooyooo/cdd/repo"
	"github.com/spf13/cobra"
	"log"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "add new alias: $ cdd add {name} {dir-path}",
	Long:    ``,
	Example: "cdd add github ${HOME}/github",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println(cmd.UsageString())
			return
		}
		name := args[0]
		dir := args[1]
		if err := repo.Instance().Add(&model.Alias{
			Name: name,
			Dir:  dir,
		}); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
