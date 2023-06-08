import random
import time


def generate_list(n):
    lst = [random.randint(0, n) for _ in range(n)]
    return lst


def selection_sort(unsorted_list):
    for current_index in range(len(unsorted_list)):
        # At the start of each iteration, we assume the current index is the index of the smallest element.
        smallest_value_index = current_index

        # We compare the assumed smallest value with all other elements in the list.
        for comparator_index in range(current_index + 1, len(unsorted_list)):
            # If we find a smaller value, we update our smallest value index.
            if unsorted_list[smallest_value_index] > unsorted_list[comparator_index]:
                smallest_value_index = comparator_index

        # Once we've gone through all other elements, we swap the current index's value with the smallest value we found.
        unsorted_list[current_index], unsorted_list[smallest_value_index] = (
            unsorted_list[smallest_value_index],
            unsorted_list[current_index],
        )

    return unsorted_list


def main():
    haystack = generate_list(1000)
    print(f"unsorted: {haystack}\n")

    start_time = time.time()
    print(f"sorted: {selection_sort(haystack)}")
    duration = time.time() - start_time

    print(f"Execution time: {duration*1000} ms")


if __name__ == "__main__":
    main()
