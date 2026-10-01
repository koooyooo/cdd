package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveAddArgs(t *testing.T) {
	t.Run("two args explicit name", func(t *testing.T) {
		name, dir, err := resolveAddArgs([]string{"github", "/tmp/github"}, "")
		require.NoError(t, err)
		assert.Equal(t, "github", name)
		assert.Equal(t, "/tmp/github", dir)
	})

	t.Run("one arg derives basename", func(t *testing.T) {
		name, dir, err := resolveAddArgs([]string{"/tmp/myproj"}, "")
		require.NoError(t, err)
		assert.Equal(t, "myproj", name)
		assert.Equal(t, "/tmp/myproj", dir)
	})

	t.Run("one arg with --alias", func(t *testing.T) {
		name, dir, err := resolveAddArgs([]string{"/tmp/myproj"}, "mp")
		require.NoError(t, err)
		assert.Equal(t, "mp", name)
		assert.Equal(t, "/tmp/myproj", dir)
	})

	t.Run("two args with --alias is error", func(t *testing.T) {
		_, _, err := resolveAddArgs([]string{"github", "/tmp/github"}, "mp")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--alias")
	})

	t.Run("zero args is error", func(t *testing.T) {
		_, _, err := resolveAddArgs(nil, "")
		require.Error(t, err)
	})

	t.Run("three args is error", func(t *testing.T) {
		_, _, err := resolveAddArgs([]string{"a", "b", "c"}, "")
		require.Error(t, err)
	})
}

func TestAliasFromPath(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		got, err := aliasFromPath("/var/log/nginx")
		require.NoError(t, err)
		assert.Equal(t, "nginx", got)
	})

	t.Run("trailing slash", func(t *testing.T) {
		got, err := aliasFromPath("/var/log/nginx/")
		require.NoError(t, err)
		assert.Equal(t, "nginx", got)
	})

	t.Run("dot uses cwd basename", func(t *testing.T) {
		orig, err := os.Getwd()
		require.NoError(t, err)
		dir := t.TempDir()
		require.NoError(t, os.Chdir(dir))
		t.Cleanup(func() {
			_ = os.Chdir(orig)
		})
		got, err := aliasFromPath(".")
		require.NoError(t, err)
		assert.Equal(t, filepath.Base(dir), got)
	})

	t.Run("home var expands", func(t *testing.T) {
		home, err := os.UserHomeDir()
		require.NoError(t, err)
		got, err := aliasFromPath("${HOME}")
		require.NoError(t, err)
		assert.Equal(t, filepath.Base(home), got)
	})
}
