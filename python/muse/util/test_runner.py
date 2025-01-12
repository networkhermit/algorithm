import sys
from collections.abc import Callable

ItemIndex = 0


def pick(fn: Callable[[], bool]) -> None:
    global ItemIndex

    if fn():
        print("✓ Item [%d] PASSED" % ItemIndex)
    else:
        print("✗ Item [%d] FAILED" % ItemIndex, file=sys.stderr)

    ItemIndex += 1
