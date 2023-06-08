import random
import time


def generate_list(n):
    lst = [random.randint(0, n) for _ in range(n)]
    return lst


def bubble_sort(unsorted_list):
    list_length = len(unsorted_list)

    # Pass through the entire list as once per element
    for pass_number in range(list_length):
        # For each pass, go through the list up to the point where the end of the list
        # starts to be sorted from previous passes.
        for current_index in range(0, list_length - pass_number - 1):
            # If the current element is larger than the next one,
            # they are out of order and we swap them.
            if unsorted_list[current_index] > unsorted_list[current_index + 1]:
                unsorted_list[current_index], unsorted_list[current_index + 1] = (
                    unsorted_list[current_index + 1],
                    unsorted_list[current_index],
                )

    # The list is now sorted.
    return unsorted_list


def main():
    haystack = generate_list(1000)
    print(f"unsorted: {haystack}\n")

    start_time = time.time()
    print(f"sorted: {bubble_sort(haystack)}")
    duration = time.time() - start_time

    print(f"Execution time: {duration*1000} ms")


if __name__ == "__main__":
    main()
