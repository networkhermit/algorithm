from collections.abc import Callable
from typing import Literal

from muse.algorithms.number_theory import parity, tests
from muse.algorithms.number_theory.tests import UniqueCategorySample as Sample
from muse.util import test_runner

type Category = Literal["Even", "Odd"]
EVEN = "Even"
ODD = "Odd"

sample = [
    Sample(0, EVEN),
    Sample(1, ODD),
    Sample(-1, ODD),
    Sample(60, EVEN),
    Sample(62, EVEN),
    Sample(42, EVEN),
    Sample(-87, ODD),
    Sample(-98, EVEN),
    Sample(-56, EVEN),
    Sample(41, ODD),
    Sample(14, EVEN),
    Sample(20, EVEN),
    Sample(-63, ODD),
    Sample(-81, ODD),
    Sample(-39, ODD),
    Sample(33, ODD),
    Sample(-32, EVEN),
    Sample(-50, EVEN),
    Sample(-1471, ODD),
    Sample(4046, EVEN),
    Sample(5019, ODD),
    Sample(4521, ODD),
    Sample(-8744, EVEN),
    Sample(4787, ODD),
    Sample(4973, ODD),
    Sample(-1184, EVEN),
    Sample(-4437, ODD),
    Sample(-5871, ODD),
    Sample(-4298, EVEN),
    Sample(-2027, ODD),
    Sample(-2796, EVEN),
    Sample(-2209, ODD),
    Sample(-6094, EVEN),
    Sample(3257, ODD),
    Sample(-4732, EVEN),
    Sample(7495, ODD),
    Sample(-3916, EVEN),
    Sample(1469, ODD),
    Sample(6164, EVEN),
    Sample(-7545, ODD),
    Sample(-7763, ODD),
    Sample(7523, ODD),
    Sample(-8076, EVEN),
    Sample(-3778, EVEN),
    Sample(-1648, EVEN),
    Sample(4220, EVEN),
    Sample(-8551, ODD),
    Sample(9692, EVEN),
    Sample(5394, EVEN),
    Sample(2472, EVEN),
    Sample(4056, EVEN),
    Sample(5769, ODD),
    Sample(-2322, EVEN),
    Sample(503, ODD),
    Sample(-8721, ODD),
    Sample(-6344, EVEN),
    Sample(-4335, ODD),
    Sample(1677, ODD),
    Sample(-1703, ODD),
    Sample(-4086, EVEN),
    Sample(7076, EVEN),
    Sample(-7165, ODD),
    Sample(7636, EVEN),
    Sample(-8043, ODD),
    Sample(-3753, ODD),
    Sample(4007, ODD),
    Sample(-261, ODD),
    Sample(-6538, EVEN),
    Sample(9766, EVEN),
    Sample(-7563, ODD),
    Sample(-7944, EVEN),
    Sample(8922, EVEN),
    Sample(-5759, ODD),
    Sample(-8791, ODD),
    Sample(-2211, ODD),
    Sample(3493, ODD),
    Sample(5573, ODD),
    Sample(-2645, ODD),
    Sample(-603656, EVEN),
    Sample(807727, ODD),
    Sample(-69847, ODD),
    Sample(-843676, EVEN),
    Sample(-830961, ODD),
    Sample(-608772, EVEN),
    Sample(931043, ODD),
    Sample(855512, EVEN),
    Sample(358482, EVEN),
    Sample(-98919, ODD),
    Sample(215211, ODD),
    Sample(-933334, EVEN),
    Sample(-613634, EVEN),
    Sample(-95643, ODD),
    Sample(53934, EVEN),
    Sample(161818, EVEN),
    Sample(67_041_621, ODD),
    Sample(99_662_694, EVEN),
    Sample(-94_392_019, ODD),
    Sample(-20_543_495, ODD),
    Sample(-27_591_794, EVEN),
    Sample(-8_314_396, EVEN),
    Sample(97_455_764, EVEN),
    Sample(59_367_920, EVEN),
    Sample(26_856_309, ODD),
    Sample(64_178_815, ODD),
    Sample(-11_480_741, ODD),
    Sample(45_428_276, EVEN),
    Sample(46_193_175, ODD),
    Sample(-31_079_636, EVEN),
    Sample(63_115_980, EVEN),
    Sample(42_559_270, EVEN),
    Sample(-1_645_871_504, EVEN),
    Sample(-1_233_918_598, EVEN),
    Sample(722_012_346, EVEN),
    Sample(-1_525_999_934, EVEN),
    Sample(-365_543_955, ODD),
    Sample(2_008_798_151, ODD),
    Sample(-1_300_713_468, EVEN),
    Sample(1_425_587_979, ODD),
    Sample(1_324_445_673, ODD),
    Sample(2_136_612_365, ODD),
    Sample(-995_371_213, ODD),
    Sample(-2_048_365_905, ODD),
    Sample(2_096_138_065, ODD),
    Sample(-768_738_192, EVEN),
    Sample(-846_034_014, EVEN),
    Sample(411_817_058, EVEN),
    Sample(2_147_483_647, ODD),
    Sample(-2_147_483_648, EVEN),
]


def derive(fn: Callable[[int], bool], c: Category) -> Callable[[], bool]:
    return tests.unique_category_derive(fn, sample, c)


def test_parity() -> bool:
    return (
        derive(parity.modulo_is_even, EVEN)()
        and derive(parity.modulo_is_odd, ODD)()
        and derive(parity.bitwise_is_even, EVEN)()
        and derive(parity.bitwise_is_odd, ODD)()
    )


def main() -> None:
    test_runner.pick(test_parity)


if __name__ == "__main__":
    main()
