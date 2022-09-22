package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		as []int
		w  int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				as: []int{1, 3, 5},
				w:  4,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				as: []int{1, 4, 8},
				w:  6,
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				as: []int{1, 4, 8, 5, 10, 13, 11, 30, 91, 2},
				w:  31,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.args.as, tt.args.w)
			if got != tt.want {
				t.Errorf("expect %v, but %v", tt.want, got)
			}
		})
	}
}
