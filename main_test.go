package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Multiple ways of cutting tree",
			args: args{
				A: []int{3, 4, 5, 3, 7},
			},
			want: 3,
		},
		{
			name: "No way of cutting tree",
			args: args{
				A: []int{1, 2, 3, 4},
			},
			want: -1,
		},
		{
			name: "Already asthetically pleasing",
			args: args{
				A: []int{1, 3, 1, 2},
			},
			want: 0,
		},
		// Can add more cases here to test the implementation.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
