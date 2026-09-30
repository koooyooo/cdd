package cmd

import (
	"bytes"
	"testing"

	"github.com/koooyooo/cdd/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAliasNames(t *testing.T) {
	r := &fakeRepo{aliases: []*model.Alias{
		{Name: "home", Dir: "/h"},
		{Name: "", Dir: "/x"},
		{Name: "docs", Dir: "/d"},
	}}

	names, err := aliasNames(r)
	require.NoError(t, err)
	assert.Equal(t, []string{"home", "docs"}, names)
}

func TestPrintAliasNames(t *testing.T) {
	r := &fakeRepo{aliases: []*model.Alias{
		{Name: "home", Dir: "/h"},
		{Name: "docs", Dir: "/d"},
	}}
	var buf bytes.Buffer
	require.NoError(t, printAliasNames(&buf, r))
	assert.Equal(t, "home\ndocs\n", buf.String())
}
