package main

import "testing"

func TestTribonacci(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{n: 5},
			want: 4,
		},
		{
			name: "example 2",
			args: args{n: 9},
			want: 44,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tribonacci(tt.args.n)
			if got != tt.want {
				t.Errorf("expect %d, but %d", tt.want, got)
			}
		})
	}
}
