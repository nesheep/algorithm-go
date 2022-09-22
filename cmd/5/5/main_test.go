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
				w:  2,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				as: []int{10, 11, 13},
				w:  32,
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				as: []int{2, 5, 7},
				w:  3,
			},
			want: false,
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
