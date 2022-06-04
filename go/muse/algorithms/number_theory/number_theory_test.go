package number_theory

import (
	"testing"

	"muse/algorithms/number_theory/primality"
)

func TestIsCoprime(t *testing.T) {
	sample := [][3]int64{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
		{-1, -1, 1},
		{15, -45, 0},
		{-7, 51, 1},
		{-3, 75, 0},
		{-84, -32, 0},
		{-21, -71, 1},
		{89, -11, 1},
		{-40, -9, 1},
		{92, 26, 0},
		{14, -95, 1},
		{5, -99, 1},
		{-45, -57, 0},
		{59, -97, 1},
		{-16, 37, 1},
		{97, 49, 1},
		{5998, -3871, 1},
		{-2152, -2909, 1},
		{8823, 629, 0},
		{7589, -7035, 1},
		{-8669, 6398, 1},
		{1059, 3522, 0},
		{9892, -2990, 0},
		{-1819, 4740, 1},
		{7900, -7244, 0},
		{8072, -1226, 0},
		{-6875, 5777, 1},
		{-828, 8859, 0},
		{-3929, 1143, 1},
		{-5397, 1880, 1},
		{-9397, -4343, 1},
		{2389, -4026, 1},
		{-6838, -7802, 0},
		{-6118, 7944, 0},
		{1408, -8483, 1},
		{-507, 3510, 0},
		{5030, -4818, 0},
		{2563, -1867, 1},
		{3317, -4585, 1},
		{-2448, -2215, 1},
		{-1982, -3811, 1},
		{645, 7907, 1},
		{5812, 1657, 1},
		{3944, -5997, 1},
		{-2410, -6522, 0},
		{4565, 9055, 0},
		{8478, 1099, 0},
		{6444, -7298, 0},
		{7606, -7006, 0},
		{2491, -2017, 1},
		{7151, -9724, 1},
		{2958, -5697, 0},
		{5134, -7701, 0},
		{-3544, 9042, 0},
		{8826, -4548, 0},
		{6794, -5322, 0},
		{-6906, 7502, 0},
		{172, 8480, 0},
		{8244, 4622, 0},
		{-7315, 5253, 1},
		{7011, 4285, 1},
		{-3132, -4999, 1},
		{-7332, -655, 1},
		{9661, -1750, 1},
		{7560, -3128, 0},
		{1334, 6234, 0},
		{8075, 6450, 0},
		{3283, 8980, 1},
		{3364, 8482, 0},
		{4909, 8302, 1},
		{9332, -2523, 1},
		{-8515, 5209, 1},
		{-1624, 7640, 0},
		{-6463, 2562, 1},
		{7912, 5868, 0},
		{-4825, -2173, 1},
		{28876, 30106, 0},
		{-929840, 887043, 1},
		{684923, -588038, 1},
		{-291411, -299801, 1},
		{-905447, -402122, 1},
		{830872, -223425, 1},
		{831033, -753398, 1},
		{-575558, 711716, 0},
		{-312296, 515492, 0},
		{595308, 963205, 0},
		{-770718, -388434, 0},
		{353889, 330806, 1},
		{-174566, 613742, 0},
		{-884075, 26687, 1},
		{390743, 204874, 1},
		{930615, 86524, 1},
		{-84945524, 15427487, 1},
		{76038602, -89688904, 0},
		{50294214, -65802481, 1},
		{83436075, 44248708, 1},
		{54031677, 92370464, 1},
		{-6019575, -1302029, 1},
		{86952873, -75470243, 1},
		{-53215372, 28947490, 0},
		{-37109442, 75623090, 0},
		{-11656183, 16147085, 1},
		{86439843, -12134914, 1},
		{-52977427, 48349835, 1},
		{16234161, 26103566, 1},
		{73756826, -6586797, 1},
		{6836355, 11529043, 1},
		{29864126, -80782077, 1},
		{1646787325, 1961513442, 1},
		{-1755035190, -1801169490, 0},
		{-1509018194, 1829751775, 1},
		{-720160017, -1425309680, 1},
		{1216287038, 1821933798, 0},
		{-1925479607, -1842455762, 1},
		{-795996486, -1859155567, 1},
		{-367280505, 321267794, 1},
		{829304526, -1575808585, 1},
		{1457917042, 1083382210, 0},
		{-377130535, 1526538188, 1},
		{-1109347700, -1819333109, 1},
		{-1740794578, 278770346, 0},
		{448156480, -1008775746, 0},
		{442691160, 1680572092, 0},
		{1241208470, -647438045, 0},
		{2147483647, -561158902, 1},
		{761395308, -2147483647, 1},
	}

	for i := range sample {
		if sample[i][2] == 0 {
			if IsCoprime(sample[i][0], sample[i][1]) {
				t.FailNow()
			}
		} else {
			if !IsCoprime(sample[i][0], sample[i][1]) {
				t.FailNow()
			}
		}
	}
}

