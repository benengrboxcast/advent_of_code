from operator import truediv
from typing import TypedDict
from enum import Enum
from timeit import default_timer as timer
import pygame

class WarehouseSpace(Enum):
    Empty = 0
    Wall = 1
    Barrel = 2
    BarrelLeft = 3
    BarrelRight = 4

class Movement(Enum):
    Up = 0
    Down = 1
    Left = 2
    Right = 3

type Warehouse = list[list[WarehouseSpace]]
type Location = tuple[int, int]

class ParsedInput(TypedDict):
    warehouse: Warehouse
    robot_location: Location
    move_list: list[Movement]

def parse_input(filename) -> ParsedInput:
    result: ParsedInput = {
        "warehouse": [],
        'robot_location': (0, 0),
        'move_list': [],
    }

    parsing_moves = False
    row = 0
    with open(filename) as f:
        for line in f:
            line = line.strip()
            if len(line) == 0:
                parsing_moves = True
            elif parsing_moves:
                col = 0
                for c in line:
                    if c == '^':
                        move = Movement.Up
                    elif c== 'v':
                        move = Movement.Down
                    elif c == '<':
                        move = Movement.Left
                    elif c == '>':
                        move = Movement.Right
                    else:
                        raise RuntimeError(f'Failed to parse move {c} at position ({row}, {col})')
                    result['move_list'].append(move)
                    col += 1
                pass
            else:
                col = 0
                warehouse_row: list[WarehouseSpace] = []
                for c in line:
                    if c == '#':
                        warehouse_row.append(WarehouseSpace.Wall)
                    elif c == '.':
                        warehouse_row.append(WarehouseSpace.Empty)
                    elif c == '@':
                        result['robot_location'] = (row, col)
                        warehouse_row.append(WarehouseSpace.Empty)
                    elif c == 'O':
                        warehouse_row.append(WarehouseSpace.Barrel)
                    else:
                        raise RuntimeError(f'Failed to parse space {c} at position ({row}, {col})')
                    col += 1
                result['warehouse'].append(warehouse_row)
            row += 1
        pass
    return result


def next_location(location: Location, direction: Movement) -> Location:
    if direction == Movement.Up:
        return location[0] - 1, location[1]
    elif direction == Movement.Down:
        return location[0] + 1, location[1]
    elif direction == Movement.Left:
        return location[0], location[1] - 1
    elif direction == Movement.Right:
        return location[0], location[1] + 1

def reverse_direction(direction: Movement) -> Movement:
    if direction == Movement.Up:
        return Movement.Down
    elif direction == Movement.Down:
        return Movement.Up
    elif direction == Movement.Left:
        return Movement.Right
    elif direction == Movement.Right:
        return Movement.Left

def print_warehouse(warehouse: Warehouse, location: Location) -> None:
    for row in range(len(warehouse)):
        warehouse_row = warehouse[row]
        warehouse_str = ""
        for col in range(len(warehouse_row)):
            if row == location[0] and col == location[1]:
                warehouse_str += "@"
            else:
                space = warehouse[row][col]
                if space == WarehouseSpace.Empty:
                    warehouse_str += "."
                elif space == WarehouseSpace.Wall:
                    warehouse_str += "#"
                elif space == WarehouseSpace.Barrel:
                    warehouse_str += "O"
        print(warehouse_str)
    pass

def get_space(warehouse: Warehouse, location: Location) -> WarehouseSpace:
    return warehouse[location[0]][location[1]]

def perform_movement(warehouse: Warehouse, location: Location, direction: Movement) -> Location:
    space: WarehouseSpace = warehouse[location[0]][location[1]]
    if space != WarehouseSpace.Empty:
        print_warehouse(warehouse, location)
        raise RuntimeError(f'Tried to perform movement when robot space was not empty {location}')

    offset = 1
    current_location = next_location(location, direction)

    space = get_space(warehouse, current_location)
    while (space != WarehouseSpace.Wall) and (space != WarehouseSpace.Empty):
        offset += 1
        current_location = next_location(current_location, direction)
        space = get_space(warehouse, current_location)
    if space == WarehouseSpace.Empty:
        rev_dir = reverse_direction(direction)
        loc = next_location(current_location, rev_dir)
        for i in range(offset):
            space = warehouse[loc[0]][loc[1]]
            warehouse[current_location[0]][current_location[1]] = space
            current_location = loc
            loc = next_location(current_location, rev_dir)
        location = next_location(location, direction)
    return location

