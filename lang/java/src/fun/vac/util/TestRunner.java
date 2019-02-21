package fun.vac.util;

public class TestRunner {

    private static int TestRunner_TestIndex = 0;

    private TestRunner() {}

    public static void parseTest(boolean ok) {
        if (ok) {
            System.out.printf("< Test [%d] Passed >\n", TestRunner_TestIndex);
        } else {
            System.err.printf("X Test [%d] Failed X\n", TestRunner_TestIndex);
        }

        TestRunner_TestIndex += 1;
    }
}
