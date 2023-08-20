use super::Solution;

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        let nums = &mut nums;
        Self::qsort_kth_largest(nums, 0, nums.len(), k as usize).unwrap_or_default()
    }

    fn qsort_kth_largest(arr: &mut Vec<i32>, lo: usize, hi: usize, k: usize) -> Option<i32> {
        if lo >= hi {
            return None;
        }

        let p = Self::partition(arr, lo, hi);

        match arr.len() - p {
            i if i > k => Self::qsort_kth_largest(arr, p + 1, hi, k),
            i if i < k => Self::qsort_kth_largest(arr, lo, p, k),
            _ => arr.get(arr.len() - k).copied(),
        }
    }

    fn partition(arr: &mut [i32], lo: usize, hi: usize) -> usize {
        let pivot = arr[lo];

        let mut left = lo;
        let mut right = hi - 1;

        while left != right {
            let mut i = right;
            while i > left {
                if arr[i] < pivot {
                    arr[left] = arr[i];
                    arr[i] = pivot;
                    break;
                }
                right -= 1;
                i -= 1;
            }
            i = left;
            while i < right {
                if arr[i] > pivot {
                    arr[right] = arr[i];
                    arr[i] = pivot;
                    break;
                }
                left += 1;
                i += 1;
            }
        }

        left
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        assert_eq!(Solution::find_kth_largest(vec![3, 2, 1, 5, 6, 4], 2), 5);
        assert_eq!(
            Solution::find_kth_largest(vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4),
            4
        );
    }
}
