import muse.algorithms.number_theory.Parity;
import muse.util.TestRunner;

public class ParityTest {

    public static boolean testParity() {
        long[][] sample = {
            {             0, 0},
            {             1, 1},
            {            -1, 1},
            {            60, 0},
            {            62, 0},
            {            42, 0},
            {           -87, 1},
            {           -98, 0},
            {           -56, 0},
            {            41, 1},
            {            14, 0},
            {            20, 0},
            {           -63, 1},
            {           -81, 1},
            {           -39, 1},
            {            33, 1},
            {           -32, 0},
            {           -50, 0},
            {         -1471, 1},
            {          4046, 0},
            {          5019, 1},
            {          4521, 1},
            {         -8744, 0},
            {          4787, 1},
            {          4973, 1},
            {         -1184, 0},
            {         -4437, 1},
            {         -5871, 1},
            {         -4298, 0},
            {         -2027, 1},
            {         -2796, 0},
            {         -2209, 1},
            {         -6094, 0},
            {          3257, 1},
            {         -4732, 0},
            {          7495, 1},
            {         -3916, 0},
            {          1469, 1},
            {          6164, 0},
            {         -7545, 1},
            {         -7763, 1},
            {          7523, 1},
            {         -8076, 0},
            {         -3778, 0},
            {         -1648, 0},
            {          4220, 0},
            {         -8551, 1},
            {          9692, 0},
            {          5394, 0},
            {          2472, 0},
            {          4056, 0},
            {          5769, 1},
            {         -2322, 0},
            {           503, 1},
            {         -8721, 1},
            {         -6344, 0},
            {         -4335, 1},
            {          1677, 1},
            {         -1703, 1},
            {         -4086, 0},
            {          7076, 0},
            {         -7165, 1},
            {          7636, 0},
            {         -8043, 1},
            {         -3753, 1},
            {          4007, 1},
            {          -261, 1},
            {         -6538, 0},
            {          9766, 0},
            {         -7563, 1},
            {         -7944, 0},
            {          8922, 0},
            {         -5759, 1},
            {         -8791, 1},
            {         -2211, 1},
            {          3493, 1},
            {          5573, 1},
            {         -2645, 1},
            {       -603656, 0},
            {        807727, 1},
            {        -69847, 1},
            {       -843676, 0},
            {       -830961, 1},
            {       -608772, 0},
            {        931043, 1},
            {        855512, 0},
            {        358482, 0},
            {        -98919, 1},
            {        215211, 1},
            {       -933334, 0},
            {       -613634, 0},
            {        -95643, 1},
            {         53934, 0},
            {        161818, 0},
            {    67_041_621, 1},
            {    99_662_694, 0},
            {   -94_392_019, 1},
            {   -20_543_495, 1},
            {   -27_591_794, 0},
            {    -8_314_396, 0},
            {    97_455_764, 0},
            {    59_367_920, 0},
            {    26_856_309, 1},
            {    64_178_815, 1},
            {   -11_480_741, 1},
            {    45_428_276, 0},
            {    46_193_175, 1},
            {   -31_079_636, 0},
            {    63_115_980, 0},
            {    42_559_270, 0},
            {-1_645_871_504, 0},
            {-1_233_918_598, 0},
            {   722_012_346, 0},
            {-1_525_999_934, 0},
            {  -365_543_955, 1},
            { 2_008_798_151, 1},
            {-1_300_713_468, 0},
            { 1_425_587_979, 1},
            { 1_324_445_673, 1},
            { 2_136_612_365, 1},
            {  -995_371_213, 1},
            {-2_048_365_905, 1},
            { 2_096_138_065, 1},
            {  -768_738_192, 0},
            {  -846_034_014, 0},
            {   411_817_058, 0},
            { 2_147_483_647, 1},
            {-2_147_483_648, 0},
        };

        for (int i = 0, size = sample.length; i < size; i++) {
            if (sample[i][1] == 0) {
                if (!Parity.moduloIsEven(sample[i][0])) {
                    return false;
                }
            } else {
                if (Parity.moduloIsEven(sample[i][0])) {
                    return false;
                }
            }
        }

        for (int i = 0, size = sample.length; i < size; i++) {
            if (sample[i][1] == 0) {
                if (!Parity.bitwiseIsEven(sample[i][0])) {
                    return false;
                }
            } else {
                if (Parity.bitwiseIsEven(sample[i][0])) {
                    return false;
                }
            }
        }

        for (int i = 0, size = sample.length; i < size; i++) {
            if (sample[i][1] == 0) {
                if (Parity.moduloIsOdd(sample[i][0])) {
                    return false;
                }
            } else {
                if (!Parity.moduloIsOdd(sample[i][0])) {
                    return false;
                }
            }
        }

        for (int i = 0, size = sample.length; i < size; i++) {
            if (sample[i][1] == 0) {
                if (Parity.bitwiseIsOdd(sample[i][0])) {
                    return false;
                }
            } else {
                if (!Parity.bitwiseIsOdd(sample[i][0])) {
                    return false;
                }
            }
        }

        return true;
    }

    public static void main(String[] args) {
        TestRunner.pick(ParityTest::testParity);
    }
}
