def find(arr: list, key: object) -> int:
    for i, v in enumerate(arr):
        if v == key:
            return i

    return len(arr)
