package model

import "github.com/koooyooo/cdd/common"

type Alias struct {
	Name string `yaml:"name"`
	Dir  string `yaml:"dir"`
}

func (a *Alias) ReplacedDir() (string, error) {
	return common.Replace4Get(a.Dir)
}
