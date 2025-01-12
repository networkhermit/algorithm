from collections.abc import Callable

from muse.algorithms.number_theory import parity, tests
from muse.util import test_runner

EVEN = "Even"
ODD = "Odd"

sample = [
    {"n": 0, "category": EVEN},
    {"n": 1, "category": ODD},
    {"n": -1, "category": ODD},
    {"n": 60, "category": EVEN},
    {"n": 62, "category": EVEN},
    {"n": 42, "category": EVEN},
    {"n": -87, "category": ODD},
    {"n": -98, "category": EVEN},
    {"n": -56, "category": EVEN},
    {"n": 41, "category": ODD},
    {"n": 14, "category": EVEN},
    {"n": 20, "category": EVEN},
    {"n": -63, "category": ODD},
    {"n": -81, "category": ODD},
    {"n": -39, "category": ODD},
    {"n": 33, "category": ODD},
    {"n": -32, "category": EVEN},
    {"n": -50, "category": EVEN},
    {"n": -1471, "category": ODD},
    {"n": 4046, "category": EVEN},
    {"n": 5019, "category": ODD},
    {"n": 4521, "category": ODD},
    {"n": -8744, "category": EVEN},
    {"n": 4787, "category": ODD},
    {"n": 4973, "category": ODD},
    {"n": -1184, "category": EVEN},
    {"n": -4437, "category": ODD},
    {"n": -5871, "category": ODD},
    {"n": -4298, "category": EVEN},
    {"n": -2027, "category": ODD},
    {"n": -2796, "category": EVEN},
    {"n": -2209, "category": ODD},
    {"n": -6094, "category": EVEN},
    {"n": 3257, "category": ODD},
    {"n": -4732, "category": EVEN},
    {"n": 7495, "category": ODD},
    {"n": -3916, "category": EVEN},
    {"n": 1469, "category": ODD},
    {"n": 6164, "category": EVEN},
    {"n": -7545, "category": ODD},
    {"n": -7763, "category": ODD},
    {"n": 7523, "category": ODD},
    {"n": -8076, "category": EVEN},
    {"n": -3778, "category": EVEN},
    {"n": -1648, "category": EVEN},
    {"n": 4220, "category": EVEN},
    {"n": -8551, "category": ODD},
    {"n": 9692, "category": EVEN},
    {"n": 5394, "category": EVEN},
    {"n": 2472, "category": EVEN},
    {"n": 4056, "category": EVEN},
    {"n": 5769, "category": ODD},
    {"n": -2322, "category": EVEN},
    {"n": 503, "category": ODD},
    {"n": -8721, "category": ODD},
    {"n": -6344, "category": EVEN},
    {"n": -4335, "category": ODD},
    {"n": 1677, "category": ODD},
    {"n": -1703, "category": ODD},
    {"n": -4086, "category": EVEN},
    {"n": 7076, "category": EVEN},
    {"n": -7165, "category": ODD},
    {"n": 7636, "category": EVEN},
    {"n": -8043, "category": ODD},
    {"n": -3753, "category": ODD},
    {"n": 4007, "category": ODD},
    {"n": -261, "category": ODD},
    {"n": -6538, "category": EVEN},
    {"n": 9766, "category": EVEN},
    {"n": -7563, "category": ODD},
    {"n": -7944, "category": EVEN},
    {"n": 8922, "category": EVEN},
    {"n": -5759, "category": ODD},
    {"n": -8791, "category": ODD},
    {"n": -2211, "category": ODD},
    {"n": 3493, "category": ODD},
    {"n": 5573, "category": ODD},
    {"n": -2645, "category": ODD},
    {"n": -603656, "category": EVEN},
    {"n": 807727, "category": ODD},
    {"n": -69847, "category": ODD},
    {"n": -843676, "category": EVEN},
    {"n": -830961, "category": ODD},
    {"n": -608772, "category": EVEN},
    {"n": 931043, "category": ODD},
    {"n": 855512, "category": EVEN},
    {"n": 358482, "category": EVEN},
    {"n": -98919, "category": ODD},
    {"n": 215211, "category": ODD},
    {"n": -933334, "category": EVEN},
    {"n": -613634, "category": EVEN},
    {"n": -95643, "category": ODD},
    {"n": 53934, "category": EVEN},
    {"n": 161818, "category": EVEN},
    {"n": 67_041_621, "category": ODD},
    {"n": 99_662_694, "category": EVEN},
    {"n": -94_392_019, "category": ODD},
    {"n": -20_543_495, "category": ODD},
    {"n": -27_591_794, "category": EVEN},
    {"n": -8_314_396, "category": EVEN},
    {"n": 97_455_764, "category": EVEN},
    {"n": 59_367_920, "category": EVEN},
    {"n": 26_856_309, "category": ODD},
    {"n": 64_178_815, "category": ODD},
    {"n": -11_480_741, "category": ODD},
    {"n": 45_428_276, "category": EVEN},
    {"n": 46_193_175, "category": ODD},
    {"n": -31_079_636, "category": EVEN},
    {"n": 63_115_980, "category": EVEN},
    {"n": 42_559_270, "category": EVEN},
    {"n": -1_645_871_504, "category": EVEN},
    {"n": -1_233_918_598, "category": EVEN},
    {"n": 722_012_346, "category": EVEN},
    {"n": -1_525_999_934, "category": EVEN},
    {"n": -365_543_955, "category": ODD},
    {"n": 2_008_798_151, "category": ODD},
    {"n": -1_300_713_468, "category": EVEN},
    {"n": 1_425_587_979, "category": ODD},
    {"n": 1_324_445_673, "category": ODD},
    {"n": 2_136_612_365, "category": ODD},
    {"n": -995_371_213, "category": ODD},
    {"n": -2_048_365_905, "category": ODD},
    {"n": 2_096_138_065, "category": ODD},
    {"n": -768_738_192, "category": EVEN},
    {"n": -846_034_014, "category": EVEN},
    {"n": 411_817_058, "category": EVEN},
    {"n": 2_147_483_647, "category": ODD},
    {"n": -2_147_483_648, "category": EVEN},
]


def derive(fn: Callable[[int], bool], c: object) -> Callable[[], bool]:
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