func TestFactorial(t *testing.T) {
	sample := [][2]int64{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
		{11, 39916800},
		{12, 479001600},
	}

	for i := range sample {
		if Factorial(sample[i][0]) != sample[i][1] {
			t.FailNow()
		}
	}
}

func TestFibonacci(t *testing.T) {
	sample := [][2]int64{
		{-31, 1346269},
		{-30, -832040},
		{-29, 514229},
		{-28, -317811},
		{-27, 196418},
		{-26, -121393},
		{-25, 75025},
		{-24, -46368},
		{-23, 28657},
		{-22, -17711},
		{-21, 10946},
		{-20, -6765},
		{-19, 4181},
		{-18, -2584},
		{-17, 1597},
		{-16, -987},
		{-15, 610},
		{-14, -377},
		{-13, 233},
		{-12, -144},
		{-11, 89},
		{-10, -55},
		{-9, 34},
		{-8, -21},
		{-7, 13},
		{-6, -8},
		{-5, 5},
		{-4, -3},
		{-3, 2},
		{-2, -1},
		{-1, 1},
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
		{15, 610},
		{16, 987},
		{17, 1597},
		{18, 2584},
		{19, 4181},
		{20, 6765},
		{21, 10946},
		{22, 17711},
		{23, 28657},
		{24, 46368},
		{25, 75025},
		{26, 121393},
		{27, 196418},
		{28, 317811},
		{29, 514229},
		{30, 832040},
		{31, 1346269},
	}

	for i := range sample {
		if Fibonacci(sample[i][0]) != sample[i][1] {
			t.FailNow()
		}
	}
}

func TestGCD(t *testing.T) {
	sample := [][3]int64{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
		{-1, -1, 1},
		{-25, 72, 1},
		{-16, -48, 16},
		{-56, -54, 2},
		{29, -60, 1},
		{40, -58, 2},
		{-61, 97, 1},
		{-81, -99, 9},
		{-46, 59, 1},
		{59, -55, 1},
		{-36, 14, 2},
		{-42, 19, 1},
		{-22, 44, 22},
		{-14, 84, 14},
		{-99, 6, 3},
		{-7864, 2672, 8},
		{8228, 2179, 1},
		{-1622, -4471, 1},
		{8293, -4881, 1},
		{9936, 3701, 1},
		{3828, -9134, 2},
		{-9055, 934, 1},
		{-9427, 2762, 1},
		{-9960, 6904, 8},
		{-2912, 1062, 2},
		{7678, 8016, 2},
		{-3919, 5143, 1},
		{-5080, -9523, 1},
		{8388, -952, 4},
		{8996, 2305, 1},
		{-6576, 3950, 2},
		{887, -6760, 1},
		{8148, -8500, 4},
		{5170, -3469, 1},
		{-7312, 2220, 4},
		{-9841, 742, 1},
		{3525, 3121, 1},
		{-8986, 4466, 2},
		{5192, -2843, 1},
		{2196, 4569, 3},
		{8895, 644, 1},
		{7633, -1624, 1},
		{9405, -2678, 1},
		{-4444, -4640, 4},
		{2729, 6047, 1},
		{-5458, 5100, 2},
		{7727, -2644, 1},
		{9825, -3348, 3},
		{-7670, 1745, 5},
		{-9678, -2111, 1},
		{4804, -2588, 4},
		{4541, 3146, 1},
		{8691, -8489, 1},
		{-9432, -9286, 2},
		{7370, 2247, 1},
		{-6928, 7638, 2},
		{-2610, 4698, 522},
		{999, -349, 1},
		{-1954, 5542, 2},
		{7277, 7843, 1},
		{-815, 3671, 1},
		{-4961, -4919, 1},
		{-3079, -2410, 1},
		{-1594, 2267, 1},
		{5889, -1103, 1},
		{-2327, 2864, 179},
		{-9761, 2090, 1},
		{5879, 976, 1},
		{-5212, -6742, 2},
		{-1803, 3566, 1},
		{8078, -9213, 1},
		{6391, 8997, 1},
		{-7333, -2185, 1},
		{9552, 6474, 6},
		{3775, -7699, 1},
		{-975165, 693821, 1},
		{-283946, -681458, 2},
		{393833, -268145, 1},
		{93577, -938355, 517},
		{201821, -279303, 1},
		{288085, -62375, 5},
		{98342, 460581, 1},
		{875349, -88237, 1},
		{-101284, -152553, 1},
		{-711500, -228037, 1},
		{413727, 811772, 1},
		{-286175, -99810, 5},
		{959681, 662872, 1},
		{-541593, -211585, 1},
		{935328, -878838, 6},
		{544408, -614262, 2},
		{56791606, 10091252, 2},
		{76184637, 55338231, 3},
		{87827010, -2054112, 6},
		{-51336665, 47647691, 1},
		{26786890, 3898842, 2},
		{70964172, 27242363, 1},
		{50572364, 24895820, 4},
		{-10321594, 90477203, 1},
		{-90182979, -79448469, 3},
		{73849322, -30438170, 2},
		{-30825301, -45682165, 1},
		{-29366775, 65070021, 3},
		{19996288, -66270963, 1},
		{-3388067, 94646426, 1},
		{28009079, 59009075, 1},
		{-35854012, -13548532, 4},
		{-1341557375, -1706521640, 5},
		{-1703927494, 617390757, 1},
		{-1926730755, 1188983365, 5},
		{533277970, -542110746, 2},
		{1508525543, 869090704, 1},
		{1234248685, 1606411314, 1},
		{-1698647648, -1816407074, 2},
		{1149345284, 1170749266, 2},
		{-711810797, -908541229, 1},
		{424509982, 712101099, 1},
		{1996000380, -457046405, 5},
		{943706285, -419105596, 1},
		{-249463342, -2050490722, 2},
		{-1981012181, -1584006152, 1},
		{1483217656, 1658473101, 1},
		{-1069835847, 1308503268, 3},
		{2147483647, -1884119046, 1},
		{645159694, -2147483647, 1},
	}

	for i := range sample {
		if GCD(sample[i][0], sample[i][1]) != sample[i][2] {
			t.FailNow()
		}
	}
}

