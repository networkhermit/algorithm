def sort(arr: list) -> None:
    margin = 0
    unsorted = len(arr)

    while unsorted > 1:
        margin = 0
        for i in range(1, unsorted):
            if arr[i - 1] > arr[i]:
                arr[i - 1], arr[i] = arr[i], arr[i - 1]
                margin = i
        unsorted = margin
