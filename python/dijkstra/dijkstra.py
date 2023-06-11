import heapq


def dijkstra(graph, start):
    # Initialize the shortest paths with infinity
    shortest_paths = {node: float("infinity") for node in graph}
    shortest_paths[start] = 0

    # Initialize the priority queue
    priority_queue = [(0, start)]

    while priority_queue:
        # Get the node with the smallest weight
        current_weight, current_node = heapq.heappop(priority_queue)

        # Ensure we don't process a node more than once
        if current_weight > shortest_paths[current_node]:
            continue

        # Examine the edges leading to the neighbors of the current node
        for neighbor, weight in graph[current_node].items():
            distance = current_weight + weight

            # If we found a shorter path to the neighbor node, update it
            if distance < shortest_paths[neighbor]:
                shortest_paths[neighbor] = distance
                heapq.heappush(priority_queue, (distance, neighbor))

    return shortest_paths


# Define a graph - each key is a node and each value is a dictionary of the neighboring nodes and the weights of the edges to them
graph = {
    "A": {"B": 2, "C": 5},
    "B": {"A": 2, "C": 4, "D": 6},
    "C": {"A": 5, "B": 4, "D": 3},
    "D": {"B": 6, "C": 3},
}

shortest_paths = dijkstra(graph, "A")

# Print the shortest paths
print(shortest_paths)
