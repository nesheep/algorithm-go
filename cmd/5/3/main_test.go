package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		ps []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				ps: []int{2, 3, 5},
			},
			want: 7,
		},
		{
			name: "example 2",
			args: args{
				ps: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.args.ps)
			if got != tt.want {
				t.Errorf("expect %v, but %v", tt.want, got)
			}
		})
	}
}
