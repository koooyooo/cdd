/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/koooyooo/cdd/repo"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cdd",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"hello", "world"}, cobra.ShellCompDirectiveFilterFileExt
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			listCmd.Run(cmd, args)
			return
		}
		tgt := args[0]
		path, find, err := findPath(tgt)
		if err != nil {
			log.Fatal(err)
		}
		if !find {
			fmt.Printf("no path found: %s\n", tgt)
			return
		}
		cd(path)
	},
}

func findPath(tgt string) (string, bool, error) {
	r := repo.Instance()
	a, foundByName, err := r.Get(tgt)
	if err != nil {
		return "", false, err
	}
	// name-based selection
	if foundByName {
		path, err := a.ReplacedDir()
		if err != nil {
			return "", false, err
		}
		return path, true, nil
	}
	// num-based selection
	num, err := strconv.Atoi(tgt)
	if err != nil {
		return "", false, err
	}
	list, err := r.List()
	if len(list) <= num {
		return "", false, nil
	}
	path, err := list[num].ReplacedDir()
	if err != nil {
		return "", false, nil
	}
	return path, true, nil

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func cd(path string) {
	cmd := exec.Command(detectShell())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func detectShell() string {
	shell := os.Getenv("SHELL")
	if shell != "" {
		return shell
	}
	if runtime.GOOS == "windows" {
		return os.Getenv("COMSPEC")
	}
	return "/bin/sh"
}