func TestLCM(t *testing.T) {
	sample := [][3]int64{
		{1, 1, 1},
		{-1, -1, 1},
		{-85, -8, 680},
		{33, -2, 66},
		{8, 32, 32},
		{8, 89, 712},
		{16, 78, 624},
		{-66, -2, 66},
		{-50, -62, 1550},
		{12, -25, 300},
		{-74, -25, 1850},
		{4, 24, 24},
		{90, 29, 2610},
		{-62, -85, 5270},
		{-9, -50, 450},
		{90, -59, 5310},
		{-8732, -1743, 15219876},
		{-8329, 8430, 70213470},
		{3300, -1326, 729300},
		{-5969, -523, 3121787},
		{-7044, -8745, 20533260},
		{-4683, 2491, 11665353},
		{9729, 4329, 4679649},
		{-871, 3189, 2777619},
		{-1158, -3122, 1807638},
		{9912, -9910, 49113960},
		{4924, 5842, 14383004},
		{3980, -6455, 5138180},
		{5420, -1507, 8167940},
		{-1090, 7747, 8444230},
		{-8008, 7290, 29189160},
		{-2260, -6189, 13987140},
		{-962, 9376, 4509856},
		{-351, 8756, 3073356},
		{-171, 8401, 1436571},
		{-3110, -7937, 24684070},
		{6362, 1928, 6132968},
		{-8230, 964, 3966860},
		{-5791, 6186, 35823126},
		{-4204, 9556, 10043356},
		{-3338, 6848, 11429312},
		{-760, 8766, 3331080},
		{-1958, 4928, 438592},
		{-3830, -107, 409810},
		{-7809, 9720, 25301160},
		{7665, 209, 1601985},
		{-6060, 4881, 9859620},
		{2346, -9979, 1377102},
		{7125, -5604, 13309500},
		{9862, 3015, 29733930},
		{-1148, 8092, 331772},
		{8627, 3929, 33895483},
		{-5320, 8927, 47491640},
		{-2301, -8803, 20255703},
		{6395, -8793, 56231235},
		{6278, -2847, 244842},
		{1623, 3406, 5527938},
		{-3974, 1259, 5003266},
		{-4014, -3066, 2051154},
		{7546, -3833, 28923818},
		{4058, -4338, 8801802},
		{5066, 7450, 126650},
		{-9458, -5234, 24751586},
		{4142, 5319, 22031298},
		{4119, -9963, 13679199},
		{813, 1892, 1538196},
		{4375, 9055, 7923125},
		{388, -4329, 1679652},
		{3927, 5720, 2042040},
		{5099, -5563, 28365737},
		{-4174, 4772, 9959164},
		{-8829, 6842, 60408018},
		{9834, 8283, 2468334},
		{-6028, -2949, 17776572},
		{2627, -2463, 6470301},
		{9431, -5198, 49022338},
		{4553, 6198, 28219494},
		{1368, -8772, 1000008},
		{-43336, -36185, 1568113160},
		{33606, -19136, 321542208},
		{-23381, 31250, 730656250},
		{-22377, -22397, 501177669},
		{-22766, -10042, 114308086},
		{-37379, -13431, 502037349},
		{-23248, -18591, 432203568},
		{-34910, 37688, 657844040},
		{-29156, 26259, 765607404},
		{33346, 19239, 641543694},
		{-31634, 13587, 429811158},
		{22300, 14536, 81038200},
		{-25315, 26148, 661936620},
		{-21579, -43116, 310133388},
		{-25141, 37870, 952089670},
		{23741, 10471, 248592011},
		{26852, -41498, 557152148},
		{26627, -21629, 575915383},
		{-13538, -16865, 228318370},
		{25331, 31258, 791796398},
		{33885, 13483, 456871455},
		{-43671, -16383, 238487331},
		{-29483, 26616, 784719528},
		{15792, -11008, 10864896},
		{23749, 40466, 961027034},
		{-37432, -42397, 1587004504},
		{-46035, -24237, 123972255},
		{-38120, 10830, 41283960},
		{-11808, -10396, 30688992},
		{-15667, -15508, 242963836},
		{-15336, -35057, 537634152},
		{31050, -13757, 427154850},
		{42602, -20701, 881904002},
		{-20505, 18475, 75765975},
		{35445, 12527, 444019515},
		{-20600, -23702, 244130600},
		{21432, 15715, 336803880},
		{-24635, 45838, 86863010},
		{20616, 41256, 35438904},
		{-40536, 29009, 1175908824},
		{-29015, -40786, 1183405790},
		{17092, -30650, 261934900},
		{-44577, 31261, 1393521597},
		{-35517, 13916, 494254572},
		{-22680, 39616, 112311360},
		{33207, -33473, 1111537911},
		{30250, -28845, 174512250},
		{45203, -23693, 1070994679},
		{46340, 46341, 2147441940},
		{-46340, -46341, 2147441940},
	}

	for i := range sample {
		if LCM(sample[i][0], sample[i][1]) != sample[i][2] {
			t.FailNow()
		}
	}
}

