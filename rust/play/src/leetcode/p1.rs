use std::collections::HashMap;

pub struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = HashMap::new();

        for (j, n) in nums.iter().enumerate() {
            let m = target - n;
            if let Some(&i) = map.get(&m) {
                return vec![i as i32, j as i32];
            } else {
                map.insert(n, j);
            }
        }

        Default::default()
    }

    pub fn two_sum_brute_force(nums: Vec<i32>, target: i32) -> Vec<i32> {
        for i in 0..nums.len() {
            for j in i + 1..nums.len() {
                if nums[i] + nums[j] == target {
                    return vec![i as i32, j as i32];
                }
            }
        }

        Default::default()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        assert_eq!(Solution::two_sum(vec![2, 7, 11, 15], 9), vec![0, 1]);
        assert_eq!(Solution::two_sum(vec![3, 2, 4], 6), vec![1, 2]);
        assert_eq!(Solution::two_sum(vec![3, 3], 6), vec![0, 1]);
    }

    #[test]
    fn test_brute_force() {
        assert_eq!(
            Solution::two_sum_brute_force(vec![2, 7, 11, 15], 9),
            vec![0, 1]
        );
        assert_eq!(Solution::two_sum_brute_force(vec![3, 2, 4], 6), vec![1, 2]);
        assert_eq!(Solution::two_sum_brute_force(vec![3, 3], 6), vec![0, 1]);
    }
}
