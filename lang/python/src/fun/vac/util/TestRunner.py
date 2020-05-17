import sys

TestRunnerItemIndex = 0

def parseTest(ok: bool) -> None:
    global TestRunnerItemIndex

    if ok:
        print("✓ Item [%d] PASSED" % TestRunnerItemIndex)
    else:
        print("✗ Item [%d] FAILED" % TestRunnerItemIndex, file = sys.stderr)

    TestRunnerItemIndex += 1
