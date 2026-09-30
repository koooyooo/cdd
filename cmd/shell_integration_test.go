package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShellIntegrationContents(t *testing.T) {
	bashScript, err := shellIntegration("bash")
	require.NoError(t, err)
	assert.Contains(t, bashScript, "builtin cd --")
	assert.Contains(t, bashScript, "|complete-aliases)")
	assert.Contains(t, bashScript, "command cdd complete-aliases")
	assert.Contains(t, bashScript, "complete -F _cdd_complete cdd")
	assert.NotContains(t, bashScript, "compdef")

	zshScript, err := shellIntegration("zsh")
	require.NoError(t, err)
	assert.Contains(t, zshScript, "builtin cd --")
	assert.Contains(t, zshScript, "compdef _cdd cdd")
	assert.Contains(t, zshScript, "whence compdef")
	assert.NotContains(t, zshScript, "complete -F")

	_, err = shellIntegration("fish")
	require.EqualError(t, err, "unsupported shell: fish")
}

func TestShellIntegrationSyntax(t *testing.T) {
	for _, shell := range []string{"bash", "zsh"} {
		shell := shell
		t.Run(shell, func(t *testing.T) {
			if _, err := exec.LookPath(shell); err != nil {
				t.Skip(shell + " not installed")
			}
			script, err := shellIntegration(shell)
			require.NoError(t, err)
			cmd := exec.Command(shell, "-n")
			cmd.Stdin = strings.NewReader(script)
			out, err := cmd.CombinedOutput()
			require.NoError(t, err, string(out))
		})
	}
}

func TestBashCompletionCandidates(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not installed")
	}

	binDir := t.TempDir()
	workDir := t.TempDir()
	fake := filepath.Join(binDir, "cdd")
	fakeSrc := "#!/bin/sh\nif [ \"$1\" = complete-aliases ]; then\n  printf '%s\\n' home docs 'my docs'\n  exit 0\nfi\necho \"unexpected: $*\" >&2\nexit 1\n"
	require.NoError(t, os.WriteFile(fake, []byte(fakeSrc), 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(workDir, "note.txt"), []byte("x"), 0o644))

	script, err := shellIntegration("bash")
	require.NoError(t, err)

	candidates := func(t *testing.T, words string) []string {
		t.Helper()
		body := script + "\n" + words + "\n" +
			"_cdd_complete\n" +
			"printf '%s\\n' \"${COMPREPLY[@]}\"\n"
		cmd := exec.Command("bash", "--noprofile", "--norc", "-c", body)
		cmd.Dir = workDir
		cmd.Env = []string{"PATH=" + binDir}
		out, err := cmd.CombinedOutput()
		require.NoError(t, err, string(out))
		text := strings.TrimSuffix(string(out), "\n")
		if text == "" {
			return nil
		}
		return strings.Split(text, "\n")
	}

	t.Run("first word prefix", func(t *testing.T) {
		got := candidates(t, "COMP_WORDS=(cdd d)\nCOMP_CWORD=1")
		assert.Equal(t, []string{"docs", "delete", "del", "down"}, got)
	})

	t.Run("move-up prefix", func(t *testing.T) {
		got := candidates(t, "COMP_WORDS=(cdd move)\nCOMP_CWORD=1")
		assert.Equal(t, []string{"move-up", "move-down"}, got)
	})

	t.Run("move-up argument", func(t *testing.T) {
		got := candidates(t, "COMP_WORDS=(cdd move-up '')\nCOMP_CWORD=2")
		assert.Equal(t, []string{"home", "docs", "my docs"}, got)
	})

	t.Run("alias with space", func(t *testing.T) {
		got := candidates(t, "COMP_WORDS=(cdd my)\nCOMP_CWORD=1")
		assert.Equal(t, []string{"my docs"}, got)
	})

	t.Run("remove argument", func(t *testing.T) {
		got := candidates(t, "COMP_WORDS=(cdd remove '')\nCOMP_CWORD=2")
		assert.Equal(t, []string{"home", "docs", "my docs"}, got)
	})

	t.Run("init argument", func(t *testing.T) {
		got := candidates(t, "COMP_WORDS=(cdd init b)\nCOMP_CWORD=2")
		assert.Equal(t, []string{"bash"}, got)
	})

	t.Run("add path", func(t *testing.T) {
		got := candidates(t, "COMP_WORDS=(cdd add myname '')\nCOMP_CWORD=3")
		assert.Contains(t, got, "note.txt")
	})

	t.Run("add name is not a path", func(t *testing.T) {
		got := candidates(t, "COMP_WORDS=(cdd add '')\nCOMP_CWORD=2")
		assert.Empty(t, got)
	})
}
