/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/koooyooo/cdd/repo"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// downCmd represents the move-down command (alias: down)
var downCmd = &cobra.Command{
	Use:     "move-down name [amount]",
	Aliases: []string{"down"},
	Short:   "make specified alias to be lower on the list",
	Long:    ``,
	Example: "$ cdd move-down github\n$ cdd move-down github 2\n$ cdd down github",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || len(args) > 2 {
			fmt.Println(cmd.UsageString())
			return
		}
		name := args[0]
		var num = -1
		if len(args) == 2 {
			numStr := args[1]
			n, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("failed to parse 2nd arg as number: %s\n", numStr)
				return
			}
			num = 0 - n
		}
		if err := repo.Instance().Move(name, num); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
