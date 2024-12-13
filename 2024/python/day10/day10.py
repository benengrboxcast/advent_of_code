from timeit import default_timer as timer
from typing import TypedDict


type MapLocation = tuple[int, int]

class MapNode(TypedDict):
    value: int
    destinations: set[MapLocation]
    trailhead_count: int
    parsed: bool

type RawMap = dict[MapLocation, MapNode]

class InputData(TypedDict):
    graph: RawMap
    rows: int
    cols:  int

def parse_input(filename) -> InputData:
    results: RawMap = {}
    row_idx = 0
    with open(filename) as f:

        for line in f:
            line = line.strip()
            col_idx = 0
            for c in  line:
                val = int(c)
                locs = set()
                parsed = False
                trailhead_count = 0
                if val == 9:
                    locs.add((row_idx, col_idx))
                    parsed = True
                    trailhead_count = 1
                n: MapNode = {
                    'value': val,
                    'destinations': locs,
                    'parsed': parsed,
                    'trailhead_count': trailhead_count,
                }
                results[(row_idx, col_idx)] = n

                col_idx += 1
            row_idx += 1
        pass
    return {
        'graph': results,
        'rows': row_idx,
        'cols': col_idx
    }

def parse_map(loc: MapLocation, data: InputData) -> set[MapLocation]:
    # Have we already parsed this location
    graph = data['graph']
    node = graph[loc]
    value = node['value']
    dests = node['destinations']

    # If we have already parsed this node, just return the values
    if node['parsed']:
        return dests

    #If this is a 9, we can't go anywhere else.
    if value == 9:
        node['destinations'].add(loc)
        node['trailhead_count'] = 1
        return node['destinations']

    # Up
    count = 0
    next_node = (loc[0] - 1, loc[1])
    if loc[0] > 0 and graph[next_node]['value'] == value + 1:
        dests.update(parse_map((loc[0] - 1, loc[1]), data))
        count += graph[next_node]['trailhead_count']

    # Down
    next_node = (loc[0] + 1, loc[1])
    if loc[0] < data['rows'] - 1 and graph[next_node]['value'] == value + 1:
        dests.update(parse_map((loc[0] + 1, loc[1]), data))
        count += graph[next_node]['trailhead_count']

    # Left
    next_node = (loc[0], loc[1] - 1)
    if loc[1] > 0 and graph[next_node]['value'] == value + 1:
        dests.update(parse_map((loc[0], loc[1] - 1), data))
        count += graph[next_node]['trailhead_count']

    # Right
    next_node = (loc[0], loc[1] + 1)
    if loc[1] < data['cols'] - 1 and graph[(loc[0], loc[1] + 1)]['value'] == value + 1:
        dests.update(parse_map((loc[0], loc[1] + 1), data))
        count += graph[next_node]['trailhead_count']

    node['parsed'] = True
    node['trailhead_count'] = count
    return dests


def part1(data: InputData) -> (int, int):
    score = 0
    count = 0
    graph = data['graph']
    for location in graph.keys():
        if data['graph'][location]['value'] == 0:
            score += len(parse_map(location, data))
            count += graph[location]['trailhead_count']
    return score, count


if __name__ == "__main__":
    start = timer()
    input_data = parse_input("example")
    print(f'Parsing input took {timer() - start:.3f}s')
    start = timer()
    s, c = part1(input_data)
    print(f'Part 1 took {timer() - start:.3f}s: {s} / {c}')