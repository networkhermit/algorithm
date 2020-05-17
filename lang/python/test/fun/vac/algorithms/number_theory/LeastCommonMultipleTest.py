from fun.vac.algorithms.number_theory import LeastCommonMultiple
from fun.vac.util import TestRunner

def testLeastCommonMultiple() -> bool:
    sample = [
        # fmt: off
        [     1,      1,             1],
        [    -1,     -1,             1],
        [   -85,     -8,           680],
        [    33,     -2,            66],
        [     8,     32,            32],
        [     8,     89,           712],
        [    16,     78,           624],
        [   -66,     -2,            66],
        [   -50,    -62,          1550],
        [    12,    -25,           300],
        [   -74,    -25,          1850],
        [     4,     24,            24],
        [    90,     29,          2610],
        [   -62,    -85,          5270],
        [    -9,    -50,           450],
        [    90,    -59,          5310],
        [ -8732,  -1743,    15_219_876],
        [ -8329,   8430,    70_213_470],
        [  3300,  -1326,        729300],
        [ -5969,   -523,     3_121_787],
        [ -7044,  -8745,    20_533_260],
        [ -4683,   2491,    11_665_353],
        [  9729,   4329,     4_679_649],
        [  -871,   3189,     2_777_619],
        [ -1158,  -3122,     1_807_638],
        [  9912,  -9910,    49_113_960],
        [  4924,   5842,    14_383_004],
        [  3980,  -6455,     5_138_180],
        [  5420,  -1507,     8_167_940],
        [ -1090,   7747,     8_444_230],
        [ -8008,   7290,    29_189_160],
        [ -2260,  -6189,    13_987_140],
        [  -962,   9376,     4_509_856],
        [  -351,   8756,     3_073_356],
        [  -171,   8401,     1_436_571],
        [ -3110,  -7937,    24_684_070],
        [  6362,   1928,     6_132_968],
        [ -8230,    964,     3_966_860],
        [ -5791,   6186,    35_823_126],
        [ -4204,   9556,    10_043_356],
        [ -3338,   6848,    11_429_312],
        [  -760,   8766,     3_331_080],
        [ -1958,   4928,        438592],
        [ -3830,   -107,        409810],
        [ -7809,   9720,    25_301_160],
        [  7665,    209,     1_601_985],
        [ -6060,   4881,     9_859_620],
        [  2346,  -9979,     1_377_102],
        [  7125,  -5604,    13_309_500],
        [  9862,   3015,    29_733_930],
        [ -1148,   8092,        331772],
        [  8627,   3929,    33_895_483],
        [ -5320,   8927,    47_491_640],
        [ -2301,  -8803,    20_255_703],
        [  6395,  -8793,    56_231_235],
        [  6278,  -2847,        244842],
        [  1623,   3406,     5_527_938],
        [ -3974,   1259,     5_003_266],
        [ -4014,  -3066,     2_051_154],
        [  7546,  -3833,    28_923_818],
        [  4058,  -4338,     8_801_802],
        [  5066,   7450,        126650],
        [ -9458,  -5234,    24_751_586],
        [  4142,   5319,    22_031_298],
        [  4119,  -9963,    13_679_199],
        [   813,   1892,     1_538_196],
        [  4375,   9055,     7_923_125],
        [   388,  -4329,     1_679_652],
        [  3927,   5720,     2_042_040],
        [  5099,  -5563,    28_365_737],
        [ -4174,   4772,     9_959_164],
        [ -8829,   6842,    60_408_018],
        [  9834,   8283,     2_468_334],
        [ -6028,  -2949,    17_776_572],
        [  2627,  -2463,     6_470_301],
        [  9431,  -5198,    49_022_338],
        [  4553,   6198,    28_219_494],
        [  1368,  -8772,     1_000_008],
        [-43336, -36185, 1_568_113_160],
        [ 33606, -19136,   321_542_208],
        [-23381,  31250,   730_656_250],
        [-22377, -22397,   501_177_669],
        [-22766, -10042,   114_308_086],
        [-37379, -13431,   502_037_349],
        [-23248, -18591,   432_203_568],
        [-34910,  37688,   657_844_040],
        [-29156,  26259,   765_607_404],
        [ 33346,  19239,   641_543_694],
        [-31634,  13587,   429_811_158],
        [ 22300,  14536,    81_038_200],
        [-25315,  26148,   661_936_620],
        [-21579, -43116,   310_133_388],
        [-25141,  37870,   952_089_670],
        [ 23741,  10471,   248_592_011],
        [ 26852, -41498,   557_152_148],
        [ 26627, -21629,   575_915_383],
        [-13538, -16865,   228_318_370],
        [ 25331,  31258,   791_796_398],
        [ 33885,  13483,   456_871_455],
        [-43671, -16383,   238_487_331],
        [-29483,  26616,   784_719_528],
        [ 15792, -11008,    10_864_896],
        [ 23749,  40466,   961_027_034],
        [-37432, -42397, 1_587_004_504],
        [-46035, -24237,   123_972_255],
        [-38120,  10830,    41_283_960],
        [-11808, -10396,    30_688_992],
        [-15667, -15508,   242_963_836],
        [-15336, -35057,   537_634_152],
        [ 31050, -13757,   427_154_850],
        [ 42602, -20701,   881_904_002],
        [-20505,  18475,    75_765_975],
        [ 35445,  12527,   444_019_515],
        [-20600, -23702,   244_130_600],
        [ 21432,  15715,   336_803_880],
        [-24635,  45838,    86_863_010],
        [ 20616,  41256,    35_438_904],
        [-40536,  29009, 1_175_908_824],
        [-29015, -40786, 1_183_405_790],
        [ 17092, -30650,   261_934_900],
        [-44577,  31261, 1_393_521_597],
        [-35517,  13916,   494_254_572],
        [-22680,  39616,   112_311_360],
        [ 33207, -33473, 1_111_537_911],
        [ 30250, -28845,   174_512_250],
        [ 45203, -23693, 1_070_994_679],
        [ 46340,  46341, 2_147_441_940],
        [-46340, -46341, 2_147_441_940],
        # fmt: on
    ]

    for i, _ in enumerate(sample):
        if LeastCommonMultiple.reduceToBinaryGCD(sample[i][0], sample[i][1]) != sample[i][2]:
            return False

    for i, _ in enumerate(sample):
        if LeastCommonMultiple.reduceToEuclidean(sample[i][0], sample[i][1]) != sample[i][2]:
            return False

    return True

def main() -> None:
    TestRunner.parseTest(testLeastCommonMultiple())

if __name__ == "__main__":
    main()
