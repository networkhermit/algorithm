from collections.abc import Callable
from dataclasses import dataclass
from typing import Any


@dataclass
class UniqueCategorySample:
    n: Any
    category: Any


def unique_category_derive(
    fn: Callable[[Any], bool], sample: list[UniqueCategorySample], c: Any
) -> Callable[[], bool]:
    def f() -> bool:
        for tc in sample:
            actual = fn(tc.n)
            expected = tc.category == c
            if actual != expected:
                return False
        return True

    return f


@dataclass
class MNUniqueCategorySample:
    m: Any
    n: Any
    category: Any


def mn_unique_category_derive(
    fn: Callable[[Any, Any], bool], sample: list[MNUniqueCategorySample], c: Any
) -> Callable[[], bool]:
    def f() -> bool:
        for tc in sample:
            actual = fn(tc.m, tc.n)
            expected = tc.category == c
            if actual != expected:
                return False
        return True

    return f
