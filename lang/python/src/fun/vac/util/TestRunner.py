import sys

global TestRunner_TestIndex
TestRunner_TestIndex = 0

def parseTest(ok: bool) -> None:
    global TestRunner_TestIndex

    if ok:
        print("< Test [%d] Passed >" % TestRunner_TestIndex)
    else:
        print("X Test [%d] Failed X" % TestRunner_TestIndex, file = sys.stderr)

    TestRunner_TestIndex += 1
