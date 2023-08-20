use super::Solution;

impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        let mut mat = obstacle_grid;
        for i in 0..mat.len() {
            for j in 0..mat[i].len() {
                if i == 0 && j == 0 {
                    match mat[i][j] {
                        0 => {
                            mat[i][j] = 1;
                            continue;
                        }
                        _ => {
                            return 0;
                        }
                    }
                }
                match mat[i][j] {
                    0 => {
                        if i > 0 {
                            mat[i][j] += mat[i - 1][j];
                        }
                        if j > 0 {
                            mat[i][j] += mat[i][j - 1];
                        }
                    }
                    _ => {
                        mat[i][j] = 0;
                    }
                }
            }
        }
        match mat.last() {
            Some(vec) => match vec.last() {
                Some(&n) => n,
                None => 0,
            },
            None => 0,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        assert_eq!(
            Solution::unique_paths_with_obstacles(vec![
                vec![0, 0, 0],
                vec![0, 1, 0],
                vec![0, 0, 0]
            ]),
            2
        );
        assert_eq!(
            Solution::unique_paths_with_obstacles(vec![vec![0, 1], vec![0, 0],]),
            1
        );
    }
}
