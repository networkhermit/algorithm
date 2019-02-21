import fun.vac.util.TestRunner;

public class TestRunnerTest {

    public static void testParseTest() {
        for (int i = 0; i < 10; i++) {
            if ((i & 1) == 0) {
                TestRunner.parseTest(false);
            } else {
                TestRunner.parseTest(true);
            }
        }
    }

    public static void main(String[] args) {
        testParseTest();
    }
}
