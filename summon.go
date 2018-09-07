package summon

import (
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cast"
)

// Summon Summon
type Summon struct{}

var (
	// Local Lcal key and value
	Local = struct {
		Key, Val string
	}{"ENV", "LOCAL"}
	once    sync.Once
	isLocal = false
	vals    = map[string]interface{}{}
)

// New New Summon
func New() *Summon {
	return &Summon{}
}

// Init Make sure execute init first
func Init() {
	once.Do(func() {
		if e := os.Getenv(Local.Key); e == "" || e == Local.Val {
			isLocal = true
		}
	})
}

func getEnv(k string) (string, error) {
	s := os.Getenv(k)

	if s == "" {
		return "", errors.New(`env "` + k + `" is not exist`)
	}

	return s, nil
}

// Inject Get a string value from paramer or env and save in map vals
func Inject(k string, localVal interface{}) {
	upper := strings.ToUpper(k)
	if isLocal {
		vals[upper] = localVal
	}

	v, err := getEnv(upper)

	if err != nil {
		panic(err.Error())
	}

	vals[upper] = v
}

// GetString Get string value from map vals
func GetString(k string) (string, error) {
	if s, ok := vals[strings.ToUpper(k)]; ok {
		return cast.ToStringE(s)
	}
	return "", errors.New(`env "` + k + `" is not injected`)
}
