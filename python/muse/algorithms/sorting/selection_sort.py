def sort(arr: list) -> None:
    if not arr:
        return

    i_min = 0

    bound = len(arr) - 1
    for i in range(bound):
        i_min = i
        for j in range(i + 1, bound + 1):
            if arr[j] < arr[i_min]:
                i_min = j
        if i_min != i:
            arr[i], arr[i_min] = arr[i_min], arr[i]
