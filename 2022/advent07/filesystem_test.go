package main

import "testing"

func Test_popdirname(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "3", args: args{path: "/foo/bar/baz/"}, want: "/foo/bar/"},
		{name: "2", args: args{path: "/foo/bar/"}, want: "/foo/"},
		{name: "1", args: args{path: "/foo/"}, want: "/"},
		{name: "0", args: args{path: "/"}, want: "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := popdirname(tt.args.path); got != tt.want {
				t.Errorf("popdirname() = %v, want %v", got, tt.want)
			}
		})
	}
}
