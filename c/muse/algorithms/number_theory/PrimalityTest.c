#include <muse/algorithms/number_theory/Primality.h>
#include <muse/util/TestRunner.h>

bool testPrimality(void) {
    long sample[][2] = {
        // clang-format off
        {       0, 1},
        {       1, 1},
        {       2, 2},
        {       3, 2},
        {       4, 4},
        {       5, 2},
        {       6, 4},
        {       7, 2},
        {       8, 4},
        {       9, 4},
        {      10, 4},
        {      11, 2},
        {      12, 4},
        {      13, 2},
        {      14, 4},
        {      15, 4},
        {      16, 4},
        {      17, 2},
        {      18, 4},
        {      19, 2},
        {      20, 4},
        {      21, 4},
        {      22, 4},
        {      23, 2},
        {      24, 4},
        {      25, 4},
        {      26, 4},
        {      27, 4},
        {      28, 4},
        {      29, 2},
        {      30, 4},
        {      31, 2},
        {      32, 4},
        {      33, 4},
        {      34, 4},
        {      35, 4},
        {      36, 4},
        {      37, 2},
        {      38, 4},
        {      39, 4},
        {      40, 4},
        {      41, 2},
        {      42, 4},
        {      43, 2},
        {      44, 4},
        {      45, 4},
        {      46, 4},
        {      47, 2},
        {      48, 4},
        {      49, 4},
        {      50, 4},
        {      51, 4},
        {      52, 4},
        {      53, 2},
        {      54, 4},
        {      55, 4},
        {      56, 4},
        {      57, 4},
        {      58, 4},
        {      59, 2},
        {      60, 4},
        {      61, 2},
        {      62, 4},
        {      63, 4},
        {      64, 4},
        {      65, 4},
        {      66, 4},
        {      67, 2},
        {      68, 4},
        {      69, 4},
        {      70, 4},
        {      71, 2},
        {      72, 4},
        {      73, 2},
        {      74, 4},
        {      75, 4},
        {      76, 4},
        {      77, 4},
        {      78, 4},
        {      79, 2},
        {      80, 4},
        {      81, 4},
        {      82, 4},
        {      83, 2},
        {      84, 4},
        {      85, 4},
        {      86, 4},
        {      87, 4},
        {      88, 4},
        {      89, 2},
        {      90, 4},
        {      91, 4},
        {      92, 4},
        {      93, 4},
        {      94, 4},
        {      95, 4},
        {      96, 4},
        {      97, 2},
        {      98, 4},
        {      99, 4},
        {     100, 4},
        {     101, 2},
        {     102, 4},
        {     103, 2},
        {     104, 4},
        {     105, 4},
        {     106, 4},
        {     107, 2},
        {     108, 4},
        {     109, 2},
        {     110, 4},
        {     111, 4},
        {     112, 4},
        {     113, 2},
        {     114, 4},
        {     115, 4},
        {     116, 4},
        {     117, 4},
        {     118, 4},
        {     119, 4},
        {     120, 4},
        {     121, 4},
        {     122, 4},
        {     123, 4},
        {     124, 4},
        {     125, 4},
        {     126, 4},
        {     127, 2},
        {     170, 4},
        {     271, 2},
        {     357, 4},
        {     225, 4},
        {     440, 4},
        {     235, 4},
        {     274, 4},
        {     383, 2},
        {     313, 2},
        {     434, 4},
        {     222, 4},
        {     137, 2},
        {     155, 4},
        {     465, 4},
        {     249, 4},
        {     141, 4},
        {    5307, 4},
        {    6374, 4},
        {     951, 4},
        {    7236, 4},
        {    2027, 2},
        {    2382, 4},
        {    2251, 2},
        {     757, 2},
        {    6678, 4},
        {    3206, 4},
        {    6257, 2},
        {     615, 4},
        {    1677, 4},
        {    2086, 4},
        {    4158, 4},
        {    4723, 2},
        {    7755, 4},
        {    1583, 2},
        {    6547, 2},
        {    4219, 2},
        {     660, 4},
        {    7698, 4},
        {    3121, 2},
        {    3881, 2},
        {    2541, 4},
        {    4696, 4},
        {    4091, 2},
        {    5222, 4},
        {    4021, 2},
        {    5839, 2},
        {    6048, 4},
        {    5521, 2},
        {   47908, 4},
        {   35060, 4},
        {   47361, 4},
        {   48859, 2},
        {   56067, 4},
        {   70451, 2},
        {    9735, 4},
        {   19763, 2},
        {   97943, 2},
        {    9933, 4},
        {   61651, 2},
        {   11407, 4},
        {    8774, 4},
        {   48383, 2},
        {   51001, 2},
        {   73029, 4},
        {   27690, 4},
        {   30466, 4},
        {   71479, 2},
        {   84701, 2},
        {   28643, 2},
        {   57075, 4},
        {   99745, 4},
        {  100921, 4},
        {   40496, 4},
        {    9798, 4},
        {   41603, 2},
        {   46912, 4},
        {   49852, 4},
        {   55871, 2},
        {   10993, 2},
        {   79657, 2},
        {  609680, 4},
        {  180540, 4},
        {  147672, 4},
        {  819031, 2},
        {  149623, 2},
        { 1056048, 4},
        {  483389, 2},
        {  452831, 2},
        {  415109, 2},
        {  185021, 2},
        {  715823, 2},
        {  744081, 4},
        { 1276157, 2},
        {  192978, 4},
        {  631537, 2},
        {  554226, 4},
        {  653111, 2},
        {  607346, 4},
        {  452539, 2},
        {  815939, 2},
        {  247199, 4},
        { 1245953, 2},
        {  974803, 2},
        {  185813, 2},
        { 1261831, 2},
        {  443227, 2},
        { 1057294, 4},
        {  427241, 2},
        {  627391, 2},
        { 1019663, 2},
        {  629142, 4},
        {  164503, 2},
        { 6006421, 2},
        { 9499199, 2},
        {12598247, 2},
        {13919909, 2},
        { 8975950, 4},
        { 6655578, 4},
        { 2388697, 2},
        {14018237, 4},
        { 7871261, 2},
        { 1678013, 2},
        { 2654027, 2},
        {10142801, 2},
        { 2291487, 4},
        { 3893849, 2},
        { 1308913, 4},
        {14162880, 4},
        // clang-format on
    };

    for (size_t i = 0, size = *(&sample + 1) - sample; i < size; i++) {
        switch (sample[i][1]) {
        case 1:
            if (Primality_isPrime(sample[i][0])) {
                return false;
            }
            if (Primality_isComposite(sample[i][0])) {
                return false;
            }
            break;
        case 2:
            if (!Primality_isPrime(sample[i][0])) {
                return false;
            }
            if (Primality_isComposite(sample[i][0])) {
                return false;
            }
            break;
        default:
            if (Primality_isPrime(sample[i][0])) {
                return false;
            }
            if (!Primality_isComposite(sample[i][0])) {
                return false;
            }
        }
    }

    return true;
}

int main(void) {
    TestRunner_pick(&testPrimality);
}
