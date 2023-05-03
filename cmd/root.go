/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/koooyooo/cdd/repo"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cdd",
	Short: "",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r := repo.Instance()
		a, ok, err := r.Get(args[0])
		if err != nil {
			log.Fatal(err)
		}
		if !ok {
			fmt.Printf("No alias found for :%s\n", args[0])
		}
		cd(a.Dir)
	},
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
