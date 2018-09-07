package summon

import (
	"errors"
	"os"
	"sync"

	"github.com/spf13/cast"
)

// env Env data
type env struct {
	Key, Val string
}

var (
	// Local Lcal key and value
	Local       = &env{"ENV", "LOCAL"}
	once        sync.Once
	isLocal     = false
	isInitialed = false
	vals        = map[string]interface{}{}
)

// Init Make sure execute init first before Get____(k string)
func Init() {
	once.Do(func() {
		if e := os.Getenv(Local.Key); e == "" || e == Local.Val {
			isLocal = true
		}
		isInitialed = true
	})
}

func getEnv(k string) string {
	if s := os.Getenv(k); s != "" {
		return s
	}
	panic(`env "` + k + `" is not exist`)
}

// Inject Get a string value from paramer or env and save in map vals
func Inject(k string, localVal interface{}) {
	if isLocal {
		vals[k] = localVal
	}

	vals[k] = getEnv(k)
}

// GetString Get string value from map vals
func GetString(k string) (string, error) {
	if !isInitialed {
		return "", errors.New("summon has not ready yet")
	}

	if s, ok := vals[k]; ok {
		return cast.ToStringE(s)
	}

	return "", errors.New(`env "` + k + `" is not injected`)
}
