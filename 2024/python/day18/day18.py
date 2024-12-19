from library.algorithms.dijkstra import dijkstra_with_neighbors
from library.structures.MinHeap import MinHeap
from typing import Tuple

type Location = tuple[int, int]

def parse_input(filename) -> list[Location]:
    l = []
    with open(filename) as f:
        for line in f:
            a, b = line.strip().split(',')
            l.append((int(a), int(b)))
    return l


def part1(steps: int, width: int, height: int, falling: list[Location]) -> int:
    graph = []
    for row in range(height):
        for col in range(width):
            graph.append((col, row))

    queue = MinHeap[Tuple[int, Location]]()
    dist = {}
    prev = {}
    queue.insert((0, (0,0)))
    dist[(0,0)] = 0

    for i in range(steps):
        node = falling[i]
        graph.remove(node)

    step = 0
    while queue.size() > 0:
        current_score, current_node = queue.pop()

        # If we've already found a shorter path, we don't need to investigate any more.
        if dist[current_node] < current_score:
            continue

        # If this is not a valid node, continue
        if current_node not in graph:
            continue

        up = (current_node[0] - 1, current_node[1])
        down = (current_node[0] + 1, current_node[1])
        left = (current_node[0], current_node[1] - 1)
        right = (current_node[0], current_node[1] + 1)
        neighbors = []
        if current_node[0] > 0:
            neighbors.append(up)
        if current_node[0] < height - 1:
            neighbors.append(down)
        if current_node[1] > 0:
            neighbors.append(left)
        if current_node[1] < width - 1:
            neighbors.append(right)
        for next_node in neighbors:
            new_score = current_score + 1
            if next_node not in dist or dist[next_node] > new_score:
                dist[next_node] = new_score
                queue.insert((new_score, next_node))
                prev[next_node] = current_node


    visited = set()
    visited.add((width - 1, height - 1))
    next_node = prev.get((width - 1, height - 1))
    while next_node is not None:
        visited.add(next_node)
        next_node = prev.get(next_node)


    # for row in range(height):
    #     s = ""
    #     for col in range(width):
    #         if (col, row) in visited:
    #              s += 'O'
    #         elif (col, row) in falling[:steps]:
    #             s += '#'
    #         else:
    #             s += '.'
    #     print(s)


    return dist[width - 1, height - 1]


def part1_walk_while_falling(width: int, height: int, falling: list[Location]) -> int:
    """
    Let's just try dijkstra after everything has fallen
    :param width:
    :param height:
    :param falling:
    :return:
    """

    graph = []
    for row in range(height):
        for col in range(width):
            graph.append((col, row))

    queue = MinHeap[Tuple[int, Location]]()
    dist = {}
    prev = {}
    queue.insert((0, (0,0)))
    dist[(0,0)] = 0
    fell = set()

    step = 0
    while queue.size() > 0:
        current_score, current_node = queue.pop()


        # If we've already found a shorter path, we don't need to investigate any more.
        if dist[current_node] < current_score:
            continue

        # If this is not a valid node, continue
        if current_node not in graph:
            continue

        up = (current_node[0] - 1, current_node[1])
        down = (current_node[0] + 1, current_node[1])
        left = (current_node[0], current_node[1] - 1)
        right = (current_node[0], current_node[1] + 1)
        neighbors = []
        if current_node[0] > 0:
            neighbors.append(up)
        if current_node[0] < height - 1:
            neighbors.append(down)
        if current_node[1] > 0:
            neighbors.append(left)
        if current_node[1] < width - 1:
            neighbors.append(right)
        for next_node in neighbors:
            new_score = current_score + 1
            if next_node not in dist or dist[next_node] > new_score:
                dist[next_node] = new_score
                queue.insert((new_score, next_node))
                prev[next_node] = current_node

        if step < 12:
            block = falling[step]
            if block in graph:
                graph.remove(block)
                fell.add(block)
        step += 1

    visited = set()
    visited.add((width - 1, height - 1))
    next_node = prev.get((width - 1, height - 1))
    while next_node is not None:
        visited.add(next_node)
        next_node = prev.get(next_node)


    # for row in range(height):
    #     s = ""
    #     for col in range(width):
    #         # if (col, row) in visited:
    #         #      s += 'O'
    #         if (col, row) in fell:
    #             s += '#'
    #         else:
    #             s += '.'
    #     print(s)


    return dist[width - 1, height - 1]


def part2(lo, width, height, falling):
    hi = len(falling)
    result = None
    while lo < hi:
        mid = (lo + hi) // 2
        found = True
        print(f'Testing {mid}', end="... ")
        try:
            part1(mid, width, height, falling)
        except KeyError:
            found = False
        if found:
            # If we found a path, the number of steps is larger
            lo = mid + 1
            result = mid
        else:
            # If we did not find a path, the number of steps is smaller
            hi = mid
        print(found)
    return falling[result]

if __name__ == "__main__":
    puzzle_width = 71
    puzzle_height = 71
    s = 1024
    f = parse_input("data")
    print(part1(s, puzzle_width, puzzle_height, f))
    print(part2(s, puzzle_width, puzzle_height, f))
