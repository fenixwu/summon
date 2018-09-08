package summon

import (
	"os"
	"sync"
)

var (
	// env is environment variable's Key.
	env = "ENV"
	// envLocal Local environment's value.
	local   = "local"
	once    sync.Once
	isLocal = false
	vals    = map[string]interface{}{}
)

type Summoner struct {
	isLocal bool
}

// New a "Summoner" and set "isLocal" once.
func New() *Summoner {
	once.Do(func() {
		if e := os.Getenv(env); e == "" || e == local {
			isLocal = true
		}
	})

	return &Summoner{
		isLocal: isLocal,
	}
}

// IsLocalEnv tells that is program running locally.
func (s *Summoner) IsLocalEnv() bool {
	return s.isLocal
}

// Inject value from local or OS into map.
func (s *Summoner) Inject(k string, localVal interface{}) {
	if isLocal {
		vals[k] = localVal
		return
	}

	vals[k] = getValFromOS(k)
}

// Get a value from map.
func (s *Summoner) Get(k string) *someEnv {
	if v, ok := vals[k]; ok {
		return &someEnv{k, v}
	}

	return &someEnv{k, nil}
}

// Get value by using "os.Getenv(key string)".
func getValFromOS(k string) string {
	if s := os.Getenv(k); s != "" {
		return s
	}
	panic(`env "` + k + `" is not exist`)
}
