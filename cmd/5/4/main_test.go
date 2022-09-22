package main

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		as []int
		w  int
		k  int
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
				k:  2,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				as: []int{1, 5, 9, 10, 13, 21},
				w:  37,
				k:  3,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc(tt.args.as, tt.args.w, tt.args.k)
			if got != tt.want {
				t.Errorf("expect %v, but %v", tt.want, got)
			}
		})
	}
}
