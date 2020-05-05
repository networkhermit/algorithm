<?php
    require_once "fun/vac/algorithms/number_theory/Coprimality.php";
    require_once "fun/vac/util/TestRunner.php";

    use fun\vac\algorithms\number_theory\Coprimality;
    use fun\vac\util\TestRunner;

    function testCoprimality(): bool {
        $sample = [
            [          0,           1, 1],
            [          1,           0, 1],
            [          1,           1, 1],
            [         -1,          -1, 1],
            [         15,         -45, 0],
            [         -7,          51, 1],
            [         -3,          75, 0],
            [        -84,         -32, 0],
            [        -21,         -71, 1],
            [         89,         -11, 1],
            [        -40,          -9, 1],
            [         92,          26, 0],
            [         14,         -95, 1],
            [          5,         -99, 1],
            [        -45,         -57, 0],
            [         59,         -97, 1],
            [        -16,          37, 1],
            [         97,          49, 1],
            [       5998,       -3871, 1],
            [      -2152,       -2909, 1],
            [       8823,         629, 0],
            [       7589,       -7035, 1],
            [      -8669,        6398, 1],
            [       1059,        3522, 0],
            [       9892,       -2990, 0],
            [      -1819,        4740, 1],
            [       7900,       -7244, 0],
            [       8072,       -1226, 0],
            [      -6875,        5777, 1],
            [       -828,        8859, 0],
            [      -3929,        1143, 1],
            [      -5397,        1880, 1],
            [      -9397,       -4343, 1],
            [       2389,       -4026, 1],
            [      -6838,       -7802, 0],
            [      -6118,        7944, 0],
            [       1408,       -8483, 1],
            [       -507,        3510, 0],
            [       5030,       -4818, 0],
            [       2563,       -1867, 1],
            [       3317,       -4585, 1],
            [      -2448,       -2215, 1],
            [      -1982,       -3811, 1],
            [        645,        7907, 1],
            [       5812,        1657, 1],
            [       3944,       -5997, 1],
            [      -2410,       -6522, 0],
            [       4565,        9055, 0],
            [       8478,        1099, 0],
            [       6444,       -7298, 0],
            [       7606,       -7006, 0],
            [       2491,       -2017, 1],
            [       7151,       -9724, 1],
            [       2958,       -5697, 0],
            [       5134,       -7701, 0],
            [      -3544,        9042, 0],
            [       8826,       -4548, 0],
            [       6794,       -5322, 0],
            [      -6906,        7502, 0],
            [        172,        8480, 0],
            [       8244,        4622, 0],
            [      -7315,        5253, 1],
            [       7011,        4285, 1],
            [      -3132,       -4999, 1],
            [      -7332,        -655, 1],
            [       9661,       -1750, 1],
            [       7560,       -3128, 0],
            [       1334,        6234, 0],
            [       8075,        6450, 0],
            [       3283,        8980, 1],
            [       3364,        8482, 0],
            [       4909,        8302, 1],
            [       9332,       -2523, 1],
            [      -8515,        5209, 1],
            [      -1624,        7640, 0],
            [      -6463,        2562, 1],
            [       7912,        5868, 0],
            [      -4825,       -2173, 1],
            [      28876,       30106, 0],
            [    -929840,      887043, 1],
            [     684923,     -588038, 1],
            [    -291411,     -299801, 1],
            [    -905447,     -402122, 1],
            [     830872,     -223425, 1],
            [     831033,     -753398, 1],
            [    -575558,      711716, 0],
            [    -312296,      515492, 0],
            [     595308,      963205, 0],
            [    -770718,     -388434, 0],
            [     353889,      330806, 1],
            [    -174566,      613742, 0],
            [    -884075,       26687, 1],
            [     390743,      204874, 1],
            [     930615,       86524, 1],
            [  -84945524,    15427487, 1],
            [   76038602,   -89688904, 0],
            [   50294214,   -65802481, 1],
            [   83436075,    44248708, 1],
            [   54031677,    92370464, 1],
            [   -6019575,    -1302029, 1],
            [   86952873,   -75470243, 1],
            [  -53215372,    28947490, 0],
            [  -37109442,    75623090, 0],
            [  -11656183,    16147085, 1],
            [   86439843,   -12134914, 1],
            [  -52977427,    48349835, 1],
            [   16234161,    26103566, 1],
            [   73756826,    -6586797, 1],
            [    6836355,    11529043, 1],
            [   29864126,   -80782077, 1],
            [ 1646787325,  1961513442, 1],
            [-1755035190, -1801169490, 0],
            [-1509018194,  1829751775, 1],
            [ -720160017, -1425309680, 1],
            [ 1216287038,  1821933798, 0],
            [-1925479607, -1842455762, 1],
            [ -795996486, -1859155567, 1],
            [ -367280505,   321267794, 1],
            [  829304526, -1575808585, 1],
            [ 1457917042,  1083382210, 0],
            [ -377130535,  1526538188, 1],
            [-1109347700, -1819333109, 1],
            [-1740794578,   278770346, 0],
            [  448156480, -1008775746, 0],
            [  442691160,  1680572092, 0],
            [ 1241208470,  -647438045, 0],
            [ 2147483647,  -561158902, 1],
            [  761395308, -2147483647, 1],
        ];

        for ($i = 0, $size = count($sample); $i < $size; $i++) {
            if ($sample[$i][2] == 0) {
                if (Coprimality\reduceToBinaryGCD($sample[$i][0], $sample[$i][1])) {
                    return false;
                }
            } else {
                if (!Coprimality\reduceToBinaryGCD($sample[$i][0], $sample[$i][1])) {
                    return false;
                }
            }
        }

        for ($i = 0, $size = count($sample); $i < $size; $i++) {
            if ($sample[$i][2] == 0) {
                if (Coprimality\reduceToEuclidean($sample[$i][0], $sample[$i][1])) {
                    return false;
                }
            } else {
                if (!Coprimality\reduceToEuclidean($sample[$i][0], $sample[$i][1])) {
                    return false;
                }
            }
        }

        return true;
    }

    if (count(debug_backtrace()) == 0) {
        TestRunner\parseTest(testCoprimality());
    }
?>
