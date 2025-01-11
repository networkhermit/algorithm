use super::*;

const SAMPLE: [(i64, i64); 13] = [
    (0, 1),
    (1, 1),
    (2, 2),
    (3, 6),
    (4, 24),
    (5, 120),
    (6, 720),
    (7, 5040),
    (8, 40320),
    (9, 362880),
    (10, 3_628_800),
    (11, 39_916_800),
    (12, 479_001_600),
];

pub(crate) fn derive(f: &dyn Fn(i64) -> i64) -> impl Fn() + '_ {
    move || {
        SAMPLE.iter().for_each(|&(n, expected)| {
            let actual = f(n);
            assert_eq!(
                actual, expected,
                "fn({}) returned the left result, expect the right result",
                n
            );
        });
    }
}

#[test]
fn test_factorial() {
    derive(&iterative_procedure)();

    derive(&recursive_procedure)();
}
