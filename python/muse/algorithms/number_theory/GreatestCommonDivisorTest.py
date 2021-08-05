from muse.algorithms.number_theory import GreatestCommonDivisor
from muse.util import TestRunner

def testGreatestCommonDivisor() -> bool:
    sample = [
        # fmt: off
        [             0,              1,   1],
        [             1,              0,   1],
        [             1,              1,   1],
        [            -1,             -1,   1],
        [           -25,             72,   1],
        [           -16,            -48,  16],
        [           -56,            -54,   2],
        [            29,            -60,   1],
        [            40,            -58,   2],
        [           -61,             97,   1],
        [           -81,            -99,   9],
        [           -46,             59,   1],
        [            59,            -55,   1],
        [           -36,             14,   2],
        [           -42,             19,   1],
        [           -22,             44,  22],
        [           -14,             84,  14],
        [           -99,              6,   3],
        [         -7864,           2672,   8],
        [          8228,           2179,   1],
        [         -1622,          -4471,   1],
        [          8293,          -4881,   1],
        [          9936,           3701,   1],
        [          3828,          -9134,   2],
        [         -9055,            934,   1],
        [         -9427,           2762,   1],
        [         -9960,           6904,   8],
        [         -2912,           1062,   2],
        [          7678,           8016,   2],
        [         -3919,           5143,   1],
        [         -5080,          -9523,   1],
        [          8388,           -952,   4],
        [          8996,           2305,   1],
        [         -6576,           3950,   2],
        [           887,          -6760,   1],
        [          8148,          -8500,   4],
        [          5170,          -3469,   1],
        [         -7312,           2220,   4],
        [         -9841,            742,   1],
        [          3525,           3121,   1],
        [         -8986,           4466,   2],
        [          5192,          -2843,   1],
        [          2196,           4569,   3],
        [          8895,            644,   1],
        [          7633,          -1624,   1],
        [          9405,          -2678,   1],
        [         -4444,          -4640,   4],
        [          2729,           6047,   1],
        [         -5458,           5100,   2],
        [          7727,          -2644,   1],
        [          9825,          -3348,   3],
        [         -7670,           1745,   5],
        [         -9678,          -2111,   1],
        [          4804,          -2588,   4],
        [          4541,           3146,   1],
        [          8691,          -8489,   1],
        [         -9432,          -9286,   2],
        [          7370,           2247,   1],
        [         -6928,           7638,   2],
        [         -2610,           4698, 522],
        [           999,           -349,   1],
        [         -1954,           5542,   2],
        [          7277,           7843,   1],
        [          -815,           3671,   1],
        [         -4961,          -4919,   1],
        [         -3079,          -2410,   1],
        [         -1594,           2267,   1],
        [          5889,          -1103,   1],
        [         -2327,           2864, 179],
        [         -9761,           2090,   1],
        [          5879,            976,   1],
        [         -5212,          -6742,   2],
        [         -1803,           3566,   1],
        [          8078,          -9213,   1],
        [          6391,           8997,   1],
        [         -7333,          -2185,   1],
        [          9552,           6474,   6],
        [          3775,          -7699,   1],
        [       -975165,         693821,   1],
        [       -283946,        -681458,   2],
        [        393833,        -268145,   1],
        [         93577,        -938355, 517],
        [        201821,        -279303,   1],
        [        288085,         -62375,   5],
        [         98342,         460581,   1],
        [        875349,         -88237,   1],
        [       -101284,        -152553,   1],
        [       -711500,        -228037,   1],
        [        413727,         811772,   1],
        [       -286175,         -99810,   5],
        [        959681,         662872,   1],
        [       -541593,        -211585,   1],
        [        935328,        -878838,   6],
        [        544408,        -614262,   2],
        [    56_791_606,     10_091_252,   2],
        [    76_184_637,     55_338_231,   3],
        [    87_827_010,     -2_054_112,   6],
        [   -51_336_665,     47_647_691,   1],
        [    26_786_890,      3_898_842,   2],
        [    70_964_172,     27_242_363,   1],
        [    50_572_364,     24_895_820,   4],
        [   -10_321_594,     90_477_203,   1],
        [   -90_182_979,    -79_448_469,   3],
        [    73_849_322,    -30_438_170,   2],
        [   -30_825_301,    -45_682_165,   1],
        [   -29_366_775,     65_070_021,   3],
        [    19_996_288,    -66_270_963,   1],
        [    -3_388_067,     94_646_426,   1],
        [    28_009_079,     59_009_075,   1],
        [   -35_854_012,    -13_548_532,   4],
        [-1_341_557_375, -1_706_521_640,   5],
        [-1_703_927_494,    617_390_757,   1],
        [-1_926_730_755,  1_188_983_365,   5],
        [   533_277_970,   -542_110_746,   2],
        [ 1_508_525_543,    869_090_704,   1],
        [ 1_234_248_685,  1_606_411_314,   1],
        [-1_698_647_648, -1_816_407_074,   2],
        [ 1_149_345_284,  1_170_749_266,   2],
        [  -711_810_797,   -908_541_229,   1],
        [   424_509_982,    712_101_099,   1],
        [ 1_996_000_380,   -457_046_405,   5],
        [   943_706_285,   -419_105_596,   1],
        [  -249_463_342, -2_050_490_722,   2],
        [-1_981_012_181, -1_584_006_152,   1],
        [ 1_483_217_656,  1_658_473_101,   1],
        [-1_069_835_847,  1_308_503_268,   3],
        [ 2_147_483_647, -1_884_119_046,   1],
        [   645_159_694, -2_147_483_647,   1],
        # fmt: on
    ]

    for i, _ in enumerate(sample):
        if GreatestCommonDivisor.iterativeBinaryGCD(sample[i][0], sample[i][1]) != sample[i][2]:
            return False

    for i, _ in enumerate(sample):
        if GreatestCommonDivisor.recursiveBinaryGCD(sample[i][0], sample[i][1]) != sample[i][2]:
            return False

    for i, _ in enumerate(sample):
        if GreatestCommonDivisor.iterativeEuclidean(sample[i][0], sample[i][1]) != sample[i][2]:
            return False

    for i, _ in enumerate(sample):
        if GreatestCommonDivisor.recursiveEuclidean(sample[i][0], sample[i][1]) != sample[i][2]:
            return False

    return True

def main() -> None:
    TestRunner.pick(testGreatestCommonDivisor)

if __name__ == "__main__":
    main()
