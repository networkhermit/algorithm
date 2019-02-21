package main

import (
    "fun/vac/algorithms/number_theory/LeastCommonMultiple"
    "fun/vac/util/TestRunner"
)

func testLeastCommonMultiple() bool {
    mapping := [][3]int64 {
        {     1,      1,          1},
        {    -1,     -1,          1},
        {   -85,     -8,        680},
        {    33,     -2,         66},
        {     8,     32,         32},
        {     8,     89,        712},
        {    16,     78,        624},
        {   -66,     -2,         66},
        {   -50,    -62,       1550},
        {    12,    -25,        300},
        {   -74,    -25,       1850},
        {     4,     24,         24},
        {    90,     29,       2610},
        {   -62,    -85,       5270},
        {    -9,    -50,        450},
        {    90,    -59,       5310},
        { -8732,  -1743,   15219876},
        { -8329,   8430,   70213470},
        {  3300,  -1326,     729300},
        { -5969,   -523,    3121787},
        { -7044,  -8745,   20533260},
        { -4683,   2491,   11665353},
        {  9729,   4329,    4679649},
        {  -871,   3189,    2777619},
        { -1158,  -3122,    1807638},
        {  9912,  -9910,   49113960},
        {  4924,   5842,   14383004},
        {  3980,  -6455,    5138180},
        {  5420,  -1507,    8167940},
        { -1090,   7747,    8444230},
        { -8008,   7290,   29189160},
        { -2260,  -6189,   13987140},
        {  -962,   9376,    4509856},
        {  -351,   8756,    3073356},
        {  -171,   8401,    1436571},
        { -3110,  -7937,   24684070},
        {  6362,   1928,    6132968},
        { -8230,    964,    3966860},
        { -5791,   6186,   35823126},
        { -4204,   9556,   10043356},
        { -3338,   6848,   11429312},
        {  -760,   8766,    3331080},
        { -1958,   4928,     438592},
        { -3830,   -107,     409810},
        { -7809,   9720,   25301160},
        {  7665,    209,    1601985},
        { -6060,   4881,    9859620},
        {  2346,  -9979,    1377102},
        {  7125,  -5604,   13309500},
        {  9862,   3015,   29733930},
        { -1148,   8092,     331772},
        {  8627,   3929,   33895483},
        { -5320,   8927,   47491640},
        { -2301,  -8803,   20255703},
        {  6395,  -8793,   56231235},
        {  6278,  -2847,     244842},
        {  1623,   3406,    5527938},
        { -3974,   1259,    5003266},
        { -4014,  -3066,    2051154},
        {  7546,  -3833,   28923818},
        {  4058,  -4338,    8801802},
        {  5066,   7450,     126650},
        { -9458,  -5234,   24751586},
        {  4142,   5319,   22031298},
        {  4119,  -9963,   13679199},
        {   813,   1892,    1538196},
        {  4375,   9055,    7923125},
        {   388,  -4329,    1679652},
        {  3927,   5720,    2042040},
        {  5099,  -5563,   28365737},
        { -4174,   4772,    9959164},
        { -8829,   6842,   60408018},
        {  9834,   8283,    2468334},
        { -6028,  -2949,   17776572},
        {  2627,  -2463,    6470301},
        {  9431,  -5198,   49022338},
        {  4553,   6198,   28219494},
        {  1368,  -8772,    1000008},
        {-43336, -36185, 1568113160},
        { 33606, -19136,  321542208},
        {-23381,  31250,  730656250},
        {-22377, -22397,  501177669},
        {-22766, -10042,  114308086},
        {-37379, -13431,  502037349},
        {-23248, -18591,  432203568},
        {-34910,  37688,  657844040},
        {-29156,  26259,  765607404},
        { 33346,  19239,  641543694},
        {-31634,  13587,  429811158},
        { 22300,  14536,   81038200},
        {-25315,  26148,  661936620},
        {-21579, -43116,  310133388},
        {-25141,  37870,  952089670},
        { 23741,  10471,  248592011},
        { 26852, -41498,  557152148},
        { 26627, -21629,  575915383},
        {-13538, -16865,  228318370},
        { 25331,  31258,  791796398},
        { 33885,  13483,  456871455},
        {-43671, -16383,  238487331},
        {-29483,  26616,  784719528},
        { 15792, -11008,   10864896},
        { 23749,  40466,  961027034},
        {-37432, -42397, 1587004504},
        {-46035, -24237,  123972255},
        {-38120,  10830,   41283960},
        {-11808, -10396,   30688992},
        {-15667, -15508,  242963836},
        {-15336, -35057,  537634152},
        { 31050, -13757,  427154850},
        { 42602, -20701,  881904002},
        {-20505,  18475,   75765975},
        { 35445,  12527,  444019515},
        {-20600, -23702,  244130600},
        { 21432,  15715,  336803880},
        {-24635,  45838,   86863010},
        { 20616,  41256,   35438904},
        {-40536,  29009, 1175908824},
        {-29015, -40786, 1183405790},
        { 17092, -30650,  261934900},
        {-44577,  31261, 1393521597},
        {-35517,  13916,  494254572},
        {-22680,  39616,  112311360},
        { 33207, -33473, 1111537911},
        { 30250, -28845,  174512250},
        { 45203, -23693, 1070994679},
        { 46340,  46341, 2147441940},
        {-46340, -46341, 2147441940},
    }

    for i, _ := range mapping {
        if LeastCommonMultiple.ReduceToBinaryGCD(mapping[i][0], mapping[i][1]) != mapping[i][2] {
            return false
        }
    }

    for i, _ := range mapping {
        if LeastCommonMultiple.ReduceToEuclidean(mapping[i][0], mapping[i][1]) != mapping[i][2] {
            return false
        }
    }

    return true
}

func main() {
    TestRunner.ParseTest(testLeastCommonMultiple())
}