func TestIsEven(t *testing.T) {
	sample := [][2]int64{
		{0, 0},
		{1, 1},
		{-1, 1},
		{60, 0},
		{62, 0},
		{42, 0},
		{-87, 1},
		{-98, 0},
		{-56, 0},
		{41, 1},
		{14, 0},
		{20, 0},
		{-63, 1},
		{-81, 1},
		{-39, 1},
		{33, 1},
		{-32, 0},
		{-50, 0},
		{-1471, 1},
		{4046, 0},
		{5019, 1},
		{4521, 1},
		{-8744, 0},
		{4787, 1},
		{4973, 1},
		{-1184, 0},
		{-4437, 1},
		{-5871, 1},
		{-4298, 0},
		{-2027, 1},
		{-2796, 0},
		{-2209, 1},
		{-6094, 0},
		{3257, 1},
		{-4732, 0},
		{7495, 1},
		{-3916, 0},
		{1469, 1},
		{6164, 0},
		{-7545, 1},
		{-7763, 1},
		{7523, 1},
		{-8076, 0},
		{-3778, 0},
		{-1648, 0},
		{4220, 0},
		{-8551, 1},
		{9692, 0},
		{5394, 0},
		{2472, 0},
		{4056, 0},
		{5769, 1},
		{-2322, 0},
		{503, 1},
		{-8721, 1},
		{-6344, 0},
		{-4335, 1},
		{1677, 1},
		{-1703, 1},
		{-4086, 0},
		{7076, 0},
		{-7165, 1},
		{7636, 0},
		{-8043, 1},
		{-3753, 1},
		{4007, 1},
		{-261, 1},
		{-6538, 0},
		{9766, 0},
		{-7563, 1},
		{-7944, 0},
		{8922, 0},
		{-5759, 1},
		{-8791, 1},
		{-2211, 1},
		{3493, 1},
		{5573, 1},
		{-2645, 1},
		{-603656, 0},
		{807727, 1},
		{-69847, 1},
		{-843676, 0},
		{-830961, 1},
		{-608772, 0},
		{931043, 1},
		{855512, 0},
		{358482, 0},
		{-98919, 1},
		{215211, 1},
		{-933334, 0},
		{-613634, 0},
		{-95643, 1},
		{53934, 0},
		{161818, 0},
		{67041621, 1},
		{99662694, 0},
		{-94392019, 1},
		{-20543495, 1},
		{-27591794, 0},
		{-8314396, 0},
		{97455764, 0},
		{59367920, 0},
		{26856309, 1},
		{64178815, 1},
		{-11480741, 1},
		{45428276, 0},
		{46193175, 1},
		{-31079636, 0},
		{63115980, 0},
		{42559270, 0},
		{-1645871504, 0},
		{-1233918598, 0},
		{722012346, 0},
		{-1525999934, 0},
		{-365543955, 1},
		{2008798151, 1},
		{-1300713468, 0},
		{1425587979, 1},
		{1324445673, 1},
		{2136612365, 1},
		{-995371213, 1},
		{-2048365905, 1},
		{2096138065, 1},
		{-768738192, 0},
		{-846034014, 0},
		{411817058, 0},
		{2147483647, 1},
		{-2147483648, 0},
	}

	for i := range sample {
		if sample[i][1] == 0 {
			if !IsEven(sample[i][0]) {
				t.FailNow()
			}
		} else {
			if IsEven(sample[i][0]) {
				t.FailNow()
			}
		}
	}
}

