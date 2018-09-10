package summon

import (
	"os"
)

var (
	env, local = "ENV", "local"
	vals       = map[string]interface{}{}
)

type Summoner struct {
	isEnvLocal bool
}

// New a "Summoner" and set "isLocal" once.
func New(env string) *Summoner {
	isLocal := false
	if env == "" || env == local {
		isLocal = true
	}

	return &Summoner{
		isEnvLocal: isLocal,
	}
}

// IsLocalEnv tells that is program running locally.
func (s *Summoner) IsLocalEnv() bool {
	return s.isEnvLocal
}

// Inject value from local or OS into map.
func (s *Summoner) Inject(k string, localVal interface{}) {
	if s.isEnvLocal {
		vals[k] = localVal
		return
	}

	vals[k] = getValFromOS(k)
}

// Get a value from map.
func (s *Summoner) Get(k string) *SomeEnv {
	if v, ok := vals[k]; ok {
		return &SomeEnv{k, v}
	}

	return &SomeEnv{k, nil}
}

// Get value by using "os.Getenv(key string)".
func getValFromOS(k string) string {
	if s := os.Getenv(k); s != "" {
		return s
	}
	panic(`env "` + k + `" is not exist`)
}
