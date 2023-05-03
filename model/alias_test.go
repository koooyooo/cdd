package model

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestAliases(t *testing.T) {
	cdds := []Alias{
		Alias{
			Name: "name1",
			Dir:  "${HOME}/work/project1/actions",
		},
		Alias{
			Name: "name2",
			Dir:  "/var/log",
		},
	}
	b, err := yaml.Marshal(cdds)
	assert.NoError(t, err)
	assert.Equal(t, `- name: name1
  dir: ${HOME}/work/project1/actions
- name: name2
  dir: /var/log
`, string(b))
}
