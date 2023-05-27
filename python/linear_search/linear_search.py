def linear_search(haystack, needle):
    for i in range(len(haystack)):
        if haystack[i] == needle:
            return i  # needle found
    return -1  # needle not found
