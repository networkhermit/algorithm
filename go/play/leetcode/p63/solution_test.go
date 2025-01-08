package p63

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	var sample = []struct {
		obstacleGrid [][]int
		expected     int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
	}

	for _, tt := range sample {
		actual := uniquePathsWithObstacles(tt.obstacleGrid)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("%s(%v) returned %d, expect %d", t.Name(), tt.obstacleGrid, actual, tt.expected)
		}
	}
}
