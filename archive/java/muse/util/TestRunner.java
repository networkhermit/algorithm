package muse.util;

import java.util.concurrent.Callable;

public final class TestRunner {

  private static int itemIndex = 0;

  private TestRunner() {}

  public static void pick(Callable<Boolean> func) {
    try {
      if (func.call()) {
        System.out.printf("✓ Item [%d] PASSED%n", itemIndex);
      } else {
        System.err.printf("✗ Item [%d] FAILED%n", itemIndex);
      }
    } catch (Exception e) {
      e.printStackTrace();
      System.err.printf("✗ Item [%d] FAILED%n", itemIndex);
    }

    itemIndex++;
  }
}
