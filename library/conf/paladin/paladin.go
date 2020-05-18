package paladin

import (
	"errors"
	"flag"
	"github.com/lam000/go-common/library/conf/env"
	"strings"
)

var (
	DefaultClient Client
	confPath      string
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

func Init(args ...interface{}) (err error) {
	if confPath != "" {
		path := strings.TrimRight(confPath, "/") + "/" + env.DeployEnv
		DefaultClient, err = NewFile(path)
	} else {
		path, ok := args[0].(string)
		if !ok {
			panic(errors.New("lack of conf path args"))
		}
		DefaultClient, err = NewFile(path)
	}

	if err != nil {
		return
	}

	return
}

func Get(key string) *Value {
	return DefaultClient.Get(key)
}

func GetAll() *Map {
	return DefaultClient.GetAll()
}

func Keys() []string {
	return DefaultClient.GetAll().Keys()
}
