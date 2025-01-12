def find(arr: list[int], key: int) -> int:
    lo = 0
    hi = len(arr)

    while lo < hi:
        mid = lo + ((hi - lo) >> 1)
        if key < arr[mid]:
            hi = mid
        elif key > arr[mid]:
            lo = mid + 1
        else:
            return mid

    return len(arr)
