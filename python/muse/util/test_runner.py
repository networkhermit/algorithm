import sys
from typing import Callable

ItemIndex = 0


def pick(func: Callable[[], bool]) -> None:
    global ItemIndex

    if func():
        print("✓ Item [%d] PASSED" % ItemIndex)
    else:
        print("✗ Item [%d] FAILED" % ItemIndex, file=sys.stderr)

    ItemIndex += 1
