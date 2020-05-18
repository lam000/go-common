package paladin

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type file struct {
	values *Map
}

func NewFile(base string) (Client, error) {
	base = filepath.FromSlash(base)
	fi, err := os.Stat(base)
	if err != nil {
		panic(err)
	}

	var paths []string
	if fi.IsDir() {
		files, err := ioutil.ReadDir(base)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				paths = append(paths, path.Join(base, file.Name()))
			}
		}
	} else {
		paths = append(paths, base)
	}

	values := make(map[string]*Value, len(paths))
	for _, file := range paths {
		if file == "" {
			return nil, errors.New("paladin: path is emtpy")
		}
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		s := string(b)
		values[path.Base(file)] = &Value{val: s, raw: s}
	}
	m := new(Map)
	m.Store(values)
	return &file{values: m}, nil
}

func (f *file) Get(key string) *Value {
	return f.values.Get(key)
}

func (f *file) GetAll() *Map {
	return f.values
}
