package cmd

import (
	"testing"

	"github.com/koooyooo/cdd/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeRepo struct {
	aliases []*model.Alias
}

func (f *fakeRepo) Init() error { return nil }

func (f *fakeRepo) List() ([]*model.Alias, error) {
	return f.aliases, nil
}

func (f *fakeRepo) Get(name string) (*model.Alias, bool, error) {
	for _, a := range f.aliases {
		if a.Name == name {
			return a, true, nil
		}
	}
	return nil, false, nil
}

func (f *fakeRepo) Add(alias *model.Alias) error    { return nil }
func (f *fakeRepo) Remove(name string) error        { return nil }
func (f *fakeRepo) Move(name string, num int) error { return nil }

func TestFindPathWithRepo(t *testing.T) {
	home := t.TempDir()
	docs := t.TempDir()
	r := &fakeRepo{aliases: []*model.Alias{
		{Name: "home", Dir: home},
		{Name: "docs", Dir: docs},
	}}

	t.Run("by name", func(t *testing.T) {
		path, found, err := findPathWithRepo(r, "docs")
		require.NoError(t, err)
		assert.True(t, found)
		assert.Equal(t, docs, path)
	})

	t.Run("by index", func(t *testing.T) {
		path, found, err := findPathWithRepo(r, "0")
		require.NoError(t, err)
		assert.True(t, found)
		assert.Equal(t, home, path)
	})

	t.Run("unknown name", func(t *testing.T) {
		_, found, err := findPathWithRepo(r, "missing")
		require.NoError(t, err)
		assert.False(t, found)
	})

	t.Run("index out of range", func(t *testing.T) {
		_, found, err := findPathWithRepo(r, "9")
		require.NoError(t, err)
		assert.False(t, found)
	})
}
