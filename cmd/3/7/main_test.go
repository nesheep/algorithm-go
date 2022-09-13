package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{s: "125"},
			want: 176,
		},
		{
			name: "example 1",
			args: args{s: "9999999999"},
			want: 12656242944,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.args.s)
			if got != tt.want {
				t.Errorf("expect %d, but %d", tt.want, got)
			}
		})
	}
}
