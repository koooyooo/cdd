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
	"regexp"
	"strconv"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove { name | num }",
	Short:   "remove specified alias",
	Long:    ``,
	Aliases: []string{"rm", "delete", "del"},
	Example: `$ cdd remove home`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(cmd.UsageString())
			return
		}
		name := args[0]
		r := repo.Instance()

		// 番号指定対応
		aliases, err := r.List()
		if err != nil {
			log.Fatal(err)
		}
		idx := index(aliases, name)
		if idx == -1 {
			findName, ok, err := findByNum(aliases, name)
			if err != nil {
				log.Fatal(err)
			}
			if !ok {
				fmt.Printf("no alias found for %s\n", name)
				return
			}
			name = findName
		}
		if err := r.Remove(name); err != nil {
			log.Fatal(err)
		}
	},
}

func index(aliases []*model.Alias, name string) int {
	for i, a := range aliases {
		if a.Name == name {
			return i
		}
	}
	return -1
}

func findByNum(aliases []*model.Alias, name string) (string, bool, error) {
	isDig, err := isDigit(name)
	if err != nil {
		log.Fatal(err)
	}
	if !isDig {
		return "", false, nil
	}
	idx, err := strconv.Atoi(name)
	if err != nil {
		return "", false, nil
	}
	if len(aliases) <= idx {
		return "", false, fmt.Errorf("index out of range max=%d, idx=%d", len(aliases)-1, idx)
	}
	return aliases[idx].Name, true, nil

}

func isDigit(name string) (bool, error) {
	return regexp.MatchString("[0-9]*", name)
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
