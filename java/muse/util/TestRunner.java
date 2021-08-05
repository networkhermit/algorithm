package muse.util;

import java.lang.Boolean;
import java.util.concurrent.Callable;

public final class TestRunner {

    private static int TestRunnerItemIndex = 0;

    private TestRunner() {}

    public static void pick(Callable<Boolean> func) {
        try {
            if (func.call()) {
                System.out.printf("✓ Item [%d] PASSED%n", TestRunnerItemIndex);
            } else {
                System.err.printf("✗ Item [%d] FAILED%n", TestRunnerItemIndex);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

        TestRunnerItemIndex++;
    }
}
