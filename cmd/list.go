/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/koooyooo/cdd/repo"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "list pre-added aliases",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := repo.Instance()
		as, err := r.List()
		if err != nil {
			log.Fatal(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		for i, a := range as {
			table.Append([]string{fmt.Sprintf("%3d", i), a.Name, a.Dir})
		}
		table.Render()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
