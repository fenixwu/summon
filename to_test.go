package summon

import (
	"reflect"
	"testing"
	"time"
)

func TestSomeEnv_ToTime(t *testing.T) {
	var def time.Time
	now := time.Now()
	tests := []struct {
		name    string
		s       *SomeEnv
		wantTim time.Time
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_TIME_A", now}, now, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_TIME_B", "now"}, def, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_TIME_C", nil}, def, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTim, err := tt.s.ToTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTim, tt.wantTim) {
				t.Errorf("SomeEnv.ToTime() = %v, want %v", gotTim, tt.wantTim)
			}
		})
	}
}

func TestSomeEnv_ToDuration(t *testing.T) {
	var def time.Duration
	tests := []struct {
		name    string
		s       *SomeEnv
		wantD   time.Duration
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_DURATION_A", 1 * time.Second}, 1 * time.Second, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_DURATION_B", "now"}, def, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_DURATION_C", nil}, def, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := tt.s.ToDuration()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("SomeEnv.ToDuration() = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func TestSomeEnv_ToBool(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    bool
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_BOOL_A", "true"}, true, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_BOOL_B", "a"}, false, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_BOOL_C", nil}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToFloat64(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    float64
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_FLOAT64_A", 1.1}, 1.1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_FLOAT64_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_FLOAT64_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToFloat64()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToFloat32(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    float32
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_FLOAT32_A", 1.1}, 1.1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_FLOAT32_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_FLOAT32_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToFloat32()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToInt64(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    int64
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_INT64_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_INT64_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_INT64_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToInt64()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToInt32(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    int32
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_INT32_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_INT32_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_INT32_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToInt32()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToInt16(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    int16
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_INT16_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_INT16_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_INT16_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToInt16()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToInt8(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    int8
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_INT8_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_INT8_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_INT8_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToInt8()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToInt(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    int
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_INT_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_INT_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_INT_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToUint(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    uint
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_UINT_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_UINT_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_UINT_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToUint()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToUint64(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    uint64
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_UINT64_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_UINT64_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_UINT64_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToUint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToUint32(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    uint32
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_UINT32_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_UINT32_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_UINT32_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToUint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToUint16(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    uint16
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_UINT16_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_UINT16_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_UINT16_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToUint16()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToUint8(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    uint8
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_UINT8_A", 1}, 1, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_UINT8_B", "a"}, 0, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_UINT8_C", nil}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToUint8()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToString(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    string
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_STRING_A", 1}, "1", false},
		{"Test return cast error", &SomeEnv{"TEST_TO_STRING_B", &Summoner{}}, "", true},
		{"Test return nil error", &SomeEnv{"TEST_TO_STRING_C", nil}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToString()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SomeEnv.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToStringMapString(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    map[string]string
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_STRING_MAP_STRING_A", map[string]string{"test": "a"}}, map[string]string{"test": "a"}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_STRING_MAP_STRING_B", ""}, map[string]string{}, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_STRING_MAP_STRING_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToStringMapString()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToStringMapString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToStringMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToStringMapStringSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    map[string][]string
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_STRING_MAP_STRING_SLICE_A", map[string][]string{"test": []string{"a"}}}, map[string][]string{"test": []string{"a"}}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_STRING_MAP_STRING_SLICE_B", ""}, map[string][]string{}, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_STRING_MAP_STRING_SLICE_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToStringMapStringSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToStringMapStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToStringMapStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToStringMapBool(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    map[string]bool
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_STRING_MAP_BOOL_A", map[string]bool{"test": true}}, map[string]bool{"test": true}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_STRING_MAP_BOOL_B", ""}, map[string]bool{}, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_STRING_MAP_BOOL_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToStringMapBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToStringMapBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToStringMapBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToStringMap(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    map[string]interface{}
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_STRING_MAP_A", map[string]interface{}{"test": "t"}}, map[string]interface{}{"test": "t"}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_STRING_MAP_B", ""}, map[string]interface{}{}, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_STRING_MAP_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToStringMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToStringMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    []interface{}
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_SLICE_A", []interface{}{1, 3}}, []interface{}{1, 3}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_SLICE_B", ""}, nil, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_SLICE_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToBoolSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    []bool
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_BOOL_SLICE_A", []bool{true, false}}, []bool{true, false}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_BOOL_SLICE_B", ""}, []bool{}, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_BOOL_SLICE_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToBoolSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToBoolSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToBoolSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToStringSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    []string
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_STRING_SLICE_A", []string{"true", "false"}}, []string{"true", "false"}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_STRING_SLICE_B", &Summoner{}}, nil, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_STRING_SLICE_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToStringSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToIntSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    []int
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_INT_SLICE_A", []int{1, 2}}, []int{1, 2}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_INT_SLICE_B", &Summoner{}}, []int{}, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_INT_SLICE_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToIntSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToIntSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeEnv_ToDurationSlice(t *testing.T) {
	tests := []struct {
		name    string
		s       *SomeEnv
		want    []time.Duration
		wantErr bool
	}{
		{"Test return correctly", &SomeEnv{"TEST_TO_DURATION_SLICE_A", []time.Duration{1 * time.Second, 2 * time.Second}}, []time.Duration{1 * time.Second, 2 * time.Second}, false},
		{"Test return cast error", &SomeEnv{"TEST_TO_DURATION_SLICE_B", &Summoner{}}, []time.Duration{}, true},
		{"Test return nil error", &SomeEnv{"TEST_TO_DURATION_SLICE_C", nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToDurationSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeEnv.ToDurationSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SomeEnv.ToDurationSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
