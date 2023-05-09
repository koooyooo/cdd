/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit .cdd.yaml directly",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		hd, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		cName, cArgs := "vim", filepath.Join(hd, ".cdd.yaml")
		switch runtime.GOOS {
		case "darwin", "linux", "freebsd":
			// default
		case "windows":
			// TODO
		}
		c := exec.Command(cName, cArgs) // mac can also use 'open'
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
