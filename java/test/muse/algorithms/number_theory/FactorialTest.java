import muse.algorithms.number_theory.Factorial;
import muse.util.TestRunner;

public class FactorialTest {

    public static boolean testFactorial() {
        long[][] sample = {
            { 0,           1},
            { 1,           1},
            { 2,           2},
            { 3,           6},
            { 4,          24},
            { 5,         120},
            { 6,         720},
            { 7,        5040},
            { 8,       40320},
            { 9,      362880},
            {10,   3_628_800},
            {11,  39_916_800},
            {12, 479_001_600},
        };

        for (int i = 0, size = sample.length; i < size; i++) {
            if (Factorial.iterativeProcedure(sample[i][0]) != sample[i][1]) {
                return false;
            }
        }

        for (int i = 0, size = sample.length; i < size; i++) {
            if (Factorial.recursiveProcedure(sample[i][0]) != sample[i][1]) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) {
        TestRunner.parseTest(testFactorial());
    }
}
