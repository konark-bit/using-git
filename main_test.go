package main

import "testing"

func Test_heightSum(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				heights: []int{3, 7, 8},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heightSum(tt.args.heights); got != tt.want {
				t.Errorf("heightSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
