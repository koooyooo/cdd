/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/koooyooo/cdd/common"
	"github.com/koooyooo/cdd/model"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var defaultAliases = []model.Alias{
	{
		Name: "home",
		Dir:  "${HOME}",
	},
	{
		Name: "desktop",
		Dir:  "${HOME}/Desktop",
	},
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize basic settings",
	Long:  `- creating ${HOME}/.cdd.yaml with basic settings`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := common.CDDPath()
		if err != nil {
			log.Fatal(err)
		}
		if common.Exists(path) {
			return
		}
		b, err := yaml.Marshal(defaultAliases)
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile(path, b, 0655); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
