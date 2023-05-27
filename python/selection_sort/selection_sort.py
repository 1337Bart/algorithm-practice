def selection_sort(haystack):
    for i in range(len(haystack)):
        min_idx = i
        for j in range(i + 1, len(haystack)):
            if haystack[min_idx] > haystack[j]:
                min_idx = j

        haystack[i], haystack[min_idx] = haystack[min_idx], haystack[i]

    return haystack
