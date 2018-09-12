package summon

import (
	"errors"
	"strings"
	"time"

	"github.com/spf13/cast"
)

type someEnv struct {
	key string
	val interface{}
}

// ToTime parse value to time type.
func (s *someEnv) ToTime() (tim time.Time, err error) {
	if s.val == nil {
		err = s.notInjectedYetError()
		return
	}
	tim, err = cast.ToTimeE(s.val)
	return
}

// ToDuration parse value to duration type.
func (s *someEnv) ToDuration() (d time.Duration, err error) {
	if s.val == nil {
		err = s.notInjectedYetError()
		return
	}
	d, err = cast.ToDurationE(s.val)
	return
}

// ToBool parse value to bool type.
func (s *someEnv) ToBool() (bool, error) {
	if s.val == nil {
		return false, s.notInjectedYetError()
	}
	return cast.ToBoolE(s.val)
}

// ToFloat64 parse value to float64 type.
func (s *someEnv) ToFloat64() (float64, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToFloat64E(s.val)
}

// ToFloat32 parse value to float32 type.
func (s *someEnv) ToFloat32() (float32, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToFloat32E(s.val)
}

// ToInt64 parse value to int64 type.
func (s *someEnv) ToInt64() (int64, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToInt64E(s.val)
}

// ToInt32 parse value to int32 type.
func (s *someEnv) ToInt32() (int32, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToInt32E(s.val)
}

// ToInt16 parse value to int16 type.
func (s *someEnv) ToInt16() (int16, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToInt16E(s.val)
}

// ToInt8 parse value to int8 type.
func (s *someEnv) ToInt8() (int8, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToInt8E(s.val)
}

// ToInt parse value to int type.
func (s *someEnv) ToInt() (int, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToIntE(s.val)
}

// ToUint parse value to uint type.
func (s *someEnv) ToUint() (uint, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToUintE(s.val)
}

// ToUint64 parse value to uint64 type.
func (s *someEnv) ToUint64() (uint64, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToUint64E(s.val)
}

// ToUint32 parse value to uint32 type.
func (s *someEnv) ToUint32() (uint32, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToUint32E(s.val)
}

// ToUint16 parse value to uint16 type.
func (s *someEnv) ToUint16() (uint16, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToUint16E(s.val)
}

// ToUint8 parse value to uint8 type.
func (s *someEnv) ToUint8() (uint8, error) {
	if s.val == nil {
		return 0, s.notInjectedYetError()
	}
	return cast.ToUint8E(s.val)
}

// ToString parse value to string type.
func (s *someEnv) ToString() (string, error) {
	if s.val == nil {
		return "", s.notInjectedYetError()
	}
	return cast.ToStringE(s.val)
}

// ToStringMapString parse value to map[string]string type.
func (s *someEnv) ToStringMapString() (map[string]string, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToStringMapStringE(s.val)
}

// ToStringMapStringSlice parse value to map[string][]string type.
func (s *someEnv) ToStringMapStringSlice() (map[string][]string, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToStringMapStringSliceE(s.val)
}

// ToStringMapBool parse value to map[string]bool type.
func (s *someEnv) ToStringMapBool() (map[string]bool, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToStringMapBoolE(s.val)
}

// ToStringMap parse value to map[string]interface{} type.
func (s *someEnv) ToStringMap() (map[string]interface{}, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToStringMapE(s.val)
}

// ToSlice parse value to slice type.
func (s *someEnv) ToSlice() ([]interface{}, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToSliceE(s.val)
}

// ToBoolSlice parse value to bool slice type.
func (s *someEnv) ToBoolSlice() ([]bool, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToBoolSliceE(s.val)
}

// ToStringSlice parse value to string slice type.
func (s *someEnv) ToStringSlice() ([]string, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToStringSliceE(s.val)
}

// ToIntSlice parse value to int slice type.
func (s *someEnv) ToIntSlice() ([]int, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToIntSliceE(s.val)
}

// ToDurationSlice parse value to duration slice type.
func (s *someEnv) ToDurationSlice() ([]time.Duration, error) {
	if s.val == nil {
		return nil, s.notInjectedYetError()
	}
	return cast.ToDurationSliceE(s.val)
}

func (s *someEnv) notInjectedYetError() error {
	var e strings.Builder
	e.WriteString(`"`)
	e.WriteString(s.key)
	e.WriteString(`" s not injected yet`)
	return errors.New(e.String())
}
