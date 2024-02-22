use super::*;

#[test]
fn test_gen_int_n() {
    (0..8192).for_each(|_| {
        assert_eq!(gen_int_n(0, 0), 0);

        assert_eq!(gen_int_n(1, 1), 1);

        let value = gen_int_n(0, 1);
        assert!((0..=1).contains(&value));

        let value = gen_int_n(100, 10000);
        assert_eq!(gen_int_n(value, value), value);
        assert!((100..=10000).contains(&value));
    });
}

#[test]
fn test_gen_even() {
    (0..8192).for_each(|_| {
        assert!((gen_even() & 1) == 0);
    });
}

#[test]
fn test_gen_odd() {
    (0..8192).for_each(|_| {
        assert!((gen_odd() & 1) != 0);
    });
}
