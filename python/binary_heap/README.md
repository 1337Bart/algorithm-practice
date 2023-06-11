A binary heap is a complete binary tree which satisfies the heap property. It is a tree-based data structure, that can satisfy either the "min-heap" property or the "max-heap" property. In a min heap, for any given node i, the value of i is greater than or equal to the value of its parent. This property must be recursively true for all nodes in the Binary Heap.

Operations
The main operations provided by a binary heap include:

insert: Adds a new element to the heap while maintaining the heap property. The time complexity of this operation is O(log n), where n is the number of nodes in the heap.

delete_min: Removes and returns the smallest element from the heap, and then restructures the heap to maintain the heap property. The time complexity of this operation is also O(log n).

Binary heap is a complete binary tree. This means that it can be efficiently implemented using an array, where for any given index i, its left child is located at index 2i, and its right child is located at index 2i + 1. Its parent is located at index i // 2.