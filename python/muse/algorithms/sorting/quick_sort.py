def partition(arr: list, lo: int, hi: int) -> int:
    pivot = arr[lo]

    left = lo
    right = hi - 1

    while left != right:
        for i in range(right, left, -1):
            if arr[i] < pivot:
                arr[left] = arr[i]
                arr[i] = pivot
                break
            right -= 1
        for i in range(left, right, +1):
            if arr[i] > pivot:
                arr[right] = arr[i]
                arr[i] = pivot
                break
            left += 1

    return left


def sort(arr: list) -> None:
    sort_with_range(arr, 0, len(arr))


def sort_with_range(arr: list, lo: int, hi: int) -> None:
    if lo >= hi:
        return

    p = partition(arr, lo, hi)

    sort_with_range(arr, lo, p)
    sort_with_range(arr, p + 1, hi)
