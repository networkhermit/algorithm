use crate::util::{sequence_builder, sequences};

pub(crate) fn derive<'a>(
    f: &'a dyn Fn(&mut [i32]),
    size: usize,
    pack: &'a dyn Fn(&mut [i32]),
) -> impl Fn() + 'a {
    move || {
        let mut arr = vec![0; size];
        pack(&mut arr);

        let checksum = sequences::parity_checksum(&arr);

        f(&mut arr);

        assert_eq!(
            sequences::parity_checksum(&arr),
            checksum,
            "elements returned have been altered"
        );

        assert!(
            sequences::is_sorted(&arr),
            "elements returned are not in sorted order"
        );
    }
}

pub(crate) fn derive_decreasing(f: &dyn Fn(&mut [i32]), size: usize) -> impl Fn() + '_ {
    derive(f, size, &sequence_builder::pack_decreasing)
}

pub(crate) fn derive_empty(f: &dyn Fn(&mut [i32])) -> impl Fn() + '_ {
    derive(f, 0, &sequence_builder::pack_identical)
}

pub(crate) fn derive_identical(f: &dyn Fn(&mut [i32]), size: usize) -> impl Fn() + '_ {
    derive(f, size, &sequence_builder::pack_identical)
}

pub(crate) fn derive_increasing(f: &dyn Fn(&mut [i32]), size: usize) -> impl Fn() + '_ {
    derive(f, size, &sequence_builder::pack_increasing)
}

pub(crate) fn derive_random(f: &dyn Fn(&mut [i32]), size: usize) -> impl Fn() + '_ {
    derive(f, size, &sequence_builder::pack_random)
}
