package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		k, n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				k: 2,
				n: 2,
			},
			want: 6,
		},
		{
			name: "example 2",
			args: args{
				k: 5,
				n: 15,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.args.k, tt.args.n)
			if got != tt.want {
				t.Errorf("expect %d, but %d", tt.want, got)
			}
		})
	}
}
