from typing import Callable
import sys

TestRunnerItemIndex = 0

def pick(func: Callable[[], bool]) -> None:
    global TestRunnerItemIndex

    if func():
        print("✓ Item [%d] PASSED" % TestRunnerItemIndex)
    else:
        print("✗ Item [%d] FAILED" % TestRunnerItemIndex, file = sys.stderr)

    TestRunnerItemIndex += 1
