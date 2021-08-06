<?php
    require_once "muse/algorithms/number_theory/GreatestCommonDivisor.php";
    require_once "muse/util/TestRunner.php";

    use muse\algorithms\number_theory\GreatestCommonDivisor;
    use muse\util\TestRunner;

    function testGreatestCommonDivisor(): bool {
        $sample = [
            [          0,           1,   1],
            [          1,           0,   1],
            [          1,           1,   1],
            [         -1,          -1,   1],
            [        -25,          72,   1],
            [        -16,         -48,  16],
            [        -56,         -54,   2],
            [         29,         -60,   1],
            [         40,         -58,   2],
            [        -61,          97,   1],
            [        -81,         -99,   9],
            [        -46,          59,   1],
            [         59,         -55,   1],
            [        -36,          14,   2],
            [        -42,          19,   1],
            [        -22,          44,  22],
            [        -14,          84,  14],
            [        -99,           6,   3],
            [      -7864,        2672,   8],
            [       8228,        2179,   1],
            [      -1622,       -4471,   1],
            [       8293,       -4881,   1],
            [       9936,        3701,   1],
            [       3828,       -9134,   2],
            [      -9055,         934,   1],
            [      -9427,        2762,   1],
            [      -9960,        6904,   8],
            [      -2912,        1062,   2],
            [       7678,        8016,   2],
            [      -3919,        5143,   1],
            [      -5080,       -9523,   1],
            [       8388,        -952,   4],
            [       8996,        2305,   1],
            [      -6576,        3950,   2],
            [        887,       -6760,   1],
            [       8148,       -8500,   4],
            [       5170,       -3469,   1],
            [      -7312,        2220,   4],
            [      -9841,         742,   1],
            [       3525,        3121,   1],
            [      -8986,        4466,   2],
            [       5192,       -2843,   1],
            [       2196,        4569,   3],
            [       8895,         644,   1],
            [       7633,       -1624,   1],
            [       9405,       -2678,   1],
            [      -4444,       -4640,   4],
            [       2729,        6047,   1],
            [      -5458,        5100,   2],
            [       7727,       -2644,   1],
            [       9825,       -3348,   3],
            [      -7670,        1745,   5],
            [      -9678,       -2111,   1],
            [       4804,       -2588,   4],
            [       4541,        3146,   1],
            [       8691,       -8489,   1],
            [      -9432,       -9286,   2],
            [       7370,        2247,   1],
            [      -6928,        7638,   2],
            [      -2610,        4698, 522],
            [        999,        -349,   1],
            [      -1954,        5542,   2],
            [       7277,        7843,   1],
            [       -815,        3671,   1],
            [      -4961,       -4919,   1],
            [      -3079,       -2410,   1],
            [      -1594,        2267,   1],
            [       5889,       -1103,   1],
            [      -2327,        2864, 179],
            [      -9761,        2090,   1],
            [       5879,         976,   1],
            [      -5212,       -6742,   2],
            [      -1803,        3566,   1],
            [       8078,       -9213,   1],
            [       6391,        8997,   1],
            [      -7333,       -2185,   1],
            [       9552,        6474,   6],
            [       3775,       -7699,   1],
            [    -975165,      693821,   1],
            [    -283946,     -681458,   2],
            [     393833,     -268145,   1],
            [      93577,     -938355, 517],
            [     201821,     -279303,   1],
            [     288085,      -62375,   5],
            [      98342,      460581,   1],
            [     875349,      -88237,   1],
            [    -101284,     -152553,   1],
            [    -711500,     -228037,   1],
            [     413727,      811772,   1],
            [    -286175,      -99810,   5],
            [     959681,      662872,   1],
            [    -541593,     -211585,   1],
            [     935328,     -878838,   6],
            [     544408,     -614262,   2],
            [   56791606,    10091252,   2],
            [   76184637,    55338231,   3],
            [   87827010,    -2054112,   6],
            [  -51336665,    47647691,   1],
            [   26786890,     3898842,   2],
            [   70964172,    27242363,   1],
            [   50572364,    24895820,   4],
            [  -10321594,    90477203,   1],
            [  -90182979,   -79448469,   3],
            [   73849322,   -30438170,   2],
            [  -30825301,   -45682165,   1],
            [  -29366775,    65070021,   3],
            [   19996288,   -66270963,   1],
            [   -3388067,    94646426,   1],
            [   28009079,    59009075,   1],
            [  -35854012,   -13548532,   4],
            [-1341557375, -1706521640,   5],
            [-1703927494,   617390757,   1],
            [-1926730755,  1188983365,   5],
            [  533277970,  -542110746,   2],
            [ 1508525543,   869090704,   1],
            [ 1234248685,  1606411314,   1],
            [-1698647648, -1816407074,   2],
            [ 1149345284,  1170749266,   2],
            [ -711810797,  -908541229,   1],
            [  424509982,   712101099,   1],
            [ 1996000380,  -457046405,   5],
            [  943706285,  -419105596,   1],
            [ -249463342, -2050490722,   2],
            [-1981012181, -1584006152,   1],
            [ 1483217656,  1658473101,   1],
            [-1069835847,  1308503268,   3],
            [ 2147483647, -1884119046,   1],
            [  645159694, -2147483647,   1],
        ];

        for ($i = 0, $size = count($sample); $i < $size; $i++) {
            if (GreatestCommonDivisor\iterativeBinaryGCD($sample[$i][0], $sample[$i][1]) != $sample[$i][2]) {
                return false;
            }
        }

        for ($i = 0, $size = count($sample); $i < $size; $i++) {
            if (GreatestCommonDivisor\recursiveBinaryGCD($sample[$i][0], $sample[$i][1]) != $sample[$i][2]) {
                return false;
            }
        }

        for ($i = 0, $size = count($sample); $i < $size; $i++) {
            if (GreatestCommonDivisor\iterativeEuclidean($sample[$i][0], $sample[$i][1]) != $sample[$i][2]) {
                return false;
            }
        }

        for ($i = 0, $size = count($sample); $i < $size; $i++) {
            if (GreatestCommonDivisor\recursiveEuclidean($sample[$i][0], $sample[$i][1]) != $sample[$i][2]) {
                return false;
            }
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\pick("testGreatestCommonDivisor");
    }
?>
