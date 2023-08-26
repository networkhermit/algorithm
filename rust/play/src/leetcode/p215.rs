pub struct Solution {}

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        let length = nums.len();
        Self::quick_select(&mut nums, 0, length, length - k as usize).unwrap_or_default()
    }

    fn quick_select(arr: &mut [i32], lo: usize, hi: usize, k: usize) -> Option<i32> {
        if lo == hi - 1 {
            return match lo == k {
                true => Some(arr[lo]),
                false => None,
            };
        }

        let p = Self::partition(arr, lo, hi);

        match p {
            p if k <= p => Self::quick_select(arr, lo, p + 1, k),
            _ => Self::quick_select(arr, p + 1, hi, k),
        }
    }

    fn partition(arr: &mut [i32], lo: usize, hi: usize) -> usize {
        let pivot = arr[lo + ((hi - 1 - lo) >> 1)];

        let mut i = lo.wrapping_sub(1);
        let mut j = hi;

        loop {
            loop {
                i = i.wrapping_add(1);
                if arr[i] >= pivot {
                    break;
                }
            }
            loop {
                j -= 1;
                if arr[j] <= pivot {
                    break;
                }
            }
            if i >= j {
                return j;
            }
            // (arr[i], arr[j]) = (arr[j], arr[i]);
            arr.swap(i, j);
        }
    }

    pub fn find_kth_largest_inefficient(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        let length = nums.len();
        Self::quick_select_inefficient(&mut nums, 0, length, length - k as usize)
            .unwrap_or_default()
    }

    fn quick_select_inefficient(arr: &mut [i32], lo: usize, hi: usize, k: usize) -> Option<i32> {
        if lo >= hi {
            return None;
        }

        let p = Self::partition_inefficient(arr, lo, hi);

        match p {
            p if k < p => Self::quick_select_inefficient(arr, lo, p, k),
            p if p == k => arr.get(k).copied(),
            _ => Self::quick_select_inefficient(arr, p + 1, hi, k),
        }
    }

    fn partition_inefficient(arr: &mut [i32], lo: usize, hi: usize) -> usize {
        let pivot = arr[lo];

        let mut i = lo;
        let mut j = hi - 1;

        while i != j {
            while j > i {
                if arr[j] < pivot {
                    // (arr[i], arr[j]) = (arr[j], pivot);
                    arr[i] = arr[j];
                    arr[j] = pivot;
                    break;
                }
                j -= 1;
            }
            while i < j {
                if arr[i] > pivot {
                    // (arr[j], arr[i]) = (arr[i], pivot);
                    arr[j] = arr[i];
                    arr[i] = pivot;
                    break;
                }
                i += 1;
            }
        }

        i
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

    #[test]
    fn test_inefficient() {
        assert_eq!(
            Solution::find_kth_largest_inefficient(vec![3, 2, 1, 5, 6, 4], 2),
            5
        );
        assert_eq!(
            Solution::find_kth_largest_inefficient(vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4),
            4
        );
    }
}
