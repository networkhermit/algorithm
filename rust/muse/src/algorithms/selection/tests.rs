use crate::util::{random_factory, sequence_builder};

pub(crate) fn derive<'a>(
    f: &'a dyn Fn(&mut [i32], usize) -> i32,
    size: usize,
    pack: &'a dyn Fn(&mut [i32]),
) -> impl Fn() + 'a {
    move || {
        let mut arr = vec![0; size];
        pack(&mut arr);

        let mut sorted = arr.clone();
        sorted.sort_unstable();

        if arr.is_empty() {
            return;
        }

        let mut k = vec![1, arr.len()];

        if arr.len() > 2 {
            (0..40).for_each(|_| {
                k.push(random_factory::gen_int_n(2, (arr.len() - 1) as i32) as usize)
            });
        }

        k.into_iter().for_each(|k| {
            let n = f(&mut arr.clone(), k);
            assert_eq!(n, sorted[arr.len() - k]);
        });
    }
}

pub(crate) fn derive_decreasing(
    select: &dyn Fn(&mut [i32], usize) -> i32,
    size: usize,
) -> impl Fn() + '_ {
    derive(select, size, &sequence_builder::pack_decreasing)
}

pub(crate) fn derive_empty(select: &dyn Fn(&mut [i32], usize) -> i32) -> impl Fn() + '_ {
    derive(select, 0, &sequence_builder::pack_identical)
}

pub(crate) fn derive_identical(
    select: &dyn Fn(&mut [i32], usize) -> i32,
    size: usize,
) -> impl Fn() + '_ {
    derive(select, size, &sequence_builder::pack_identical)
}

pub(crate) fn derive_increasing(
    select: &dyn Fn(&mut [i32], usize) -> i32,
    size: usize,
) -> impl Fn() + '_ {
    derive(select, size, &sequence_builder::pack_increasing)
}

pub(crate) fn derive_random(
    select: &dyn Fn(&mut [i32], usize) -> i32,
    size: usize,
) -> impl Fn() + '_ {
    derive(select, size, &sequence_builder::pack_random)
}
