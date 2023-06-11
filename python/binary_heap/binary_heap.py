class BinaryHeap:
    def __init__(self):
        self.heap = [0]
        self.current_size = 0

    def sift_up(self, i):
        while i // 2 > 0:
            if (
                self.heap[i] < self.heap[i // 2]
            ):  # Swap if the child is less than the parent
                self.heap[i], self.heap[i // 2] = self.heap[i // 2], self.heap[i]
            i = i // 2

    def insert(self, k):
        self.heap.append(k)
        self.current_size += 1
        self.sift_up(self.current_size)

    def sift_down(self, i):
        while (i * 2) <= self.current_size:
            mc = self.min_child(i)
            if (
                self.heap[i] > self.heap[mc]
            ):  # Swap if the parent is greater than the child
                self.heap[i], self.heap[mc] = self.heap[mc], self.heap[i]
            i = mc

    def min_child(self, i):
        if i * 2 + 1 > self.current_size:
            return i * 2
        else:
            if self.heap[i * 2] < self.heap[i * 2 + 1]:
                return i * 2
            else:
                return i * 2 + 1

    def delete_min(self):
        retval = self.heap[1]
        self.heap[1] = self.heap[self.current_size]
        self.current_size -= 1
        self.heap.pop()
        self.sift_down(1)
        return retval


def main():
    # Create a new BinaryHeap
    heap = BinaryHeap()

    # Insert some values
    print("Inserting values...")
    heap.insert(5)
    heap.insert(9)
    heap.insert(1)
    heap.insert(7)
    heap.insert(3)

    print("Heap after insertions: ", heap.heap)

    # Delete the minimum value
    print("Deleting minimum value...")
    min_value = heap.delete_min()
    print("Minimum value: ", min_value)

    print("Heap after deletion: ", heap.heap)


if __name__ == "__main__":
    main()
