import random
import time


def generate_list(n):
    lst = [random.randint(0, n) for _ in range(n)]
    return lst


def insertion_sort(unsorted_list):
    # We start from the second element.
    for current_index in range(1, len(unsorted_list)):
        # This is the value we want to insert into the sorted part of the list.
        value_to_insert = unsorted_list[current_index]
        # We start comparing with the element before it.
        comparison_index = current_index - 1

        # While the comparison index is valid (not out of range)
        # and the value at this index is greater than the value we want to insert,
        # we shift the larger values to the right.
        while (
            comparison_index >= 0 and value_to_insert < unsorted_list[comparison_index]
        ):
            unsorted_list[comparison_index + 1] = unsorted_list[comparison_index]
            comparison_index -= 1

        # Now we have found the correct position for the value,
        # we insert it into its proper place in the sorted part of the list.
        unsorted_list[comparison_index + 1] = value_to_insert

    return unsorted_list


def main():
    unsorted_list = generate_list(1000)
    print(f"Unsorted: {unsorted_list}\n")

    start_time = time.time()
    sorted_list = insertion_sort(unsorted_list)
    duration = time.time() - start_time

    print(f"Sorted: {sorted_list}")
    print(f"Execution time: {duration*1000} ms")


if __name__ == "__main__":
    main()
