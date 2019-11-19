package graph

import (
	"reflect"
	"testing"
)

func Test_findRedundantDirectedConnection(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				[][]int{
					[]int{1, 2},
					[]int{1, 3},
					[]int{2, 3},
				},
			},
			[]int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantDirectedConnection(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantDirectedConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
