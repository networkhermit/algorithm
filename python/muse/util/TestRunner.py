from typing import Callable
import sys

itemIndex = 0

def pick(func: Callable[[], bool]) -> None:
    global itemIndex

    if func():
        print("✓ Item [%d] PASSED" % itemIndex)
    else:
        print("✗ Item [%d] FAILED" % itemIndex, file = sys.stderr)

    itemIndex += 1
