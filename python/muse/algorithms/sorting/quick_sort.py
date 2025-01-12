def partition(arr: list[int], lo: int, hi: int) -> int:
    pivot = arr[lo]

    i = lo
    j = hi - 1

    while i != j:
        while j > i:
            if arr[j] < pivot:
                arr[i], arr[j] = arr[j], pivot
                break
            j -= 1
        while i < j:
            if arr[i] > pivot:
                arr[j], arr[i] = arr[i], pivot
                break
            i += 1

    return i


def sort(arr: list[int]) -> None:
    sort_with_range(arr, 0, len(arr))


def sort_with_range(arr: list[int], lo: int, hi: int) -> None:
    if lo >= hi:
        return

    p = partition(arr, lo, hi)

    sort_with_range(arr, lo, p)
    sort_with_range(arr, p + 1, hi)
