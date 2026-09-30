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

// upCmd represents the move-up command (alias: up)
var upCmd = &cobra.Command{
	Use:     "move-up name [amount]",
	Aliases: []string{"up"},
	Short:   "make specified alias to be upper on the list",
	Long:    ``,
	Example: "$ cdd move-up github\n$ cdd move-up github 2\n$ cdd up github",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || len(args) > 2 {
			fmt.Println(cmd.UsageString())
			return
		}
		name := args[0]
		var num = 1
		if len(args) == 2 {
			numStr := args[1]
			n, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("failed to parse 2nd arg as number: %s\n", numStr)
				return
			}
			num = n
		}
		if err := repo.Instance().Move(name, num); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
