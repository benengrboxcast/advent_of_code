from timeit import default_timer as timer
from typing import TypedDict
from enum import Enum
from sortedcontainers import SortedList

type ParsedData = dict[tuple[int, int], str]

class GardenPlot(TypedDict):
    plant: str
    nodes: SortedList[tuple[int, int]]
    perimeter: int
    internal_plots: list
    edges: int
    id: int

class Direction(Enum):
    Up = 0
    Right = 1
    Down = 2
    Left = 3

type NodeMap = dict[tuple[int, int], int]

def get_next_node(current: tuple[int, int], direction: Direction) -> tuple[int, int]:
    if direction == Direction.Up:
        return current[0] - 1, current[1]
    if direction == Direction.Down:
        return current[0] + 1, current[1]
    if direction == Direction.Left:
        return current[0], current[1] - 1
    if direction == Direction.Right:
        return current[0], current[1] + 1

def plot_area(p: GardenPlot):
        return len(p['nodes'])

type GardenPlots = list[GardenPlot]

def parse_input(filename: str) -> ParsedData:
    parsed_data: ParsedData = {}
    with open(filename) as f:
        row = 0
        for line in f:
            line = line.strip()
            col = 0
            for c in line:
                parsed_data[(row, col)] = c
                col += 1
            row += 1
    return parsed_data

def add_to_plot(current_node: tuple[int, int], parsed_data: ParsedData, plot: GardenPlot):
    found_nodes = []
    for d in Direction:
        next_node = get_next_node(current_node, d)
        if next_node in plot['nodes']:
            continue
        if next_node in parsed_data and parsed_data[next_node] == plot['plant']:
            parsed_data.pop(next_node)
            plot['nodes'].add(next_node)
            found_nodes.append(next_node)
        else:
            plot['perimeter'] += 1

    for node in found_nodes:
        add_to_plot(node, parsed_data, plot)

def create_plot(parsed_data: ParsedData, plot_id=0) -> GardenPlot:
    node = next(iter(parsed_data))
    plant = parsed_data.pop(node)
    plot: GardenPlot = {
        'plant': plant,
        'nodes': SortedList(),
        'perimeter': 0,
        'id': plot_id,
        'edges': 0,
        'internal_plots': []
    }
    plot['nodes'].add(node)
    add_to_plot(node, parsed_data, plot)
    return plot

def part1(parsed_data: ParsedData) -> int:
    plots: list[GardenPlot] = []
    while len(parsed_data) > 0:
        plots.append(create_plot(parsed_data))
    total_cost = 0
    for plot in plots:
        area = plot_area(plot)
        perimeter = plot['perimeter']
        cost = perimeter * area
        total_cost += cost
    return total_cost



def get_left(current: tuple[int, int], direction: Direction) -> tuple[int, int]:
    if direction == Direction.Up:
        return current[0], current[1] - 1
    elif direction == Direction.Right:
        return current[0] - 1, current[1]
    elif direction == Direction.Down:
        return current[0], current[1] + 1
    else:
        return current[0] + 1, current[1]

def get_straight(current: tuple[int, int], direction: Direction) -> tuple[int, int]:
    if direction == Direction.Up:
        return current[0] - 1, current[1]
    elif direction == Direction.Right:
        return current[0], current[1] + 1
    elif direction == Direction.Down:
        return current[0] + 1, current[1]
    else:
        return current[0], current[1] -1

def get_right(current: tuple[int, int], direction: Direction) -> tuple[int, int]:
    if direction == Direction.Up:
        return current[0], current[1] + 1
    elif direction == Direction.Right:
        return current[0] + 1, current[1]
    elif direction == Direction.Down:
        return current[0], current[1] - 1
    else:
        return current[0] - 1, current[1]

def get_reverse(current: tuple[int, int], direction: Direction) -> tuple[int, int]:
    if direction == Direction.Up:
        return current[0] + 1, current[1]
    elif direction == Direction.Right:
        return current[0], current[1] - 1
    elif direction == Direction.Down:
        return current[0] - 1, current[1]
    else:
        return current[0], current[1] + 1

