package repo

import (
	"github.com/koooyooo/cdd/model"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once
var singleton repo

func Instance() repo {
	once.Do(func() {
		r, err := newRepo()
		if err != nil {
			log.Fatalf("fail in loading repo: %v", err)
		}
		singleton = r
	})
	return singleton
}

type repo interface {
	Init() error
	List() ([]*model.Alias, error)
	Get(name string) (*model.Alias, bool, error)
	Add(alias *model.Alias) error
	Remove(name string) error
}

func newRepo() (repo, error) {
	repo := new(repoImpl)
	if err := repo.Init(); err != nil {
		return nil, err
	}
	return repo, nil
}

type repoImpl struct {
	cache []*model.Alias
}

func (r *repoImpl) Init() error {
	as, err := load()
	if err != nil {
		return err
	}
	r.cache = make([]*model.Alias, len(as))
	for i, _ := range as {
		//r.cache[i] = &a // iterator 一時変数のアドレスを取得するのは危険
		r.cache[i] = &as[i]
	}
	return nil
}

func (r *repoImpl) List() ([]*model.Alias, error) {
	return r.cache, nil
}

func (r *repoImpl) Get(name string) (*model.Alias, bool, error) {
	for _, a := range r.cache {
		if a.Name == name {
			return a, true, nil
		}
	}
	return nil, false, nil
}

func (r *repoImpl) Add(alias *model.Alias) error {
	r.cache = append(r.cache, alias)
	raw, err := toRaw(r.cache)
	if err != nil {
		return err
	}
	return store(raw)
}

func (r *repoImpl) Remove(name string) error {
	var removed []*model.Alias
	for _, a := range r.cache {
		if a.Name == name {
			continue
		}
		removed = append(removed, a)
	}
	raw, err := toRaw(removed)
	if err != nil {
		return err
	}
	if err := store(raw); err != nil {
		return err
	}
	r.cache = removed
	return nil
}

func toRaw(as []*model.Alias) ([]model.Alias, error) {
	var data = make([]model.Alias, len(as))
	for i, c := range as {
		data[i] = *c
	}
	return data, nil
}

func store(as []model.Alias) error {
	b, err := yaml.Marshal(as)
	if err != nil {
		return err
	}
	path, err := cddPath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 655)
}

func load() ([]model.Alias, error) {
	path, err := cddPath()
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var as []model.Alias
	if err := yaml.Unmarshal(b, &as); err != nil {
		return nil, err
	}
	return as, nil
}

func cddPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, ".cdd.yaml"), nil
}
