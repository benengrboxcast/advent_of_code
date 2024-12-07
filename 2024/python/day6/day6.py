import copy
from enum import Enum
from typing import TypedDict
from sortedcontainers import SortedList

class Direction(Enum):
    NORTH = 0
    EAST = 1
    SOUTH = 2
    WEST = 3

class GameState(TypedDict):
    layout: list[list[str]]
    position: tuple[int, int]
    direction: Direction

def parse_input(filename) -> GameState:
    layout = []
    position = (0, 0)
    with open(filename) as f:
        row = 0
        for line in f:
            col = line.find('^')
            if col >= 0:
                position = (row, col)
                line.replace('^', '.')
            layout.append(list(line.strip()))
            row += 1
    return {
        'layout': layout,
        'position': position,
        'direction': Direction.NORTH,
    }


def get_guard_char(dir: Direction) -> str:
    if dir == Direction.NORTH:
        return '^'
    elif dir == Direction.EAST:
        return '>'
    elif dir == Direction.SOUTH:
        return 'V'
    else:
        return '<'

def get_next_directon(dir: Direction) -> Direction:
    if dir == Direction.NORTH:
        return Direction.EAST
    elif dir == Direction.EAST:
        return Direction.SOUTH
    elif dir == Direction.SOUTH:
        return Direction.WEST
    else:
        return Direction.NORTH

def print_board(state: GameState) -> None:
    row = 0
    for line in state['layout']:
        if row == state['position'][0]:
            line[state['position'][1]] = get_guard_char(state['direction'])
        print(line)
        row += 1
    pass

def do_movement(state: GameState) -> GameState:
    row = state['position'][0]
    col = state['position'][1]

    state['layout'][row][col] = 'X'
    if state['direction'] == Direction.NORTH:
        while row > 0 and state['layout'][row - 1][col] != '#':
            state['layout'][row][col] = 'X'
            row -= 1
        if row == 0:
            state['layout'][0][col] = 'X'
            row = -1
    elif state['direction'] == Direction.EAST:
        while col < len(state['layout'][row]) - 1 and state['layout'][row][col + 1] != '#':
            state['layout'][row][col] = 'X'
            col += 1
        if col == len(state['layout'][row]) - 1:
            state['layout'][row][len(state['layout'][row]) - 1] = 'X'
            col = -1
    elif state['direction'] == Direction.SOUTH:
        while row < len(state['layout']) - 1 and state['layout'][row + 1][col] != '#':
            state['layout'][row][col] = 'X'
            row += 1
        if row == len(state['layout']) - 1:
            state['layout'][len(state['layout']) - 1][col] = 'X'
            row = -1
    else:
        while col > 0 and state['layout'][row][col - 1] != '#':
            state['layout'][row][col] = 'X'
            col -= 1
        if col == 0:
            state['layout'][row][0] = 'X'
            col = -1
    return {
        'layout': state['layout'],
        'position': (row, col),
        'direction': get_next_directon(state['direction']),
    }

def part1(initial_state: GameState) -> int:
    state = copy.deepcopy(initial_state)
    while state['position'][0] >= 0 and state['position'][1] >= 0:
        state = do_movement(state)
    result = 0
    for line in state['layout']:
        result += line.count('X')
    return result

def hash_state(state: GameState) -> int:
    return state['position'][0] * 10000 + state['position'][1] * 10 + state['direction'].value


def part2(state: GameState) -> int:
    result = 0
    for row in range(len(state['layout'])):
        for col in range(len(state['layout'][row])):
            current_state = copy.deepcopy(state)
            if current_state['layout'][row][col] == '#':
                continue
            elif current_state['layout'][row][col] == '^':
                continue
            current_state['layout'][row][col] = '#'
            found_states = SortedList()
            while current_state['position'][0] >= 0 and current_state['position'][1] >= 0:
                current_state = do_movement(current_state)
                h = hash_state(current_state)
                if found_states.count(h) > 0:
                    result += 1
                    break
                else:
                    found_states.add(h)
            pass
        print(f'Testing Row {row}: {result}')
    return result

if __name__ == '__main__':
    s = parse_input('data')
    print(part1(s))
    print(part2(s))