def hash_warehouse(warehouse: Warehouse) -> int:
    result = 0
    for row in range(1, len(warehouse)):
        for col in range(1, len(warehouse[row])):
            space = warehouse[row][col]
            if space == WarehouseSpace.Barrel:
                result += row * 100 + col
    return result

def part1(data: ParsedInput) -> int:
    warehouse = data['warehouse']
    location = data['robot_location']
    for move in data['move_list']:
        location = perform_movement(warehouse, location, move)
    return hash_warehouse(warehouse)

type BoxLocations = set[Location]
type WallLocations = set[Location]

def handle_vertical_push(boxes: BoxLocations, warehouse: Warehouse, start_location: Location, direction: Movement):
    affected_boxes: list[set[Location]] = []
    row_delta = 1
    if direction == Movement.Up:
        row_delta = -1

    current_row = start_location[0] + row_delta
    boxes_this_row: list[Location] = []
    wall_found = False
    for col in range(start_location[1] - 1, start_location[1] + 1):
        wall_range = start_location[1] <= col < start_location[1] + 1
        if wall_range and warehouse[current_row][col] == WarehouseSpace.Wall:
            wall_found = True
            break
        elif (current_row, col) in boxes:
            boxes_this_row.append((current_row, col))

    while not wall_found and len(boxes_this_row) > 0:
        affected_boxes.insert(0, boxes_this_row)
        boxes_this_row: set[Location] = set()
        current_row += row_delta
        for box in affected_boxes[0]:
            for col in range(box[1] - 1, box[1] + 2):
                wall_range = box[1] <= col < start_location[1] + 2
                if wall_range and warehouse[current_row][col] == WarehouseSpace.Wall:
                    wall_found = True
                    break
                if (current_row, col) in boxes:
                    boxes_this_row.add((current_row, col))

    if not wall_found:
        for row in affected_boxes:
            for box in row:
                boxes.remove(box)
                boxes.add((box[0] + row_delta, box[1]))
        return next_location(start_location, direction)
    return start_location

def handle_horizontal_push(boxes: BoxLocations, warehouse: Warehouse, start_location: Location, direction: Movement):
    col_delta = 2
    col_offset = 1
    if direction == Movement.Left:
        col_delta = -2
        col_offset = -2
        if warehouse[start_location[0]][start_location[1] - 1] == WarehouseSpace.Empty and warehouse[start_location[0]][start_location[1] - 2] == WarehouseSpace.Wall:
            return next_location(start_location, direction)

    affected_boxes: list[Location] = []
    row = start_location[0]
    col = start_location[1]
    while (row, col + col_offset) in boxes:
        affected_boxes.insert(0, (row, col + col_offset))
        col_offset += col_delta

    if direction == Movement.Left:
        wall_found = warehouse[row][col + col_offset + 1] == WarehouseSpace.Wall
    else:
        wall_found = warehouse[row][col + col_offset] == WarehouseSpace.Wall

    if not wall_found:
        col_delta /= 2
        col_offset = col_delta
        for box in affected_boxes:
            boxes.remove(box)
            boxes.add((box[0], box[1] + col_offset))
        return next_location(start_location, direction)
    return start_location


def expand_warehouse(warehouse: Warehouse) -> (Warehouse, set[Location]):
    result: Warehouse = []
    boxes: set[Location] = set()
    for row_index, row in enumerate(warehouse):
        new_row: list[WarehouseSpace] = []
        for col_index, col in enumerate(row):
            if col == WarehouseSpace.Wall:
                new_row.append(WarehouseSpace.Wall)
                new_row.append(WarehouseSpace.Wall)
            else:
                new_row.append(WarehouseSpace.Empty)
                new_row.append(WarehouseSpace.Empty)
            if col == WarehouseSpace.Barrel:
                boxes.add((row_index, col_index * 2))
        result.append(new_row)
    return result, boxes

def print_expanded_warehouse(warehouse: Warehouse, boxes: set[Location], robot: Location):
    for i in range(len(warehouse)):
        s = ""
        for j in range(len(warehouse[i])):
            if (i, j) in boxes:
              s += '['
            elif (i, j - 1) in boxes:
                s += ']'
            elif (i, j) == robot:
                s += '@'
            elif warehouse[i][j] == WarehouseSpace.Wall:
                s += '#'
            elif warehouse[i][j] == WarehouseSpace.Empty:
                s += '.'
            else:
                raise RuntimeError(f'Unhandled location: {i}, {j}, {warehouse[i][j]}')
        print(s)

