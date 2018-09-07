package summon

import (
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cast"
)

var (
	// EnvKey Environment Key
	EnvKey = "ENV"
	// EnvLocal Local environment Value
	EnvLocal   = "local"
	once       sync.Once
	isEnvLocal = false
	vals       = map[string]interface{}{}
)

// Summon Summon
type Summon struct{}

// New New a summon, set isEnvLocal once
func New() *Summon {
	once.Do(func() {
		if e := os.Getenv(EnvKey); e == "" || e == EnvLocal {
			isEnvLocal = true
		}
	})

	return &Summon{}
}

// Get value by using "os.Getenv(key string)"
func getValFromOS(k string) string {
	if s := os.Getenv(k); s != "" {
		return s
	}
	panic(`env "` + k + `" is not exist`)
}

// Inject Inject value which is from local or OS into map
func (s *Summon) Inject(k string, localVal interface{}) {
	if isEnvLocal {
		vals[k] = localVal
		return
	}

	vals[k] = getValFromOS(k)
}

// Get Get value from map
func (s *Summon) Get(k string) *SomeEnv {
	if v, ok := vals[k]; ok {
		return &SomeEnv{k, v}
	}

	return &SomeEnv{k, nil}
}

// SomeEnv Get value struct
type SomeEnv struct {
	key string
	val interface{}
}

// ToTime Parse value to time
func (s *SomeEnv) ToTime() (tim time.Time, err error) {
	if s.val == nil {
		return time.Now(), generateNotInjectedYetErr(s.key)
	}
	return cast.ToTimeE(s.val)
}

// ToDuration Parse value to duration
func (s *SomeEnv) ToDuration() (d time.Duration, err error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToDurationE(s.val)
}

// ToBool Parse value to bool
func (s *SomeEnv) ToBool() (bool, error) {
	if s.val == nil {
		return false, generateNotInjectedYetErr(s.key)
	}
	return cast.ToBoolE(s.val)
}

// ToFloat64 Parse value to float64
func (s *SomeEnv) ToFloat64() (float64, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToFloat64E(s.val)
}

// ToFloat32 Parse value to float32
func (s *SomeEnv) ToFloat32() (float32, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToFloat32E(s.val)
}

// ToInt64 Parse value to int64
func (s *SomeEnv) ToInt64() (int64, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToInt64E(s.val)
}

// ToInt32 Parse value to int32
func (s *SomeEnv) ToInt32() (int32, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToInt32E(s.val)
}

// ToInt16 Parse value to int16
func (s *SomeEnv) ToInt16() (int16, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToInt16E(s.val)
}

// ToInt8 Parse value to int8
func (s *SomeEnv) ToInt8() (int8, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToInt8E(s.val)
}

// ToInt Parse value to int
func (s *SomeEnv) ToInt() (int, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToIntE(s.val)
}

// ToUint Parse value to uint
func (s *SomeEnv) ToUint() (uint, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToUintE(s.val)
}

// ToUint64 Parse value to uint64
func (s *SomeEnv) ToUint64() (uint64, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToUint64E(s.val)
}

// ToUint32 Parse value to uint32
func (s *SomeEnv) ToUint32() (uint32, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToUint32E(s.val)
}

// ToUint16 Parse value to uint16
func (s *SomeEnv) ToUint16() (uint16, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToUint16E(s.val)
}

// ToUint8 Parse value to uint8
func (s *SomeEnv) ToUint8() (uint8, error) {
	if s.val == nil {
		return 0, generateNotInjectedYetErr(s.key)
	}
	return cast.ToUint8E(s.val)
}

// ToString Parse value to string
func (s *SomeEnv) ToString() (string, error) {
	if s.val == nil {
		return "", generateNotInjectedYetErr(s.key)
	}
	return cast.ToStringE(s.val)
}

// ToStringMapString Parse value to stringmapstring
func (s *SomeEnv) ToStringMapString() (map[string]string, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToStringMapStringE(s.val)
}

// ToStringMapStringSlice Parse value to stringmapstringslice
func (s *SomeEnv) ToStringMapStringSlice() (map[string][]string, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToStringMapStringSliceE(s.val)
}

// ToStringMapBool Parse value to stringmapbool
func (s *SomeEnv) ToStringMapBool() (map[string]bool, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToStringMapBoolE(s.val)
}

// ToStringMap Parse value to stringmap
func (s *SomeEnv) ToStringMap() (map[string]interface{}, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToStringMapE(s.val)
}

// ToSlice Parse value to slice
func (s *SomeEnv) ToSlice() ([]interface{}, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToSliceE(s.val)
}

// ToBoolSlice Parse value to boolslice
func (s *SomeEnv) ToBoolSlice() ([]bool, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToBoolSliceE(s.val)
}

// ToStringSlice Parse value to stringslice
func (s *SomeEnv) ToStringSlice() ([]string, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToStringSliceE(s.val)
}

// ToIntSlice Parse value to intslice
func (s *SomeEnv) ToIntSlice() ([]int, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToIntSliceE(s.val)
}

// ToDurationSlice Parse value to durationslice
func (s *SomeEnv) ToDurationSlice() ([]time.Duration, error) {
	if s.val == nil {
		return nil, generateNotInjectedYetErr(s.key)
	}
	return cast.ToDurationSliceE(s.val)
}

func generateNotInjectedYetErr(k string) error {
	var s strings.Builder
	s.WriteString(`"`)
	s.WriteString(k)
	s.WriteString(`" s not injected yet`)
	return errors.New(s.String())
}
