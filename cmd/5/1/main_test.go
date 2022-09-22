package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		as [][3]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{[][3]int{
				{10, 40, 70},
				{20, 50, 80},
				{30, 60, 90},
			}},
			want: 210,
		},
		{
			name: "example 2",
			args: args{[][3]int{
				{100, 10, 1},
			}},
			want: 100,
		},
		{
			name: "example 3",
			args: args{[][3]int{
				{6, 7, 8},
				{8, 8, 3},
				{2, 5, 2},
				{7, 8, 6},
				{4, 6, 8},
				{2, 3, 4},
				{7, 5, 1},
			}},
			want: 46,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.args.as)
			if got != tt.want {
				t.Errorf("expect %d, but %d", tt.want, got)
			}
		})
	}
}
