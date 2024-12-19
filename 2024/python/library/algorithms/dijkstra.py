from library.structures.MinHeap import MinHeap
from typing import TypeVar, Callable, Tuple, List
from timeit import default_timer as timer

T = TypeVar('T')

def dijkstra(
        graph: List[T],
        start: T,
        get_neighbors: Callable[[T], List[Tuple[int, int]]]) -> (dict[T], dict[T, list[T]]):
    """
    Given a graph, compute the distance of every node from the starting node.
    :param graph: List of nodes in the graph
    :param start: The starting node
    :param get_neighbors: Function to get the neighbors of a node
    :return: a dictionary of Node -> Distance from start, dictionary of node -> prev for each node
    """
    queue = MinHeap[Tuple[int, T]]()
    dist = {}
    prev = {}
    queue.insert((0, start))
    dist[start] = 0

    start = timer()
    processed = 0
    while queue.size() > 0:
        current_score, current_node = queue.pop()

        # If we've already found a shorter path, we don't need to investigate any more.
        if dist[current_node] < current_score:
            continue

        # If this is not a valid node, continue
        if current_node not in graph:
            continue

        for next_node, weight in get_neighbors(current_node):
            new_score = current_score + weight
            if next_node not in dist or dist[next_node] > new_score:
                dist[next_node] = new_score
                queue.insert((new_score, next_node))
                prev[next_node] = [current_node]
            elif dist[next_node] == new_score:
                prev[next_node].append(current_node)


        processed += 1
        elapsed = timer() - start
        if elapsed > 10:
            print(f'Processed {processed} nodes in {timer() - start:.2f} seconds')
            start = timer()
    return dist, prev


def dijkstra_with_neighbors(
        graph: dict[T, List[tuple[T, int]]],
        start: T) -> (dict[T], dict[T, T]):
    """
    Given a graph, compute the distance of every node from the starting node.
    :param graph: List of nodes in the graph
    :param start: The starting node
    :param get_neighbors: Function to get the neighbors of a node
    :return: a dictionary of Node -> Distance from start, dictionary of node -> prev for each node
    """
    queue = MinHeap[Tuple[int, T]]()
    dist = {}
    prev = {}
    queue.insert((0, start))
    dist[start] = 0

    start = timer()
    processed = 0
    while queue.size() > 0:
        current_score, current_node = queue.pop()

        # If we've already found a shorter path, we don't need to investigate any more.
        if dist[current_node] < current_score:
            continue

        # If this is not a valid node, continue
        if current_node not in graph:
            continue

        for next_node, weight in graph[current_node]:
            new_score = current_score + weight
            if next_node not in dist or dist[next_node] > new_score:
                dist[next_node] = new_score
                queue.insert((new_score, next_node))
                prev[next_node] = current_node

        processed += 1
        elapsed = timer() - start
        if elapsed > 10:
            print(f'Processed {processed} nodes in {timer() - start:.2f} seconds')
            start = timer()
    return dist, prev