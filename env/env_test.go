package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		key string
		def []string
	}
	tests := []struct {
		name   string
		args   args
		want   string
		mocker func(args)
	}{
		{
			name:   "the env key is missing and the default value is empty",
			args:   args{key: "KEY", def: []string{}},
			want:   "",
			mocker: func(a args) {},
		},
		{
			name:   "env key is missing but default value has one element",
			args:   args{key: "KEY", def: []string{"NOW"}},
			want:   "NOW",
			mocker: func(a args) {},
		},
		{
			name: "the key of env is exist but default value is empty",
			args: args{key: "IM_IN", def: []string{}},
			want: "-",
			mocker: func(a args) {
				os.Setenv("IM_IN", "-")
			},
		},
		{
			name:   "the key of env is exist but value is empty but default value has one element",
			args:   args{key: "IM_OUT", def: []string{"-"}},
			want:   "-",
			mocker: func(a args) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args)
			if got := Get(tt.args.key, tt.args.def...); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		key string
		def []int
	}
	tests := []struct {
		name   string
		args   args
		want   int
		mocker func(args)
	}{
		{
			name:   "the env key is missing and the default value is empty",
			args:   args{key: "THENUM", def: []int{}},
			want:   0,
			mocker: func(a args) {},
		},
		{
			name:   "env key is missing but default value has one element",
			args:   args{key: "THENUM", def: []int{100}},
			want:   100,
			mocker: func(a args) {},
		},
		{
			name: "the key of env is exist but default value is empty",
			args: args{key: "THENUM1", def: []int{}},
			want: 0,
			mocker: func(a args) {
				os.Setenv("THENUM1", "-")
			},
		},
		{
			name: "the key of env is exist but default value is empty",
			args: args{key: "THENUM1", def: []int{90}},
			want: 90,
			mocker: func(a args) {
				os.Setenv("THENUM1", "-")
			},
		},
		{
			name: "the key of env is exist but value is empty but default value has one element",
			args: args{key: "THENUM0", def: []int{900}},
			want: 10,
			mocker: func(a args) {
				os.Setenv("THENUM0", "10")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args)
			if got := GetInt(tt.args.key, tt.args.def...); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
