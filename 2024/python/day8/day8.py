from timeit import default_timer as timer
from typing import TypedDict

type InputMap = dict[str, list[tuple[int, int]]]
class InputData(TypedDict):
    grid: InputMap
    rows: int
    cols: int
type Part1Results = dict[tuple[int, int], list[str]]

def parse_input(filename: str) -> InputData:
    input_data: InputMap = {}
    row = 0
    with open(filename) as f:
        for line in f:
            col = 0
            line = line.strip()
            for c in line:
                if c != '.':
                    if c in input_data.keys():
                        input_data[c].append((row, col))
                    else:
                        input_data[c] = [(row,col)]
                col += 1
            row += 1
    return {
        'grid': input_data,
        'rows': row,
        'cols': col,
    }


def print_results(in_data: InputData, results: Part1Results):
    print()
    for row in range(in_data['rows']):
        line = ''
        for col in range(in_data['cols']):
            node = (row, col)
            c = '.'
            if node in results:
                c = '#'
            for key in in_data['grid']:
                if node in in_data['grid'][key]:
                    c = key
            line += c
        print(line)


def add_node(node: tuple[int, int], node_list: Part1Results, key: str, rows: int, cols: int) -> None:
    if node in node_list:
        node_list[node].append(key)
    else:
        if not (node[0] < 0 or node[0] >= rows or node[1] < 0 or node[1] >= cols):
            node_list[node] = [key]


def part1(data: InputData) -> int:
    antinodes: dict[tuple[int, int], list[str]] = {}
    for key in data['grid'].keys():
        antennas = data['grid'][key]
        for idx_1 in range(len(antennas)):
            for idx_2 in range(idx_1 + 1, len(antennas)):
                a1 = antennas[idx_1]
                a2 = antennas[idx_2]
                add_node((a1[0] + (a1[0] - a2[0]), a1[1] + (a1[1] - a2[1])), antinodes, key, data['rows'], data['cols'])
                add_node((a2[0] + (a2[0] - a1[0]), a2[1] + (a2[1] - a1[1])), antinodes, key, data['rows'], data['cols'])
    return len(antinodes)


def count_antennas_and_nodes(in_data: InputData, results: Part1Results) -> int:
    count = 0
    for row in range(in_data['rows']):
        line = ''
        for col in range(in_data['cols']):
            node = (row, col)
            if node in results:
                count += 1
            else:
                for key in in_data['grid']:
                    if node in in_data['grid'][key]:
                        count += 1
                        break
    return count

def part2(data: InputData) -> int:
    antinodes: dict[tuple[int, int], list[str]] = {}
    for key in data['grid'].keys():
        antennas = data['grid'][key]
        for idx_1 in range(len(antennas)):
            for idx_2 in range(idx_1 + 1, len(antennas)):
                a1 = antennas[idx_1]
                a2 = antennas[idx_2]
                delta_x = (a1[0] - a2[0])
                delta_y = (a1[1] - a2[1])
                x = a1[0] + delta_x
                y = a1[1] + delta_y
                while not (x < 0 or x >= data['rows'] or y < 0 or y >= data['cols']):
                    add_node((x, y), antinodes, key, data['rows'], data['cols'])
                    x += delta_x
                    y += delta_y

                x = a2[0] - delta_x
                y = a2[1] - delta_y
                while not (x < 0 or x >= data['rows'] or y < 0 or y >= data['cols']):
                    add_node((x, y), antinodes, key, data['rows'], data['cols'])
                    x -= delta_x
                    y -= delta_y
    return count_antennas_and_nodes(data, antinodes)

if __name__ == "__main__":
    start = timer()
    d = parse_input("data")
    print(f'Parsing took {timer() - start:.3f}s')

    start = timer()
    print(f'Part 1 {part1(d)} took {timer() - start:.3f}s')

    start = timer()
    print(f'Part 2 {part2(d)} took {timer() - start:.3f}s')

