use rand::{rng, seq::SliceRandom};

pub struct Solution {}

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        let length = nums.len();
        Self::quick_select(&mut nums, 0, length, length - k as usize).unwrap_or_default()
    }

    fn quick_select(arr: &mut [i32], lo: usize, hi: usize, k: usize) -> Option<i32> {
        if lo + 1 == hi {
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
            arr.swap(i, j);
        }
    }

    pub fn find_kth_largest_three_way_partition_with_shuffle(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        let length = nums.len();
        let mut rng = rng();
        nums.shuffle(&mut rng);
        Self::quick_select_three_way_partition(&mut nums, 0, length, length - k as usize)
            .unwrap_or_default()
    }

    fn quick_select_three_way_partition(
        arr: &mut [i32],
        lo: usize,
        hi: usize,
        k: usize,
    ) -> Option<i32> {
        if lo + 1 >= hi {
            return arr.get(k).copied();
        }

        let (i, j) = Self::partition_three_way(arr, lo, hi);

        match k {
            k if k <= j && j != usize::MAX => {
                Self::quick_select_three_way_partition(arr, lo, j + 1, k)
            }
            k if k >= i => Self::quick_select_three_way_partition(arr, i, hi, k),
            _ => arr.get(k).copied(),
        }
    }

    #[allow(dead_code)]
    fn quick_select_three_way_partition_iterative(
        arr: &mut [i32],
        lo: usize,
        hi: usize,
        k: usize,
    ) -> Option<i32> {
        let mut lo = lo;
        let mut hi = hi;

        while lo + 1 < hi {
            let (i, j) = Self::partition_three_way(arr, lo, hi);

            match k {
                k if k <= j && j != usize::MAX => {
                    hi = j + 1;
                }
                k if k >= i => {
                    lo = i;
                }
                _ => return arr.get(k).copied(),
            };
        }

        arr.get(k).copied()
    }

    fn partition_three_way(arr: &mut [i32], lo: usize, hi: usize) -> (usize, usize) {
        let mut i = lo.wrapping_sub(1);
        let mut j = hi - 1;

        let mut p = lo.wrapping_sub(1);
        let mut q = hi - 1;

        let pivot = arr[hi - 1];

        loop {
            loop {
                i = i.wrapping_add(1);
                if arr[i] >= pivot {
                    break;
                }
            }
            loop {
                j -= 1;
                if pivot >= arr[j] {
                    break;
                }
                if j == lo {
                    break;
                }
            }
            if i >= j {
                break;
            }
            arr.swap(i, j);
            if arr[i] == pivot {
                p = p.wrapping_add(1);
                arr.swap(p, i);
            }
            if pivot == arr[j] {
                q -= 1;
                arr.swap(j, q);
            }
        }

        arr.swap(i, hi - 1);

        j = i.wrapping_sub(1);
        i += 1;

        let mut k = lo;
        while k < p && p != usize::MAX {
            arr.swap(k, j);
            j -= 1;
            k += 1;
        }
        k = hi - 2;
        while k > q {
            arr.swap(i, k);
            i += 1;
            k -= 1;
        }

        (i, j)
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
                    (arr[i], arr[j]) = (arr[j], pivot);
                    break;
                }
                j -= 1;
            }
            while i < j {
                if arr[i] > pivot {
                    (arr[j], arr[i]) = (arr[i], pivot);
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
    fn test_find_kth_largest() {
        assert_eq!(Solution::find_kth_largest(vec![3, 2, 1, 5, 6, 4], 2), 5);
        assert_eq!(
            Solution::find_kth_largest(vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4),
            4
        );
    }

    #[test]
    fn test_find_kth_largest_three_way_partition() {
        assert_eq!(
            Solution::find_kth_largest_three_way_partition_with_shuffle(vec![3, 2, 1, 5, 6, 4], 2),
            5
        );
        assert_eq!(
            Solution::find_kth_largest_three_way_partition_with_shuffle(
                vec![3, 2, 3, 1, 2, 4, 5, 5, 6],
                4
            ),
            4
        );
    }

    #[test]
    fn test_find_kth_largest_inefficient() {
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