func TestIsOdd(t *testing.T) {
	sample := [][2]int64{
		{0, 0},
		{1, 1},
		{-1, 1},
		{60, 0},
		{62, 0},
		{42, 0},
		{-87, 1},
		{-98, 0},
		{-56, 0},
		{41, 1},
		{14, 0},
		{20, 0},
		{-63, 1},
		{-81, 1},
		{-39, 1},
		{33, 1},
		{-32, 0},
		{-50, 0},
		{-1471, 1},
		{4046, 0},
		{5019, 1},
		{4521, 1},
		{-8744, 0},
		{4787, 1},
		{4973, 1},
		{-1184, 0},
		{-4437, 1},
		{-5871, 1},
		{-4298, 0},
		{-2027, 1},
		{-2796, 0},
		{-2209, 1},
		{-6094, 0},
		{3257, 1},
		{-4732, 0},
		{7495, 1},
		{-3916, 0},
		{1469, 1},
		{6164, 0},
		{-7545, 1},
		{-7763, 1},
		{7523, 1},
		{-8076, 0},
		{-3778, 0},
		{-1648, 0},
		{4220, 0},
		{-8551, 1},
		{9692, 0},
		{5394, 0},
		{2472, 0},
		{4056, 0},
		{5769, 1},
		{-2322, 0},
		{503, 1},
		{-8721, 1},
		{-6344, 0},
		{-4335, 1},
		{1677, 1},
		{-1703, 1},
		{-4086, 0},
		{7076, 0},
		{-7165, 1},
		{7636, 0},
		{-8043, 1},
		{-3753, 1},
		{4007, 1},
		{-261, 1},
		{-6538, 0},
		{9766, 0},
		{-7563, 1},
		{-7944, 0},
		{8922, 0},
		{-5759, 1},
		{-8791, 1},
		{-2211, 1},
		{3493, 1},
		{5573, 1},
		{-2645, 1},
		{-603656, 0},
		{807727, 1},
		{-69847, 1},
		{-843676, 0},
		{-830961, 1},
		{-608772, 0},
		{931043, 1},
		{855512, 0},
		{358482, 0},
		{-98919, 1},
		{215211, 1},
		{-933334, 0},
		{-613634, 0},
		{-95643, 1},
		{53934, 0},
		{161818, 0},
		{67041621, 1},
		{99662694, 0},
		{-94392019, 1},
		{-20543495, 1},
		{-27591794, 0},
		{-8314396, 0},
		{97455764, 0},
		{59367920, 0},
		{26856309, 1},
		{64178815, 1},
		{-11480741, 1},
		{45428276, 0},
		{46193175, 1},
		{-31079636, 0},
		{63115980, 0},
		{42559270, 0},
		{-1645871504, 0},
		{-1233918598, 0},
		{722012346, 0},
		{-1525999934, 0},
		{-365543955, 1},
		{2008798151, 1},
		{-1300713468, 0},
		{1425587979, 1},
		{1324445673, 1},
		{2136612365, 1},
		{-995371213, 1},
		{-2048365905, 1},
		{2096138065, 1},
		{-768738192, 0},
		{-846034014, 0},
		{411817058, 0},
		{2147483647, 1},
		{-2147483648, 0},
	}

	for i := range sample {
		if sample[i][1] == 0 {
			if IsOdd(sample[i][0]) {
				t.FailNow()
			}
		} else {
			if !IsOdd(sample[i][0]) {
				t.FailNow()
			}
		}
	}
}

