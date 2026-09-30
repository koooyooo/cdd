/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/koooyooo/cdd/repo"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cdd",
	Short: "Jump to bookmarked directories (use shell integration for cd)",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
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

func handleOpenOpt(cmd *cobra.Command, path string) error {
	open, err := cmd.Flags().GetBool("open")
	if err != nil {
		return err
	}
	if !open {
		return nil
	}
	cmdStr, args, err := openCommandStr(path)
	if err != nil {
		return err
	}
	return exec.Command(cmdStr, args...).Run()
}

func findPath(tgt string) (string, bool, error) {
	return findPathWithRepo(repo.Instance(), tgt)
}

func findPathWithRepo(r repo.Repo, tgt string) (string, bool, error) {
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
		return "", false, nil
	}
	list, err := r.List()
	if err != nil {
		return "", false, err
	}
	if len(list) <= num || num < 0 {
		return "", false, nil
	}
	path, err := list[num].ReplacedDir()
	if err != nil {
		return "", false, err
	}
	return path, true, nil
}

func openCommandStr(path string) (string, []string, error) {
	switch runtime.GOOS {
	case "darwin":
		return "open", []string{path}, nil
	case "windows":
		return "start", []string{path}, nil
	case "linux":
		return "xdg-open", []string{path}, nil
	}
	return "", nil, fmt.Errorf("unsupported os: %s", runtime.GOOS)
}
