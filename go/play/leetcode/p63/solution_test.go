package p63

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	sample := []struct {
		obstacleGrid [][]int
		expected     int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
	}

	for _, tc := range sample {
		actual := uniquePathsWithObstacles(tc.obstacleGrid)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("%s(%v) returned %d, expect %d", t.Name(), tc.obstacleGrid, actual, tc.expected)
		}
	}
}
