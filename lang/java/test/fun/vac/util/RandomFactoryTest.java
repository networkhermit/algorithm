import fun.vac.util.RandomFactory;
import fun.vac.util.TestRunner;

public class RandomFactoryTest {

    public static boolean testGenerateEven() {
        RandomFactory.launch();

        int value;
        for (int i = 0; i < 8192; i++) {
            value = RandomFactory.generateEven();
            if ((value & 1) != 0) {
                return false;
            }
        }

        return true;
    }

    public static boolean testGenerateOdd() {
        RandomFactory.launch();

        int value;
        for (int i = 0; i < 8192; i++) {
            value = RandomFactory.generateOdd();
            if ((value & 1) == 0) {
                return false;
            }
        }

        return true;
    }

    public static void main(String[] args) {

        TestRunner.parseTest(testGenerateEven());

        TestRunner.parseTest(testGenerateOdd());
    }
}