func TestIsPrime(t *testing.T) {
	sample := [][2]int64{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 4},
		{5, 2},
		{6, 4},
		{7, 2},
		{8, 4},
		{9, 4},
		{10, 4},
		{11, 2},
		{12, 4},
		{13, 2},
		{14, 4},
		{15, 4},
		{16, 4},
		{17, 2},
		{18, 4},
		{19, 2},
		{20, 4},
		{21, 4},
		{22, 4},
		{23, 2},
		{24, 4},
		{25, 4},
		{26, 4},
		{27, 4},
		{28, 4},
		{29, 2},
		{30, 4},
		{31, 2},
		{32, 4},
		{33, 4},
		{34, 4},
		{35, 4},
		{36, 4},
		{37, 2},
		{38, 4},
		{39, 4},
		{40, 4},
		{41, 2},
		{42, 4},
		{43, 2},
		{44, 4},
		{45, 4},
		{46, 4},
		{47, 2},
		{48, 4},
		{49, 4},
		{50, 4},
		{51, 4},
		{52, 4},
		{53, 2},
		{54, 4},
		{55, 4},
		{56, 4},
		{57, 4},
		{58, 4},
		{59, 2},
		{60, 4},
		{61, 2},
		{62, 4},
		{63, 4},
		{64, 4},
		{65, 4},
		{66, 4},
		{67, 2},
		{68, 4},
		{69, 4},
		{70, 4},
		{71, 2},
		{72, 4},
		{73, 2},
		{74, 4},
		{75, 4},
		{76, 4},
		{77, 4},
		{78, 4},
		{79, 2},
		{80, 4},
		{81, 4},
		{82, 4},
		{83, 2},
		{84, 4},
		{85, 4},
		{86, 4},
		{87, 4},
		{88, 4},
		{89, 2},
		{90, 4},
		{91, 4},
		{92, 4},
		{93, 4},
		{94, 4},
		{95, 4},
		{96, 4},
		{97, 2},
		{98, 4},
		{99, 4},
		{100, 4},
		{101, 2},
		{102, 4},
		{103, 2},
		{104, 4},
		{105, 4},
		{106, 4},
		{107, 2},
		{108, 4},
		{109, 2},
		{110, 4},
		{111, 4},
		{112, 4},
		{113, 2},
		{114, 4},
		{115, 4},
		{116, 4},
		{117, 4},
		{118, 4},
		{119, 4},
		{120, 4},
		{121, 4},
		{122, 4},
		{123, 4},
		{124, 4},
		{125, 4},
		{126, 4},
		{127, 2},
		{170, 4},
		{271, 2},
		{357, 4},
		{225, 4},
		{440, 4},
		{235, 4},
		{274, 4},
		{383, 2},
		{313, 2},
		{434, 4},
		{222, 4},
		{137, 2},
		{155, 4},
		{465, 4},
		{249, 4},
		{141, 4},
		{5307, 4},
		{6374, 4},
		{951, 4},
		{7236, 4},
		{2027, 2},
		{2382, 4},
		{2251, 2},
		{757, 2},
		{6678, 4},
		{3206, 4},
		{6257, 2},
		{615, 4},
		{1677, 4},
		{2086, 4},
		{4158, 4},
		{4723, 2},
		{7755, 4},
		{1583, 2},
		{6547, 2},
		{4219, 2},
		{660, 4},
		{7698, 4},
		{3121, 2},
		{3881, 2},
		{2541, 4},
		{4696, 4},
		{4091, 2},
		{5222, 4},
		{4021, 2},
		{5839, 2},
		{6048, 4},
		{5521, 2},
		{47908, 4},
		{35060, 4},
		{47361, 4},
		{48859, 2},
		{56067, 4},
		{70451, 2},
		{9735, 4},
		{19763, 2},
		{97943, 2},
		{9933, 4},
		{61651, 2},
		{11407, 4},
		{8774, 4},
		{48383, 2},
		{51001, 2},
		{73029, 4},
		{27690, 4},
		{30466, 4},
		{71479, 2},
		{84701, 2},
		{28643, 2},
		{57075, 4},
		{99745, 4},
		{100921, 4},
		{40496, 4},
		{9798, 4},
		{41603, 2},
		{46912, 4},
		{49852, 4},
		{55871, 2},
		{10993, 2},
		{79657, 2},
		{609680, 4},
		{180540, 4},
		{147672, 4},
		{819031, 2},
		{149623, 2},
		{1056048, 4},
		{483389, 2},
		{452831, 2},
		{415109, 2},
		{185021, 2},
		{715823, 2},
		{744081, 4},
		{1276157, 2},
		{192978, 4},
		{631537, 2},
		{554226, 4},
		{653111, 2},
		{607346, 4},
		{452539, 2},
		{815939, 2},
		{247199, 4},
		{1245953, 2},
		{974803, 2},
		{185813, 2},
		{1261831, 2},
		{443227, 2},
		{1057294, 4},
		{427241, 2},
		{627391, 2},
		{1019663, 2},
		{629142, 4},
		{164503, 2},
		{6006421, 2},
		{9499199, 2},
		{12598247, 2},
		{13919909, 2},
		{8975950, 4},
		{6655578, 4},
		{2388697, 2},
		{14018237, 4},
		{7871261, 2},
		{1678013, 2},
		{2654027, 2},
		{10142801, 2},
		{2291487, 4},
		{3893849, 2},
		{1308913, 4},
		{14162880, 4},
	}

	for i := range sample {
		switch sample[i][1] {
		case 1:
			if IsPrime(sample[i][0]) {
				t.FailNow()
			}
		case 2:
			if !IsPrime(sample[i][0]) {
				t.FailNow()
			}
		default:
			if IsPrime(sample[i][0]) {
				t.FailNow()
			}
		}
	}
}

