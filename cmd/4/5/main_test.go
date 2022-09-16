package main

import "testing"

func TestCount(t *testing.T) {
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
			args: args{n: 575},
			want: 4,
		},
		{
			name: "example 2",
			args: args{n: 3600},
			want: 13,
		},
		{
			name: "example 3",
			args: args{n: 999999999},
			want: 26484,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := count(tt.args.n)
			if got != tt.want {
				t.Errorf("expect %d, but %d", tt.want, got)
			}
		})
	}
}
