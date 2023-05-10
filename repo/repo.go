package repo

import (
	"github.com/koooyooo/cdd/common"
	"github.com/koooyooo/cdd/model"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"sync"
)

var once sync.Once
var singleton Repo

func Instance() Repo {
	once.Do(func() {
		r, err := newRepo()
		if err != nil {
			log.Fatalf("fail in loading Repo: %v", err)
		}
		singleton = r
	})
	return singleton
}

type Repo interface {
	Init() error
	List() ([]*model.Alias, error)
	Get(name string) (*model.Alias, bool, error)
	Add(alias *model.Alias) error
	Remove(name string) error
	Move(name string, num int) error
}

func newRepo() (Repo, error) {
	repo := new(repoImpl)
	if err := repo.reload(); err != nil {
		return nil, err
	}
	return repo, nil
}

type repoImpl struct {
	cache []*model.Alias
}

func (r *repoImpl) Init() error {
	defaultAliases := []model.Alias{
		{
			Name: "home",
			Dir:  "${HOME}",
		},
		{
			Name: "docs",
			Dir:  "${HOME}/Documents",
		},
	}
	return store(defaultAliases)
}

func (r *repoImpl) reload() error {
	path, err := common.CDDPath()
	if err != nil {
		log.Fatal(err)
	}
	if !common.Exists(path) {
		if err := r.Init(); err != nil {
			return err
		}
	}
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
	// 同名の既存削除
	idx := -1
	for i, a := range r.cache {
		if a.Name == alias.Name {
			idx = i
		}
	}
	if idx != -1 {
		r.cache = append(r.cache[:idx], r.cache[idx+1:]...)
	}
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

//func (r *repoImpl) Up(name string, num int) error {
//	idx := -1
//	for i, a := range r.cache {
//		if a.Name == name {
//			idx = i
//		}
//	}
//	if idx == -1 {
//		if len(r.cache) <= num {
//			return nil
//		}
//		var err error
//		idx, err = strconv.Atoi(name)
//		if err != nil {
//			return nil
//		}
//	}
//	for i := idx; i > 0; i-- {
//		if num == 0 {
//			break
//		}
//		r.cache[i], r.cache[i-1] = r.cache[i-1], r.cache[i]
//		num--
//	}
//	raw, err := toRaw(r.cache)
//	if err != nil {
//		return err
//	}
//	return store(raw)
//}
//
//func (r *repoImpl) Down(name string, num int) error {
//	return nil // TODO implement
//}

func (r *repoImpl) Move(name string, num int) error {
	idx := -1
	for i, a := range r.cache {
		if a.Name == name {
			idx = i
		}
	}
	if idx == -1 {
		if len(r.cache) <= num {
			return nil
		}
		var err error
		idx, err = strconv.Atoi(name)
		if err != nil {
			return nil
		}
	}
	if num > 0 {
		for i := idx; i > 0; i-- {
			if num == 0 {
				break
			}
			r.cache[i], r.cache[i-1] = r.cache[i-1], r.cache[i]
			num--
		}
	}
	if num < 0 {
		for i := idx; i < len(r.cache)-1; i++ {
			if num == 0 {
				break
			}
			r.cache[i], r.cache[i+1] = r.cache[i+1], r.cache[i]
			num++
		}
	}
	raw, err := toRaw(r.cache)
	if err != nil {
		return err
	}
	return store(raw)
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
	path, err := common.CDDPath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0655)
}

func load() ([]model.Alias, error) {
	path, err := common.CDDPath()
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
