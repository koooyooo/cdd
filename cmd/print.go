package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:     "print {name|num}",
	Aliases: []string{"p"},
	Short:   "print resolved path for an alias or list index",
	Long:    ``,
	Example: `$ cdd print docs
$ cd "$(cdd print docs)"
$ cdd print -o docs`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tgt := args[0]
		path, found, err := findPath(tgt)
		if err != nil {
			log.Fatal(err)
		}
		if !found {
			fmt.Fprintf(os.Stderr, "no path found: %s\n", tgt)
			os.Exit(1)
		}
		if err := handleOpenOpt(cmd, path); err != nil {
			log.Fatalf("fail in handling open option: %v\n", err)
		}
		fmt.Println(path)
	},
}

func init() {
	printCmd.Flags().BoolP("open", "o", false, "Open window for the path before printing")
	rootCmd.AddCommand(printCmd)
}
