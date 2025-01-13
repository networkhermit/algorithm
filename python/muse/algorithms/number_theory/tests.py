from collections.abc import Callable
from dataclasses import dataclass


@dataclass
class UniqueCategorySample[T, C]:
    n: T
    category: C


def unique_category_derive[
    T, C
](fn: Callable[[T], bool], sample: list[UniqueCategorySample[T, C]], c: C) -> Callable[
    [], bool
]:
    def f() -> bool:
        for tc in sample:
            actual = fn(tc.n)
            expected = tc.category == c
            if actual != expected:
                return False
        return True

    return f


@dataclass
class MNUniqueCategorySample[T, C]:
    m: T
    n: T
    category: C


def mn_unique_category_derive[
    T, C
](
    fn: Callable[[T, T], bool], sample: list[MNUniqueCategorySample[T, C]], c: C
) -> Callable[[], bool]:
    def f() -> bool:
        for tc in sample:
            actual = fn(tc.m, tc.n)
            expected = tc.category == c
            if actual != expected:
                return False
        return True

    return f
