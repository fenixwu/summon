package summon

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name string
		args args
		want *Summoner
	}{
		{`TEST env = "".`, args{""}, &Summoner{true}},
		{`TEST env = "local".`, args{"local"}, &Summoner{true}},
		{`TEST env = "release".`, args{"release"}, &Summoner{false}},
		{`TEST env = "production".`, args{"production"}, &Summoner{false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummoner_IsLocalEnv(t *testing.T) {
	tests := []struct {
		name string
		s    *Summoner
		want bool
	}{
		{"TEST isLocal = true.", &Summoner{true}, true},
		{"TEST isLocal = false.", &Summoner{false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsLocalEnv(); got != tt.want {
				t.Errorf("Summoner.IsLocalEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummoner_Inject(t *testing.T) {
	type args struct {
		k        string
		localVal interface{}
	}
	tests := []struct {
		name string
		s    *Summoner
		args args
	}{
		{"Test inject with local value.", &Summoner{true}, args{"TEST", "a"}},
		{"Test inject with OS value", &Summoner{false}, args{"TEST", "a"}},
	}
	for _, tt := range tests {
		if !tt.s.isEnvLocal {
			os.Setenv(tt.args.k, "b")
		}
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Inject(tt.args.k, tt.args.localVal)
		})
	}
}

func TestSummoner_Get(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		s    *Summoner
		args args
		want *someEnv
	}{
		{"Test env's value(TEST) is from local", &Summoner{true}, args{"TEST"}, &someEnv{"TEST", 1}},
		{"Test env's value(TEST) is from OS", &Summoner{false}, args{"TEST"}, &someEnv{"TEST", "a"}},
		{"Test env's value(TEST) is not exist", &Summoner{false}, args{"TEST2"}, &someEnv{"TEST2", nil}},
	}
	for _, tt := range tests {
		if tt.args.k == "TEST" && tt.s.isEnvLocal {
			tt.s.Inject(tt.args.k, 1)
		}

		if tt.args.k == "TEST" && !tt.s.isEnvLocal {
			os.Setenv(tt.want.key, tt.want.val.(string))
			tt.s.Inject(tt.args.k, 1)
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Get(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Summoner.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValFromOS(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantPanic bool
	}{
		{"Test should panic.", args{"TESTA"}, "a", true},
		{"Test should not panic", args{"TESTB"}, "b", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); tt.wantPanic && r == nil {
					t.Errorf("%s: should have panicked!", tt.name)
				}
			}()
			if tt.args.k == "TESTB" {
				os.Setenv(tt.args.k, tt.want)
			}
			if got := getValFromOS(tt.args.k); got != tt.want {
				fmt.Println(got)
				t.Errorf("getValFromOS() = %v, want %v", got, tt.want)
			}
		})
	}
}