def count_edges(plot: GardenPlot, parsed_input: ParsedData, node_map: NodeMap):
    start_node = plot['nodes'][0]
    direction = Direction.Right
    plant = plot['plant']
    visited: set[tuple[int, int]] = set()
    visited.add(start_node)
    current = start_node
    edges = 1
    container_id = node_map.get(get_left(current, direction), None)

    if len(plot['nodes']) == 1:
        if not (node_map.get(get_straight(current, direction), None) == container_id and
                node_map.get(get_right(current, direction), None) == container_id and
                node_map.get(get_reverse(current, direction), None) == container_id):
            container_id = None
        plot['edges'] = 4
        return container_id

    done = False
    south = (start_node[0] + 1, start_node[1])
    need_south = south in plot['nodes']
    if not need_south:
        edges += 1

    while not done:
        if parsed_input.get(get_left(current, direction), '.') == plant:
            current = get_left(current, direction)
            next_dir = direction.value - 1
            edges += 1
        elif parsed_input.get(get_straight(current, direction), '.') == plant:
            current = get_straight(current, direction)
            next_dir = direction.value
        elif parsed_input.get(get_right(current, direction), '.') == plant:
            current = get_right(current, direction)
            next_dir = direction.value + 1
            edges += 1
        elif parsed_input.get(get_reverse(current, direction), '.') == plant:
            current = get_reverse(current, direction)
            next_dir = direction.value + 2
            edges += 2
            if node_map.get(get_straight(current, direction), None) != container_id:
                container_id = None
        else:
            # This should only happen with the node is size 1.
            print("THIS SHOULD NOT HAPPEN")
            if not (node_map.get(get_straight(current, direction), None) == container_id and
                node_map.get(get_right(current, direction), None) == container_id and
                node_map.get(get_reverse(current, direction), None) == container_id):
                container_id = None
            plot['edges'] = 4
            return container_id

        if next_dir < 0:
            next_dir += 4
        elif next_dir > 3:
            next_dir -= 4
        direction = Direction(next_dir)

        visited.add(current)

        if node_map.get(get_left(current, direction), None) != container_id:
            container_id = None

        if current == start_node:
            if need_south:
                done = south in visited
            else:
                done = True
        pass


    plot['edges'] = edges
    return container_id

def part2(parsed_data: ParsedData) -> int:
    plots: list[GardenPlot] = []
    while len(parsed_data) > 0:
        plots.append(create_plot(parsed_data))
    nodes_to_plots: dict[tuple[int, int], int] = {}
    plot_index = 0
    garden_data = parsed_data.copy()
    while len(garden_data) > 0:
        plot = create_plot(garden_data, plot_index)
        plots.append(plot)
        for n in plot['nodes']:
            nodes_to_plots[n] = plot_index
        plot_index += 1

    for plot in plots:
        container_id = count_edges(plot, parsed_data, nodes_to_plots)
        if container_id is not None:
            plots[container_id]['internal_plots'].append(plot)

    total_cost = 0
    for plot in plots:
        total_edges = plot['edges']
        for p in plot['internal_plots']:
            total_edges += p['edges']
        area = plot_area(plot)
        cost = total_edges * area
        print(f'Plot {plot["id"]} ({plot["plant"]}): {total_edges} edges ({plot["edges"]} outside) * {area} area = ${cost}')
        if len(plot['internal_plots']) > 0:
            print("Internal Plots:")
            for p in plot['internal_plots']:
                print(f'\tPlot {p["id"]}: {p["plant"]}')

        total_cost += cost
    return total_cost

def part2_after_hint(parsed_data: ParsedData) -> int:
    """ Part 2 done after getting two hints

    1. The number of sides is the same as the number of corners
    2. We DON'T actually have to determine if a plot is inside another one,
       we just need to know if there is a corner.

    These two things together mean that for each unique plant (not area) we
    need to count the number of nodes that have that plant and count the
    number of nodes that are corners

    |-
    :param data:
    :return:
    """
    plot_data = parsed_data.copy()
    plots: list[GardenPlot] = []
    while len(plot_data) > 0:
        plots.append(create_plot(plot_data))

    for plot in plots:
        plot['edges'] = 0
        value = plot['plant']
        corners = 0
        for node in plot['nodes']:
            # Top left
            # xx  0x  X0
            # x0  x0  00
            nw = parsed_data.get((node[0] - 1, node[1] - 1), None)
            n = parsed_data.get((node[0] - 1, node[1]), None)
            ne = parsed_data.get((node[0] - 1, node[1] + 1), None)
            w = parsed_data.get((node[0], node[1] - 1), None)
            e = parsed_data.get((node[0], node[1] + 1), None)
            sw = parsed_data.get((node[0] + 1, node[1] - 1), None)
            s = parsed_data.get((node[0] + 1, node[1]), None)
            se = parsed_data.get((node[0] + 1, node[1] + 1), None)
            if n != value and w != value:
                corners += 1
            elif nw != value:
                if n == value and w == value:
                    corners += 1
            if n != value and e != value:
                corners += 1
            elif ne != value:
                if n == value and e == value:
                    corners += 1
            if s != value and w != value:
                corners += 1
            elif sw != value:
                if w == value and s == value:
                    corners += 1
            if s != value and e != value:
                corners += 1
            elif se != value:
                if s == value and e == value:
                    corners += 1
        plot['edges'] = corners

    total_cost = 0
    for plot in plots:
        sides = plot['edges']
        ares = len(plot['nodes'])
        cost = sides * ares
        print(f'{plot["plant"]}: sides ({sides}), areas ({ares}), cost ({cost})')
        total_cost += cost
    return total_cost




if __name__ == '__main__':
    data = parse_input('example')
    start = timer()
    print(f'{part1(data.copy())} in {(timer() - start) * 1000:.2f}ms')



    # for f in ['ex2', 'ex3', 'ex4', 'ex6']:
    #     data = parse_input(f)
    #     print(f'{f}: {part2_after_hint(data.copy())}')
    data = parse_input('data')
    print(part2_after_hint(data))

    data = parse_input('data')
    print(part2(data))