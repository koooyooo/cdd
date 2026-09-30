package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/koooyooo/cdd/repo"
	"github.com/spf13/cobra"
)

var completeAliasesCmd = &cobra.Command{
	Use:    "complete-aliases",
	Short:  "print alias names for shell completion",
	Hidden: true,
	Args:   cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := printAliasNames(os.Stdout, repo.Instance()); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeAliasesCmd)
}

func aliasNames(r repo.Repo) ([]string, error) {
	aliases, err := r.List()
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(aliases))
	for _, a := range aliases {
		if a.Name == "" {
			continue
		}
		names = append(names, a.Name)
	}
	return names, nil
}

func printAliasNames(w io.Writer, r repo.Repo) error {
	names, err := aliasNames(r)
	if err != nil {
		return err
	}
	for _, name := range names {
		if _, err := fmt.Fprintln(w, name); err != nil {
			return err
		}
	}
	return nil
}