func TestIsComposite(t *testing.T) {
	sample := [][2]int64{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 4},
		{5, 2},
		{6, 4},
		{7, 2},
		{8, 4},
		{9, 4},
		{10, 4},
		{11, 2},
		{12, 4},
		{13, 2},
		{14, 4},
		{15, 4},
		{16, 4},
		{17, 2},
		{18, 4},
		{19, 2},
		{20, 4},
		{21, 4},
		{22, 4},
		{23, 2},
		{24, 4},
		{25, 4},
		{26, 4},
		{27, 4},
		{28, 4},
		{29, 2},
		{30, 4},
		{31, 2},
		{32, 4},
		{33, 4},
		{34, 4},
		{35, 4},
		{36, 4},
		{37, 2},
		{38, 4},
		{39, 4},
		{40, 4},
		{41, 2},
		{42, 4},
		{43, 2},
		{44, 4},
		{45, 4},
		{46, 4},
		{47, 2},
		{48, 4},
		{49, 4},
		{50, 4},
		{51, 4},
		{52, 4},
		{53, 2},
		{54, 4},
		{55, 4},
		{56, 4},
		{57, 4},
		{58, 4},
		{59, 2},
		{60, 4},
		{61, 2},
		{62, 4},
		{63, 4},
		{64, 4},
		{65, 4},
		{66, 4},
		{67, 2},
		{68, 4},
		{69, 4},
		{70, 4},
		{71, 2},
		{72, 4},
		{73, 2},
		{74, 4},
		{75, 4},
		{76, 4},
		{77, 4},
		{78, 4},
		{79, 2},
		{80, 4},
		{81, 4},
		{82, 4},
		{83, 2},
		{84, 4},
		{85, 4},
		{86, 4},
		{87, 4},
		{88, 4},
		{89, 2},
		{90, 4},
		{91, 4},
		{92, 4},
		{93, 4},
		{94, 4},
		{95, 4},
		{96, 4},
		{97, 2},
		{98, 4},
		{99, 4},
		{100, 4},
		{101, 2},
		{102, 4},
		{103, 2},
		{104, 4},
		{105, 4},
		{106, 4},
		{107, 2},
		{108, 4},
		{109, 2},
		{110, 4},
		{111, 4},
		{112, 4},
		{113, 2},
		{114, 4},
		{115, 4},
		{116, 4},
		{117, 4},
		{118, 4},
		{119, 4},
		{120, 4},
		{121, 4},
		{122, 4},
		{123, 4},
		{124, 4},
		{125, 4},
		{126, 4},
		{127, 2},
		{170, 4},
		{271, 2},
		{357, 4},
		{225, 4},
		{440, 4},
		{235, 4},
		{274, 4},
		{383, 2},
		{313, 2},
		{434, 4},
		{222, 4},
		{137, 2},
		{155, 4},
		{465, 4},
		{249, 4},
		{141, 4},
		{5307, 4},
		{6374, 4},
		{951, 4},
		{7236, 4},
		{2027, 2},
		{2382, 4},
		{2251, 2},
		{757, 2},
		{6678, 4},
		{3206, 4},
		{6257, 2},
		{615, 4},
		{1677, 4},
		{2086, 4},
		{4158, 4},
		{4723, 2},
		{7755, 4},
		{1583, 2},
		{6547, 2},
		{4219, 2},
		{660, 4},
		{7698, 4},
		{3121, 2},
		{3881, 2},
		{2541, 4},
		{4696, 4},
		{4091, 2},
		{5222, 4},
		{4021, 2},
		{5839, 2},
		{6048, 4},
		{5521, 2},
		{47908, 4},
		{35060, 4},
		{47361, 4},
		{48859, 2},
		{56067, 4},
		{70451, 2},
		{9735, 4},
		{19763, 2},
		{97943, 2},
		{9933, 4},
		{61651, 2},
		{11407, 4},
		{8774, 4},
		{48383, 2},
		{51001, 2},
		{73029, 4},
		{27690, 4},
		{30466, 4},
		{71479, 2},
		{84701, 2},
		{28643, 2},
		{57075, 4},
		{99745, 4},
		{100921, 4},
		{40496, 4},
		{9798, 4},
		{41603, 2},
		{46912, 4},
		{49852, 4},
		{55871, 2},
		{10993, 2},
		{79657, 2},
		{609680, 4},
		{180540, 4},
		{147672, 4},
		{819031, 2},
		{149623, 2},
		{1056048, 4},
		{483389, 2},
		{452831, 2},
		{415109, 2},
		{185021, 2},
		{715823, 2},
		{744081, 4},
		{1276157, 2},
		{192978, 4},
		{631537, 2},
		{554226, 4},
		{653111, 2},
		{607346, 4},
		{452539, 2},
		{815939, 2},
		{247199, 4},
		{1245953, 2},
		{974803, 2},
		{185813, 2},
		{1261831, 2},
		{443227, 2},
		{1057294, 4},
		{427241, 2},
		{627391, 2},
		{1019663, 2},
		{629142, 4},
		{164503, 2},
		{6006421, 2},
		{9499199, 2},
		{12598247, 2},
		{13919909, 2},
		{8975950, 4},
		{6655578, 4},
		{2388697, 2},
		{14018237, 4},
		{7871261, 2},
		{1678013, 2},
		{2654027, 2},
		{10142801, 2},
		{2291487, 4},
		{3893849, 2},
		{1308913, 4},
		{14162880, 4},
	}

	for i := range sample {
		switch sample[i][1] {
		case 1:
			if IsComposite(sample[i][0]) {
				t.FailNow()
			}
		case 2:
			if IsComposite(sample[i][0]) {
				t.FailNow()
			}
		default:
			if !IsComposite(sample[i][0]) {
				t.FailNow()
			}
		}
	}
}

