def merge(arr: list, lo: int, mid: int, hi: int) -> None:
    if lo == mid:
        return

    merge(arr, lo, (lo + mid) >> 1, mid)
    merge(arr, mid, (mid + hi) >> 1, hi)

    m = lo
    n = mid

    sorted = [0] * (hi - lo)

    for i, _ in enumerate(sorted):
        if m != mid and (n == hi or arr[m] < arr[n]):
            sorted[i] = arr[m]
            m += 1
        else:
            sorted[i] = arr[n]
            n += 1

    cursor = 0
    for i in range(lo, hi):
        arr[i] = sorted[cursor]
        cursor += 1


def sort(arr: list) -> None:
    merge(arr, 0, len(arr) >> 1, len(arr))
