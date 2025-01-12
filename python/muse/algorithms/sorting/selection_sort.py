def sort(arr: list) -> None:
    if not arr:
        return

    for i in range(len(arr) - 1):
        i_min = i
        for j in range(i + 1, len(arr)):
            if arr[j] < arr[i_min]:
                i_min = j
        if i_min != i:
            arr[i], arr[i_min] = arr[i_min], arr[i]
