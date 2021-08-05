import muse.util.TestRunner;

public class TestRunnerTest {

    public static void testPick() {
        for (int i = 0; i < 10; i++) {
            boolean result = (i & 1) != 0;
            TestRunner.pick(() -> result);
        }
    }

    public static void main(String[] args) {
        testPick();
    }
}
