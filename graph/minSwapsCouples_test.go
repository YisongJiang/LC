package graph

import "testing"

func Test_minSwapsCouples(t *testing.T) {
	type args struct {
		row []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				[]int{3, 1, 2, 0},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwapsCouples(tt.args.row); got != tt.want {
				t.Errorf("minSwapsCouples() = %v, want %v", got, tt.want)
			}
		})
	}
}
