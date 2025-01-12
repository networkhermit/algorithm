from collections.abc import Callable


def unique_category_derive(
    fn: Callable[[object], bool], sample: list, c: object
) -> Callable[[], bool]:
    def f() -> bool:
        for tc in sample:
            actual = fn(tc["n"])
            expected = tc["category"] == c
            if actual != expected:
                return False
        return True

    return f


def mn_unique_category_derive(
    fn: Callable[[object, object], bool], sample: list, c: object
) -> Callable[[], bool]:
    def f() -> bool:
        for tc in sample:
            actual = fn(tc["m"], tc["n"])
            expected = tc["category"] == c
            if actual != expected:
                return False
        return True

    return f
