pub fn modulo_is_even(n: i64) -> bool {
    n % 2 == 0
}

pub fn modulo_is_odd(n: i64) -> bool {
    n % 2 != 0
}

pub fn bitwise_is_even(n: i64) -> bool {
    (n & 1) == 0
}

pub fn bitwise_is_odd(n: i64) -> bool {
    (n & 1) != 0
}

#[cfg(test)]
pub(crate) mod tests {
    use super::*;

    #[derive(Debug, Eq, PartialEq)]
    pub(crate) enum Category {
        Even,
        Odd,
    }

    pub(crate) fn derive(f: &dyn Fn(i64) -> bool, c: Category) {
        use super::super::tests::{unique_category_derive, UniqueCategorySample as Sample};
        use Category::{Even, Odd};

        let sample = [
            Sample(0, Even),
            Sample(1, Odd),
            Sample(-1, Odd),
            Sample(60, Even),
            Sample(62, Even),
            Sample(42, Even),
            Sample(-87, Odd),
            Sample(-98, Even),
            Sample(-56, Even),
            Sample(41, Odd),
            Sample(14, Even),
            Sample(20, Even),
            Sample(-63, Odd),
            Sample(-81, Odd),
            Sample(-39, Odd),
            Sample(33, Odd),
            Sample(-32, Even),
            Sample(-50, Even),
            Sample(-1471, Odd),
            Sample(4046, Even),
            Sample(5019, Odd),
            Sample(4521, Odd),
            Sample(-8744, Even),
            Sample(4787, Odd),
            Sample(4973, Odd),
            Sample(-1184, Even),
            Sample(-4437, Odd),
            Sample(-5871, Odd),
            Sample(-4298, Even),
            Sample(-2027, Odd),
            Sample(-2796, Even),
            Sample(-2209, Odd),
            Sample(-6094, Even),
            Sample(3257, Odd),
            Sample(-4732, Even),
            Sample(7495, Odd),
            Sample(-3916, Even),
            Sample(1469, Odd),
            Sample(6164, Even),
            Sample(-7545, Odd),
            Sample(-7763, Odd),
            Sample(7523, Odd),
            Sample(-8076, Even),
            Sample(-3778, Even),
            Sample(-1648, Even),
            Sample(4220, Even),
            Sample(-8551, Odd),
            Sample(9692, Even),
            Sample(5394, Even),
            Sample(2472, Even),
            Sample(4056, Even),
            Sample(5769, Odd),
            Sample(-2322, Even),
            Sample(503, Odd),
            Sample(-8721, Odd),
            Sample(-6344, Even),
            Sample(-4335, Odd),
            Sample(1677, Odd),
            Sample(-1703, Odd),
            Sample(-4086, Even),
            Sample(7076, Even),
            Sample(-7165, Odd),
            Sample(7636, Even),
            Sample(-8043, Odd),
            Sample(-3753, Odd),
            Sample(4007, Odd),
            Sample(-261, Odd),
            Sample(-6538, Even),
            Sample(9766, Even),
            Sample(-7563, Odd),
            Sample(-7944, Even),
            Sample(8922, Even),
            Sample(-5759, Odd),
            Sample(-8791, Odd),
            Sample(-2211, Odd),
            Sample(3493, Odd),
            Sample(5573, Odd),
            Sample(-2645, Odd),
            Sample(-603656, Even),
            Sample(807727, Odd),
            Sample(-69847, Odd),
            Sample(-843676, Even),
            Sample(-830961, Odd),
            Sample(-608772, Even),
            Sample(931043, Odd),
            Sample(855512, Even),
            Sample(358482, Even),
            Sample(-98919, Odd),
            Sample(215211, Odd),
            Sample(-933334, Even),
            Sample(-613634, Even),
            Sample(-95643, Odd),
            Sample(53934, Even),
            Sample(161818, Even),
            Sample(67_041_621, Odd),
            Sample(99_662_694, Even),
            Sample(-94_392_019, Odd),
            Sample(-20_543_495, Odd),
            Sample(-27_591_794, Even),
            Sample(-8_314_396, Even),
            Sample(97_455_764, Even),
            Sample(59_367_920, Even),
            Sample(26_856_309, Odd),
            Sample(64_178_815, Odd),
            Sample(-11_480_741, Odd),
            Sample(45_428_276, Even),
            Sample(46_193_175, Odd),
            Sample(-31_079_636, Even),
            Sample(63_115_980, Even),
            Sample(42_559_270, Even),
            Sample(-1_645_871_504, Even),
            Sample(-1_233_918_598, Even),
            Sample(722_012_346, Even),
            Sample(-1_525_999_934, Even),
            Sample(-365_543_955, Odd),
            Sample(2_008_798_151, Odd),
            Sample(-1_300_713_468, Even),
            Sample(1_425_587_979, Odd),
            Sample(1_324_445_673, Odd),
            Sample(2_136_612_365, Odd),
            Sample(-995_371_213, Odd),
            Sample(-2_048_365_905, Odd),
            Sample(2_096_138_065, Odd),
            Sample(-768_738_192, Even),
            Sample(-846_034_014, Even),
            Sample(411_817_058, Even),
            Sample(2_147_483_647, Odd),
            Sample(-2_147_483_648, Even),
        ];

        unique_category_derive(f, &sample, c);
    }

    #[test]
    fn test_parity() {
        use Category::{Even, Odd};

        derive(&modulo_is_even, Even);

        derive(&modulo_is_odd, Odd);

        derive(&bitwise_is_even, Even);

        derive(&bitwise_is_odd, Odd);
    }
}
