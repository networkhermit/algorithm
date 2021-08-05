#include <muse/algorithms/number_theory/Parity.h>
#include <muse/util/TestRunner.h>

bool testParity(void) {
    long sample[][2] = {
        // clang-format off
        {          0, 0},
        {          1, 1},
        {         -1, 1},
        {         60, 0},
        {         62, 0},
        {         42, 0},
        {        -87, 1},
        {        -98, 0},
        {        -56, 0},
        {         41, 1},
        {         14, 0},
        {         20, 0},
        {        -63, 1},
        {        -81, 1},
        {        -39, 1},
        {         33, 1},
        {        -32, 0},
        {        -50, 0},
        {      -1471, 1},
        {       4046, 0},
        {       5019, 1},
        {       4521, 1},
        {      -8744, 0},
        {       4787, 1},
        {       4973, 1},
        {      -1184, 0},
        {      -4437, 1},
        {      -5871, 1},
        {      -4298, 0},
        {      -2027, 1},
        {      -2796, 0},
        {      -2209, 1},
        {      -6094, 0},
        {       3257, 1},
        {      -4732, 0},
        {       7495, 1},
        {      -3916, 0},
        {       1469, 1},
        {       6164, 0},
        {      -7545, 1},
        {      -7763, 1},
        {       7523, 1},
        {      -8076, 0},
        {      -3778, 0},
        {      -1648, 0},
        {       4220, 0},
        {      -8551, 1},
        {       9692, 0},
        {       5394, 0},
        {       2472, 0},
        {       4056, 0},
        {       5769, 1},
        {      -2322, 0},
        {        503, 1},
        {      -8721, 1},
        {      -6344, 0},
        {      -4335, 1},
        {       1677, 1},
        {      -1703, 1},
        {      -4086, 0},
        {       7076, 0},
        {      -7165, 1},
        {       7636, 0},
        {      -8043, 1},
        {      -3753, 1},
        {       4007, 1},
        {       -261, 1},
        {      -6538, 0},
        {       9766, 0},
        {      -7563, 1},
        {      -7944, 0},
        {       8922, 0},
        {      -5759, 1},
        {      -8791, 1},
        {      -2211, 1},
        {       3493, 1},
        {       5573, 1},
        {      -2645, 1},
        {    -603656, 0},
        {     807727, 1},
        {     -69847, 1},
        {    -843676, 0},
        {    -830961, 1},
        {    -608772, 0},
        {     931043, 1},
        {     855512, 0},
        {     358482, 0},
        {     -98919, 1},
        {     215211, 1},
        {    -933334, 0},
        {    -613634, 0},
        {     -95643, 1},
        {      53934, 0},
        {     161818, 0},
        {   67041621, 1},
        {   99662694, 0},
        {  -94392019, 1},
        {  -20543495, 1},
        {  -27591794, 0},
        {   -8314396, 0},
        {   97455764, 0},
        {   59367920, 0},
        {   26856309, 1},
        {   64178815, 1},
        {  -11480741, 1},
        {   45428276, 0},
        {   46193175, 1},
        {  -31079636, 0},
        {   63115980, 0},
        {   42559270, 0},
        {-1645871504, 0},
        {-1233918598, 0},
        {  722012346, 0},
        {-1525999934, 0},
        { -365543955, 1},
        { 2008798151, 1},
        {-1300713468, 0},
        { 1425587979, 1},
        { 1324445673, 1},
        { 2136612365, 1},
        { -995371213, 1},
        {-2048365905, 1},
        { 2096138065, 1},
        { -768738192, 0},
        { -846034014, 0},
        {  411817058, 0},
        { 2147483647, 1},
        {-2147483648, 0},
        // clang-format on
    };

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        if (sample[i][1] == 0) {
            if (!Parity_moduloIsEven(sample[i][0])) {
                return false;
            }
        } else {
            if (Parity_moduloIsEven(sample[i][0])) {
                return false;
            }
        }
    }

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        if (sample[i][1] == 0) {
            if (!Parity_bitwiseIsEven(sample[i][0])) {
                return false;
            }
        } else {
            if (Parity_bitwiseIsEven(sample[i][0])) {
                return false;
            }
        }
    }

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        if (sample[i][1] == 0) {
            if (Parity_moduloIsOdd(sample[i][0])) {
                return false;
            }
        } else {
            if (!Parity_moduloIsOdd(sample[i][0])) {
                return false;
            }
        }
    }

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        if (sample[i][1] == 0) {
            if (Parity_bitwiseIsOdd(sample[i][0])) {
                return false;
            }
        } else {
            if (!Parity_bitwiseIsOdd(sample[i][0])) {
                return false;
            }
        }
    }

    return true;
}

int main(void) {
    TestRunner_pick(&testParity);
}