def animate_part_2(data: ParsedInput, scale: int = 60):
    pygame.init()
    expanded = expand_warehouse(data['warehouse'])
    warehouse: Warehouse = expanded[0]
    boxes: set[Location] = expanded[1]
    location = data['robot_location'][0], data['robot_location'][1] * 2
    x_size = len(warehouse[0]) * scale
    y_size = len(warehouse) * scale
    screen = pygame.display.set_mode((x_size, y_size + 200))
    clock = pygame.time.Clock()
    running = True
    frame_limit = 20
    move_number = -1
    GAME_FONT = pygame.font.SysFont('Arial', 12)
    text_pos = (5, len(warehouse) * scale + 45)
    pause_update = True
    free_run = False
    wall_surface = pygame.surface.Surface((len(warehouse[0]) * scale, len(warehouse) * scale))
    barrel_surface = pygame.surface.Surface((len(warehouse) * scale, len(warehouse) * scale))

    while move_number < 1960:
        move_number += 1
        move = data["move_list"][move_number]
        if move == Movement.Left or move == Movement.Right:
            location = handle_horizontal_push(boxes, warehouse, location, move)
        else:
            location = handle_vertical_push(boxes, warehouse, location, move)

    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            elif event.type == pygame.KEYDOWN:
                if event.key == pygame.K_RIGHT:
                    pause_update = False
                if event.key == pygame.K_f:
                    free_run = not free_run

        move = None
        if not pause_update:
            if not free_run:
                pause_update = True
            move_number += 1
            if 0 <= move_number <= len(data["move_list"]) - 1:
                move = data["move_list"][move_number]
            else:
                move = None

        if move is not None:
            if move == Movement.Left or move == Movement.Right:
                location = handle_horizontal_push(boxes, warehouse, location, move)
            else:
                location = handle_vertical_push(boxes, warehouse, location, move)

        screen.fill("black")
        # Fill the wall surface
        for row in range(len(warehouse)):
            for col in range(len(warehouse[row])):
                if warehouse[row][col] == WarehouseSpace.Wall:
                    wall_x = col * scale
                    wall_y = row * scale
                    pygame.draw.rect(screen, "red", [wall_x + 1, wall_y + 1, scale - 2, scale -2])

        for box in boxes:
            barrel_x = box[1] * scale
            barrel_y = box[0] * scale
            pygame.draw.rect(screen, "blue", [barrel_x + 1, barrel_y + 1, scale * 2 -2, scale - 2])

        location_x = location[1] * scale + scale / 2
        location_y = location[0] * scale + scale / 2
        pygame.draw.circle(screen, "green", (location_x, location_y), scale / 2)

        if move_number < 0:
            text = "initial position"
        elif move_number > len(data["move_list"]) - 1:
            text = "final position"
        else:
            text = f'Move Number: {move_number} - {data["move_list"][move_number]} of {len(data["move_list"])}'
        text_surface = GAME_FONT.render(text, False, (255, 255, 255))
        screen.blit(text_surface, text_pos)

        pygame.display.flip()
        current_frame = 0


        current_frame += 1
        clock.tick(frame_limit)



def part2(data: ParsedInput) -> int:
    warehouse, boxes = expand_warehouse(data['warehouse'])
    location = data['robot_location'][0], data['robot_location'][1] * 2

    print("Initial Setup")
    print_expanded_warehouse(warehouse, boxes, location)
    for idx, move in enumerate(data['move_list']):
        # print(f'\nGoing to move {idx}: {move}')
        # print_expanded_warehouse(warehouse, boxes, location)
        if move == Movement.Up or move == Movement.Down:
            location = handle_vertical_push(boxes, warehouse, location, move)
        else:
            location = handle_horizontal_push(boxes, warehouse, location, move)

    result = 0
    for box in boxes:
        result += box[0] * 100 + box[1]
    return result

if __name__ == '__main__':
    filename = "data"

    start = timer()
    parsed_input = parse_input(filename)
    print(f'It tool {(timer() - start) * 1000:.2f} ms to parse the input')
    # start = timer()
    # r = part1(parsed_input)
    # print(f'Part 1: {r} took {(timer() - start) * 1000:.2f} ms')
    animate_part_2(parsed_input.copy(), 30)

    start = timer()
    r = part2(parsed_input)
    print(f'Part 2: {r} took {(timer() - start) * 1000:.2f} ms')