func TestSieveOfPrimes(t *testing.T) {
	sample := [][2]int{
		{0, 0},
		{1, 0},
		{180, 41},
		{26, 9},
		{350, 70},
		{521, 98},
		{307, 63},
		{163, 38},
		{181, 42},
		{372, 73},
		{421, 82},
		{241, 53},
		{229, 50},
		{169, 39},
		{442, 85},
		{445, 86},
		{338, 68},
		{306, 62},
		{56, 16},
		{270, 57},
		{499, 95},
		{173, 40},
		{18, 7},
		{43, 14},
		{218, 47},
		{122, 30},
		{265, 56},
		{467, 91},
		{381, 75},
		{11, 5},
		{409, 80},
		{233, 51},
		{1119, 187},
		{6976, 896},
		{4355, 594},
		{7757, 984},
		{3467, 486},
		{5824, 764},
		{7724, 980},
		{1470, 232},
		{7894, 997},
		{6788, 873},
		{7873, 994},
		{6874, 885},
		{3554, 497},
		{4137, 569},
		{7878, 995},
		{3390, 477},
		{2141, 323},
		{4642, 626},
		{2718, 396},
		{7126, 912},
		{1022, 172},
		{5651, 743},
		{4822, 649},
		{2022, 306},
		{6449, 837},
		{2153, 325},
		{1323, 216},
		{6960, 893},
		{2837, 412},
		{4212, 576},
		{2179, 327},
		{1033, 174},
		{33023, 3540},
		{65798, 6572},
		{39761, 4179},
		{12379, 1478},
		{25339, 2794},
		{50093, 5143},
		{98641, 9471},
		{84959, 8272},
		{90697, 8778},
		{20898, 2350},
		{63079, 6326},
		{18369, 2104},
		{23459, 2611},
		{55337, 5620},
		{55057, 5596},
		{94463, 9111},
		{22769, 2544},
		{22501, 2516},
		{68369, 6799},
		{65315, 6524},
		{72893, 7206},
		{44969, 4672},
		{69511, 6902},
		{24371, 2705},
		{32579, 3497},
		{26017, 2862},
		{84309, 8217},
		{100043, 9595},
		{86861, 8440},
		{79355, 7776},
		{14445, 1693},
		{95418, 9197},
		{1209025, 93544},
		{526733, 43583},
		{832339, 66339},
		{350593, 30017},
		{616581, 50358},
		{787123, 62995},
		{474923, 39612},
		{320448, 27647},
		{617049, 50394},
		{232333, 20622},
		{651727, 52944},
		{524873, 43427},
		{1008852, 79161},
		{187573, 16970},
		{432235, 36339},
		{714479, 57629},
		{669287, 54254},
		{633133, 51574},
		{964012, 75920},
		{663127, 53800},
		{1257461, 96995},
		{172171, 15681},
		{1112392, 86634},
		{912533, 72211},
		{409529, 34574},
		{691575, 55914},
		{742768, 59730},
		{479430, 39964},
		{108289, 10303},
		{1228187, 94920},
		{709431, 57257},
		{1294061, 99610},
	}

	var arr []int

	for i := range sample {
		arr = SieveOfPrimes(sample[i][0])

		if len(arr) != sample[i][1] {
			t.FailNow()
		}

		for _, v := range arr {
			if !primality.IsPrime(int64(v)) {
				t.FailNow()
			}
		}
	}
}