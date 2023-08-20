package p63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	mat := obstacleGrid
	for i := range mat {
		for j := range mat[i] {
			if i == 0 && j == 0 {
				if mat[i][j] == 0 {
					mat[i][j] = 1
					continue
				} else {
					return 0
				}
			}
			if mat[i][j] == 0 {
				if i > 0 {
					mat[i][j] += mat[i-1][j]
				}
				if j > 0 {
					mat[i][j] += mat[i][j-1]
				}
			} else {
				mat[i][j] = 0
			}
		}
	}
	if m := len(mat); m > 0 {
		if n := len(mat[m-1]); n > 0 {
			return mat[m-1][n-1]
		}
	}
	return 0
}
