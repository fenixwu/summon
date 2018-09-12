package summon

import (
	"os"
	"time"
)

var (
	env, local = "ENV", "local"
	vals       = map[string]interface{}{}
)

type To interface {
	ToTime() (tim time.Time, err error)
	ToDuration() (d time.Duration, err error)
	ToBool() (bool, error)
	ToFloat64() (float64, error)
	ToFloat32() (float32, error)
	ToInt64() (int64, error)
	ToInt32() (int32, error)
	ToInt16() (int16, error)
	ToInt8() (int8, error)
	ToInt() (int, error)
	ToUint() (uint, error)
	ToUint64() (uint64, error)
	ToUint32() (uint32, error)
	ToUint16() (uint16, error)
	ToUint8() (uint8, error)
	ToString() (string, error)
	ToStringMapString() (map[string]string, error)
	ToStringMapStringSlice() (map[string][]string, error)
	ToStringMapBool() (map[string]bool, error)
	ToStringMap() (map[string]interface{}, error)
	ToSlice() ([]interface{}, error)
	ToBoolSlice() ([]bool, error)
	ToStringSlice() ([]string, error)
	ToIntSlice() ([]int, error)
	ToDurationSlice() ([]time.Duration, error)
}

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
func (s *Summoner) Get(k string) To {
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
