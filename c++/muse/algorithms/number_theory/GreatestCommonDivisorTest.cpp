#include <muse/algorithms/number_theory/GreatestCommonDivisor.hpp>
#include <muse/util/TestRunner.hpp>

using namespace std;

bool testGreatestCommonDivisor() {
    long sample[][3] = {
        // clang-format off
        {             0,              1,   1},
        {             1,              0,   1},
        {             1,              1,   1},
        {            -1,             -1,   1},
        {           -25,             72,   1},
        {           -16,            -48,  16},
        {           -56,            -54,   2},
        {            29,            -60,   1},
        {            40,            -58,   2},
        {           -61,             97,   1},
        {           -81,            -99,   9},
        {           -46,             59,   1},
        {            59,            -55,   1},
        {           -36,             14,   2},
        {           -42,             19,   1},
        {           -22,             44,  22},
        {           -14,             84,  14},
        {           -99,              6,   3},
        {         -7864,           2672,   8},
        {          8228,           2179,   1},
        {         -1622,          -4471,   1},
        {          8293,          -4881,   1},
        {          9936,           3701,   1},
        {          3828,          -9134,   2},
        {         -9055,            934,   1},
        {         -9427,           2762,   1},
        {         -9960,           6904,   8},
        {         -2912,           1062,   2},
        {          7678,           8016,   2},
        {         -3919,           5143,   1},
        {         -5080,          -9523,   1},
        {          8388,           -952,   4},
        {          8996,           2305,   1},
        {         -6576,           3950,   2},
        {           887,          -6760,   1},
        {          8148,          -8500,   4},
        {          5170,          -3469,   1},
        {         -7312,           2220,   4},
        {         -9841,            742,   1},
        {          3525,           3121,   1},
        {         -8986,           4466,   2},
        {          5192,          -2843,   1},
        {          2196,           4569,   3},
        {          8895,            644,   1},
        {          7633,          -1624,   1},
        {          9405,          -2678,   1},
        {         -4444,          -4640,   4},
        {          2729,           6047,   1},
        {         -5458,           5100,   2},
        {          7727,          -2644,   1},
        {          9825,          -3348,   3},
        {         -7670,           1745,   5},
        {         -9678,          -2111,   1},
        {          4804,          -2588,   4},
        {          4541,           3146,   1},
        {          8691,          -8489,   1},
        {         -9432,          -9286,   2},
        {          7370,           2247,   1},
        {         -6928,           7638,   2},
        {         -2610,           4698, 522},
        {           999,           -349,   1},
        {         -1954,           5542,   2},
        {          7277,           7843,   1},
        {          -815,           3671,   1},
        {         -4961,          -4919,   1},
        {         -3079,          -2410,   1},
        {         -1594,           2267,   1},
        {          5889,          -1103,   1},
        {         -2327,           2864, 179},
        {         -9761,           2090,   1},
        {          5879,            976,   1},
        {         -5212,          -6742,   2},
        {         -1803,           3566,   1},
        {          8078,          -9213,   1},
        {          6391,           8997,   1},
        {         -7333,          -2185,   1},
        {          9552,           6474,   6},
        {          3775,          -7699,   1},
        {       -975165,         693821,   1},
        {       -283946,        -681458,   2},
        {        393833,        -268145,   1},
        {         93577,        -938355, 517},
        {        201821,        -279303,   1},
        {        288085,         -62375,   5},
        {         98342,         460581,   1},
        {        875349,         -88237,   1},
        {       -101284,        -152553,   1},
        {       -711500,        -228037,   1},
        {        413727,         811772,   1},
        {       -286175,         -99810,   5},
        {        959681,         662872,   1},
        {       -541593,        -211585,   1},
        {        935328,        -878838,   6},
        {        544408,        -614262,   2},
        {    56'791'606,     10'091'252,   2},
        {    76'184'637,     55'338'231,   3},
        {    87'827'010,     -2'054'112,   6},
        {   -51'336'665,     47'647'691,   1},
        {    26'786'890,      3'898'842,   2},
        {    70'964'172,     27'242'363,   1},
        {    50'572'364,     24'895'820,   4},
        {   -10'321'594,     90'477'203,   1},
        {   -90'182'979,    -79'448'469,   3},
        {    73'849'322,    -30'438'170,   2},
        {   -30'825'301,    -45'682'165,   1},
        {   -29'366'775,     65'070'021,   3},
        {    19'996'288,    -66'270'963,   1},
        {    -3'388'067,     94'646'426,   1},
        {    28'009'079,     59'009'075,   1},
        {   -35'854'012,    -13'548'532,   4},
        {-1'341'557'375, -1'706'521'640,   5},
        {-1'703'927'494,    617'390'757,   1},
        {-1'926'730'755,  1'188'983'365,   5},
        {   533'277'970,   -542'110'746,   2},
        { 1'508'525'543,    869'090'704,   1},
        { 1'234'248'685,  1'606'411'314,   1},
        {-1'698'647'648, -1'816'407'074,   2},
        { 1'149'345'284,  1'170'749'266,   2},
        {  -711'810'797,   -908'541'229,   1},
        {   424'509'982,    712'101'099,   1},
        { 1'996'000'380,   -457'046'405,   5},
        {   943'706'285,   -419'105'596,   1},
        {  -249'463'342, -2'050'490'722,   2},
        {-1'981'012'181, -1'584'006'152,   1},
        { 1'483'217'656,  1'658'473'101,   1},
        {-1'069'835'847,  1'308'503'268,   3},
        { 2'147'483'647, -1'884'119'046,   1},
        {   645'159'694, -2'147'483'647,   1},
        // clang-format on
    };

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        if (GreatestCommonDivisor::iterativeBinaryGCD(sample[i][0], sample[i][1]) != sample[i][2]) {
            return false;
        }
    }

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        if (GreatestCommonDivisor::recursiveBinaryGCD(sample[i][0], sample[i][1]) != sample[i][2]) {
            return false;
        }
    }

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        if (GreatestCommonDivisor::iterativeEuclidean(sample[i][0], sample[i][1]) != sample[i][2]) {
            return false;
        }
    }

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        if (GreatestCommonDivisor::recursiveEuclidean(sample[i][0], sample[i][1]) != sample[i][2]) {
            return false;
        }
    }

    return true;
}

int main() {
    TestRunner::pick(&testGreatestCommonDivisor);
}
