from typing import Protocol, runtime_checkable


@runtime_checkable
class Resizable(Protocol):
    def capacity(self) -> int: ...

    def shrink(self) -> None: ...
