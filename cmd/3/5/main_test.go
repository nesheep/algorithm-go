package main

import "testing"

func TestHowManyTimes(t *testing.T) {
	type args struct {
		a int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "want 0",
			args: args{
				a: 7,
			},
			want: 0,
		},
		{
			name: "want 1",
			args: args{
				a: 6,
			},
			want: 1,
		},
		{
			name: "want 5",
			args: args{
				a: 96,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := howManyTimes(tt.args.a)
			if got != tt.want {
				t.Errorf("expect %d, but %d", tt.want, got)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	type args struct {
		as []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "want 0",
			args: args{
				as: []int{1, 8, 16},
			},
			want: 0,
		},
		{
			name: "want 1",
			args: args{
				as: []int{2, 8, 16},
			},
			want: 1,
		},
		{
			name: "want 5",
			args: args{
				as: []int{32, 64, 128},
			},
			want: 5,
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
