package fun.vac.util;

public final class TestRunner {

    private static int TestRunnerItemIndex = 0;

    private TestRunner() {}

    public static void parseTest(boolean ok) {
        if (ok) {
            System.out.printf("✓ Item [%d] PASSED%n", TestRunnerItemIndex);
        } else {
            System.err.printf("✗ Item [%d] FAILED%n", TestRunnerItemIndex);
        }

        TestRunnerItemIndex++;
    }
}
