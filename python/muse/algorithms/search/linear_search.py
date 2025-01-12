def find(arr: list[int], key: int) -> int:
    for i, v in enumerate(arr):
        if v == key:
            return i

    return len(arr)
