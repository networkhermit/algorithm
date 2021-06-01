import muse.util.TestRunner;

public class TestRunnerTest {

    public static void testParseTest() {
        for (int i = 0; i < 10; i++) {
            TestRunner.parseTest((i & 1) != 0);
        }
    }

    public static void main(String[] args) {
        testParseTest();
    }
}
