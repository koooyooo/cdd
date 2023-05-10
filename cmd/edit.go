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
		var cName, cArgs string
		switch runtime.GOOS {
		case "darwin", "linux", "freebsd":
			cName = "vim"
			cArgs = filepath.Join(hd, ".cdd.yaml")
		case "windows":
			cName = "cmd"
			cArgs = "/c start notpad.exe .cdd.yaml"
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
