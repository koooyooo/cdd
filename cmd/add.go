/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/koooyooo/cdd/common"
	"github.com/koooyooo/cdd/model"
	"github.com/koooyooo/cdd/repo"
	"github.com/spf13/cobra"
)

var addAliasFlag string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add [name] path",
	Aliases: []string{"a"},
	Short:   "add new alias: $ cdd add {path} | $ cdd add {name} {path}",
	Long:    ``,
	Example: `$ cdd add ${HOME}/github              # alias = github
$ cdd add . -a work                    # alias from --alias
$ cdd add github ${HOME}/github        # explicit name + path
$ cdd add github '${HOME}/github'      # deny shell expansion of ${HOME}`,
	Run: func(cmd *cobra.Command, args []string) {
		name, dir, err := resolveAddArgs(args, addAliasFlag)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(cmd.UsageString())
			return
		}
		dir, err = common.Replace4Store(dir)
		if err != nil {
			log.Fatal(err)
		}
		if err := repo.Instance().Add(&model.Alias{
			Name: name,
			Dir:  dir,
		}); err != nil {
			log.Fatal(err)
		}
	},
}

func resolveAddArgs(args []string, aliasFlag string) (name, dir string, err error) {
	switch len(args) {
	case 1:
		dir = args[0]
		if aliasFlag != "" {
			return aliasFlag, dir, nil
		}
		name, err = aliasFromPath(dir)
		if err != nil {
			return "", "", err
		}
		return name, dir, nil
	case 2:
		if aliasFlag != "" {
			return "", "", fmt.Errorf("--alias cannot be used together with a positional alias name")
		}
		return args[0], args[1], nil
	default:
		return "", "", fmt.Errorf("accepts 1 or 2 arguments")
	}
}

func aliasFromPath(path string) (string, error) {
	expanded, err := common.Replace4Get(path)
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(expanded)
	if err != nil {
		return "", err
	}
	base := filepath.Base(abs)
	if base == "" || base == "." || base == ".." || base == string(filepath.Separator) {
		return "", fmt.Errorf("cannot derive alias from path: %s", path)
	}
	return base, nil
}

func init() {
	addCmd.Flags().StringVarP(&addAliasFlag, "alias", "a", "", "alias name (default: last directory of path)")
	rootCmd.AddCommand(addCmd)
}
