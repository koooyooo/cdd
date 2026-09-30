package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init {bash|zsh}",
	Short: "print shell integration code for bash or zsh",
	Long: `Print shell integration that wraps cdd so alias jumps use builtin cd.
The script also completes alias names and subcommands.

Add this to your shell rc (example for zsh, after compinit):

  eval "$(cdd init zsh)"

When stdout is a TTY, you will be asked whether to copy that eval line to the clipboard.`,
	Example: `$ eval "$(cdd init zsh)"
$ cdd init zsh   # interactive: optionally copy eval line to clipboard`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]
		script, err := shellIntegration(shell)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s (use bash or zsh)\n", err.Error())
			os.Exit(1)
		}
		fmt.Print(script)

		if !stdoutIsTTY() {
			return
		}
		maybeCopyInitLine(shell)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func stdoutIsTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

func maybeCopyInitLine(shell string) {
	fmt.Fprintf(os.Stderr, "Copy eval \"$(cdd init %s)\" to clipboard for pasting into your shell rc? [Y/n] ", shell)

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\ncannot open /dev/tty: %v\n", err)
		return
	}
	defer tty.Close()

	answer, err := bufio.NewReader(tty).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "failed to read answer: %v\n", err)
		return
	}
	answer = strings.TrimSpace(strings.ToLower(answer))
	if answer != "" && answer != "y" && answer != "yes" {
		return
	}

	line := fmt.Sprintf("eval \"$(cdd init %s)\"\n", shell)
	if err := copyToClipboard(line); err != nil {
		fmt.Fprintf(os.Stderr, "clipboard unavailable (%v). Add this line to your rc manually:\n  %s", err, line)
		return
	}
	fmt.Fprintf(os.Stderr, "copied to clipboard. Paste into ~/.%src (or equivalent).\n", shell)
}

func copyToClipboard(text string) error {
	candidates := [][]string{
		{"pbcopy"},
		{"wl-copy"},
		{"xclip", "-selection", "clipboard"},
	}
	var lastErr error
	for _, args := range candidates {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = strings.NewReader(text)
		if err := cmd.Run(); err != nil {
			lastErr = err
			continue
		}
		return nil
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("no clipboard command found")
	}
	return lastErr
}
