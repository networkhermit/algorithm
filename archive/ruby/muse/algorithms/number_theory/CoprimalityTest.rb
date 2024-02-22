require "muse/algorithms/number_theory/Coprimality"
require "muse/util/TestRunner"

def testCoprimality
  sample = [
    [0, 1, 1],
    [1, 0, 1],
    [1, 1, 1],
    [-1, -1, 1],
    [15, -45, 0],
    [-7, 51, 1],
    [-3, 75, 0],
    [-84, -32, 0],
    [-21, -71, 1],
    [89, -11, 1],
    [-40, -9, 1],
    [92, 26, 0],
    [14, -95, 1],
    [5, -99, 1],
    [-45, -57, 0],
    [59, -97, 1],
    [-16, 37, 1],
    [97, 49, 1],
    [5998, -3871, 1],
    [-2152, -2909, 1],
    [8823, 629, 0],
    [7589, -7035, 1],
    [-8669, 6398, 1],
    [1059, 3522, 0],
    [9892, -2990, 0],
    [-1819, 4740, 1],
    [7900, -7244, 0],
    [8072, -1226, 0],
    [-6875, 5777, 1],
    [-828, 8859, 0],
    [-3929, 1143, 1],
    [-5397, 1880, 1],
    [-9397, -4343, 1],
    [2389, -4026, 1],
    [-6838, -7802, 0],
    [-6118, 7944, 0],
    [1408, -8483, 1],
    [-507, 3510, 0],
    [5030, -4818, 0],
    [2563, -1867, 1],
    [3317, -4585, 1],
    [-2448, -2215, 1],
    [-1982, -3811, 1],
    [645, 7907, 1],
    [5812, 1657, 1],
    [3944, -5997, 1],
    [-2410, -6522, 0],
    [4565, 9055, 0],
    [8478, 1099, 0],
    [6444, -7298, 0],
    [7606, -7006, 0],
    [2491, -2017, 1],
    [7151, -9724, 1],
    [2958, -5697, 0],
    [5134, -7701, 0],
    [-3544, 9042, 0],
    [8826, -4548, 0],
    [6794, -5322, 0],
    [-6906, 7502, 0],
    [172, 8480, 0],
    [8244, 4622, 0],
    [-7315, 5253, 1],
    [7011, 4285, 1],
    [-3132, -4999, 1],
    [-7332, -655, 1],
    [9661, -1750, 1],
    [7560, -3128, 0],
    [1334, 6234, 0],
    [8075, 6450, 0],
    [3283, 8980, 1],
    [3364, 8482, 0],
    [4909, 8302, 1],
    [9332, -2523, 1],
    [-8515, 5209, 1],
    [-1624, 7640, 0],
    [-6463, 2562, 1],
    [7912, 5868, 0],
    [-4825, -2173, 1],
    [28876, 30106, 0],
    [-929840, 887043, 1],
    [684923, -588038, 1],
    [-291411, -299801, 1],
    [-905447, -402122, 1],
    [830872, -223425, 1],
    [831033, -753398, 1],
    [-575558, 711716, 0],
    [-312296, 515492, 0],
    [595308, 963205, 0],
    [-770718, -388434, 0],
    [353889, 330806, 1],
    [-174566, 613742, 0],
    [-884075, 26687, 1],
    [390743, 204874, 1],
    [930615, 86524, 1],
    [-84_945_524, 15_427_487, 1],
    [76_038_602, -89_688_904, 0],
    [50_294_214, -65_802_481, 1],
    [83_436_075, 44_248_708, 1],
    [54_031_677, 92_370_464, 1],
    [-6_019_575, -1_302_029, 1],
    [86_952_873, -75_470_243, 1],
    [-53_215_372, 28_947_490, 0],
    [-37_109_442, 75_623_090, 0],
    [-11_656_183, 16_147_085, 1],
    [86_439_843, -12_134_914, 1],
    [-52_977_427, 48_349_835, 1],
    [16_234_161, 26_103_566, 1],
    [73_756_826, -6_586_797, 1],
    [6_836_355, 11_529_043, 1],
    [29_864_126, -80_782_077, 1],
    [1_646_787_325, 1_961_513_442, 1],
    [-1_755_035_190, -1_801_169_490, 0],
    [-1_509_018_194, 1_829_751_775, 1],
    [-720_160_017, -1_425_309_680, 1],
    [1_216_287_038, 1_821_933_798, 0],
    [-1_925_479_607, -1_842_455_762, 1],
    [-795_996_486, -1_859_155_567, 1],
    [-367_280_505, 321_267_794, 1],
    [829_304_526, -1_575_808_585, 1],
    [1_457_917_042, 1_083_382_210, 0],
    [-377_130_535, 1_526_538_188, 1],
    [-1_109_347_700, -1_819_333_109, 1],
    [-1_740_794_578, 278_770_346, 0],
    [448_156_480, -1_008_775_746, 0],
    [442_691_160, 1_680_572_092, 0],
    [1_241_208_470, -647_438_045, 0],
    [2_147_483_647, -561_158_902, 1],
    [761_395_308, -2_147_483_647, 1]
  ]

  sample.each_index do |i|
    if sample[i][2].zero?
      return false if Coprimality.reduceToBinaryGCD(sample[i][0], sample[i][1])
    else
      return false unless Coprimality.reduceToBinaryGCD(sample[i][0], sample[i][1])
    end
  end

  sample.each_index do |i|
    if sample[i][2].zero?
      return false if Coprimality.reduceToEuclidean(sample[i][0], sample[i][1])
    else
      return false unless Coprimality.reduceToEuclidean(sample[i][0], sample[i][1])
    end
  end

  true
end

TestRunner.pick(testCoprimality) if __FILE__ == $PROGRAM_NAME
