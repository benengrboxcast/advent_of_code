import pygame

from library.algorithms.dijkstra import dijkstra
from mazegrid import MazeGrid, Direction, Turn, MazeSpace
import math
from timeit import default_timer as timer
from sortedcontainers import SortedList

type Location = tuple[int, int]
type GraphNode = tuple[Location, Direction]

# Distance of the specified node from source
distances: dict[GraphNode, float] = {}

# The node that got us to this node
prev: dict[GraphNode, GraphNode] = {}

def get_next_location(location: Location, direction: Direction) -> Location:
    if direction == Direction.North:
        return location[0] - 1, location[1]
    elif direction == Direction.East:
        return location[0], location[1] + 1
    elif direction == Direction.South:
        return location[0] + 1, location[1]
    else:
        return location[0], location[1] - 1

def update_distance(current: GraphNode, distance: float, turn: Turn | None, unvisited: dict[GraphNode, float]):
    turn_cost = 1000
    forward_cost = 1

    if turn is None:
        location = get_next_location(current[0], current[1])
        distance += forward_cost
    else:
        location = current[0]
        distance += turn_cost
    direction = current[1].turn(turn)

    node = location, direction
    if node in unvisited and (unvisited[node] == math.nan or unvisited[node] > distance):
        unvisited[node] = distance
        distances[node] = distance
        prev[node] = current


def part1_old(g: MazeGrid) -> float:
    result = None

    # Dijkstra's Algorithm From https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm

    # 1. Create the unvisited set
    unvisited_set: dict[GraphNode, float] = {}
    goal = g.goal

    for space in g.spaces:
        unvisited_set[(space, Direction.East)] = math.inf
        unvisited_set[(space, Direction.West)] = math.inf
        unvisited_set[(space, Direction.North)] = math.inf
        unvisited_set[(space, Direction.South)] =  math.inf


    # 2. Assign initial node 0 distance
    unvisited_set[g.location, g.facing] = 0

    count = 0
    start = timer()
    while len(unvisited_set) > 0:
        count += 1
        if count % 100 == 0:
            print(f'There are {len(unvisited_set)} unvisited locations. Visited 100 nodes in {(timer() - start) * 1000:2f} ms')
            start = timer()

        # 3. Select the node with the smallest distance.
        current_node = min(unvisited_set, key=unvisited_set.get)
        # Remove it from the set (step 5).  We can do this because we cannot revisit the node from here
        current_distance = unvisited_set.pop(current_node)

        if current_node[0] == goal:
            if result is None or result > current_distance:
                result = current_distance

        if current_distance == math.nan:
            print('done because we could not get to any other node')
            return result


        # 4. Consider Unvisited Neighbors
        update_distance(current_node, current_distance, None, unvisited_set)
        update_distance(current_node, current_distance, Turn.Left, unvisited_set)
        update_distance(current_node, current_distance, Turn.Right, unvisited_set)

    print('done because all nodes were visited')

    end_node = None
    end_distance = 0
    for node in distances.keys():
        distance = distances[node]
        if node[0] == goal and (end_node is None or distance < end_distance):
            end_node = node
            end_distance = distance

    path = []
    while end_node is not None:
        print(f'Visited {end_node} with distance {end_distance}')
        path.append(end_node)
        try:
            end_node = prev[end_node]
            end_distance = distances[end_node]
        except KeyError:
            end_node = None

    loc = (0, 0)
    turn_count = 0
    for p in path:
        print(p)
        if p[0] == loc:
            turn_count += 1
            print(f'This is the {turn_count} at loc')
        else:
            turn_count = 0
        loc = p[0]
    print(result)

    pygame.init()
    clock = pygame.time.Clock()
    scale = 12
    screen = pygame.display.set_mode((len(g.grid[0]) * scale, len(g.grid) * scale))
    walls = pygame.Surface((len(g.grid[0]) * scale, len(g.grid) * scale))
    for row in range(len(g.grid)):
        r = g.grid[row]
        for col in range(len(r)):
            if g.grid[row][col] == MazeSpace.Wall:
                pygame.draw.rect(walls, "red", (col * scale, row * scale, scale, scale))
    walls.set_alpha(100)
    character_dir: dict[Direction, pygame.Surface] = {}
    character = pygame.Surface((scale, scale))
    pygame.draw.polygon(character, "yellow", ((0, scale), (0, scale / 2), (scale, scale)))
    character_dir[Direction.North] = character
    character_dir[Direction.East] = pygame.transform.rotate(character, 90)
    character_dir[Direction.South] = pygame.transform.rotate(character, 180)
    character_dir[Direction.West] = pygame.transform.rotate(character, -90)

    trail = pygame.Surface((len(g.grid[0]) * scale, len(g.grid) * scale))
    trail.set_alpha(100)

    current_loc = g.location
    current_dir = g.facing
    screen.blit(character_dir[current_dir], (current_loc[1] * scale, current_loc[0] * scale))

    running = True

    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False

        if len(path) > 0:
            screen.fill("black")

            ## Add the trail
            pygame.draw.rect(trail, "green", (current_loc[1] * scale, current_loc[0] * scale, scale, scale))
            screen.blit(trail, (0,0))

            ## Add the character
            current_loc, current_dir = path.pop()
            screen.blit(character_dir[current_dir], (current_loc[1] * scale, current_loc[0] * scale))

            # Add the walls
            screen.blit(walls, (0, 0))

        pygame.display.flip()
        clock.tick(10)

    return result

def part1(g: MazeGrid):
    graph: list[GraphNode] = []
    for space in g.spaces:
        graph.append((space, Direction.North))
        graph.append((space, Direction.South))
        graph.append((space, Direction.East))
        graph.append((space, Direction.West))

    def get_neighbors(n: GraphNode) -> list[tuple[GraphNode, int]]:
        neighbors = []
        next_node = (get_next_location(n[0], n[1]), n[1])
        neighbors.append((next_node, 1))
        next_node = (n[0], n[1].turn(Turn.Left))
        neighbors.append((next_node, 1000))
        next_node = (n[0], n[1].turn(Turn.Right))
        neighbors.append((next_node, 1000))
        return neighbors


    dist, prev = dijkstra(graph, (g.location, g.facing), get_neighbors)
    result = math.inf
    end_node = None
    for d in Direction:
        c = dist.get((g.goal, d), math.inf)
        if c < result:
            end_node = (g.goal, d)
            result = c


    v = set()
    v.add(g.location)

    def count_visited(nodes: list[GraphNode], visited: set[Location], previous: dict[GraphNode, list[GraphNode]]):
        to_visit = [*nodes]
        while len(to_visit) > 0:
            node = to_visit.pop(0)
            v.add(node[0])
            if node in previous:
                next_nodes = previous[node]
                to_visit.extend(previous[node])

    count_visited([end_node], v, prev)

    return result, len(v)


if __name__ == '__main__':
    g = MazeGrid("data")
    print(part1(g))