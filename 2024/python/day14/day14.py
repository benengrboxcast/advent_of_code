import re
from timeit import default_timer as timer
import pygame
import math

type Position = tuple[int, int]
type Velocity = tuple[int, int]
type RobotInfo = tuple[Position, Velocity]
type ParsedInput = list[RobotInfo]

line_regex = re.compile(r"p=(\d+),(\d+) v=(-?\d+),(-?\d+)")
def parse_input(filename) -> ParsedInput:
    parsed_input: ParsedInput = []
    with open(filename) as f:
        for line in f:
            found = line_regex.search(line)
            position = (int(found.group(1)), int(found.group(2)))
            velocity = (int(found.group(3)), int(found.group(4)))
            parsed_input.append((position, velocity))
    return parsed_input

def update_position(info: RobotInfo, secs: int, x_limit: int, y_limit: int) -> RobotInfo:
    x = info[0][0] + info[1][0] * secs
    y = info[0][1] + info[1][1] * secs
    return (x % x_limit, y % y_limit), info[1]

def print_board(data: ParsedInput, x_limit: int, y_limit: int) -> None:
    board = []
    for i in range(y_limit):
        row = []
        for j in range(x_limit):
            row.append(0)
        board.append(row)

    for info in data:
        bot = info[0]
        board[bot[1]][bot[0]] += 1

    for row in board:
        row_str = ''
        for col in row:
            if col == 0:
                row_str += '.'
            else:
                row_str += str(col)
        print(row_str)

def hash_data(data: ParsedInput, x_limit, y_limit) -> int:
    x_div = x_limit // 2
    y_div = y_limit // 2

    quads = [0, 0, 0, 0]
    for info in data:
        loc = info[0]
        if loc[0] < x_div and loc[1] < y_div:
            quads[0] += 1
        if loc[0] > x_div and loc[1] < y_div:
            quads[1] += 1
        if loc[0] < x_div and loc[1] > y_div:
            quads[2] += 1
        if loc[0] > x_div and loc[1] > y_div:
            quads[3] += 1
    result = 1
    for quad in quads:
        result *= quad
    return result

def part1(data: ParsedInput, elapsed: int, x_limit: int, y_limit: int) -> int:
    for i in range(len(data)):
        data[i] = update_position(data[i], elapsed, x_limit, y_limit)
    return hash_data(data, x_limit, y_limit)

def part2(data: ParsedInput, x_limit, y_limit, scale = 4) -> int:
    pygame.init()
    screen = pygame.display.set_mode((x_limit * scale, y_limit * scale + 100))
    clock = pygame.time.Clock()
    running = True
    frame_limit = 60
    frame_delta = 6
    frame_update = 6
    current_frame = 0
    robot_pos = 0
    reverse = False
    GAME_FONT = pygame.font.SysFont('Arial', 12)
    text_pos = (5, y_limit * scale + 5)
    pause_update = False

    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            elif event.type == pygame.KEYDOWN:
                if event.key == pygame.K_r:
                    reverse = True
                elif event.key == pygame.K_f:
                    reverse = False
                elif event.key == pygame.K_UP:
                    frame_update -= frame_delta
                elif event.key == pygame.K_DOWN:
                    frame_update += frame_delta
                elif event.key == pygame.K_SPACE:
                    pause_update = not pause_update

        if not pause_update and current_frame >= frame_update:
            screen.fill("black")
            for robot in data:
                robot_x = robot[0][0]
                robot_y = robot[0][1]
                pygame.draw.rect(screen, "red", (robot_x * scale, robot_y * scale, scale, scale))
            text_surface = GAME_FONT.render(f'Current Iteration: {robot_pos}', False, (155, 155, 255))
            screen.blit(text_surface, text_pos)
            pygame.display.flip()
            current_frame = 0

        if current_frame == 1:
            amount = 1
            if reverse:
                amount = -1
            for i in range(len(data)):
                data[i] = update_position(data[i], amount, x_limit, y_limit)
            robot_pos += amount

        current_frame += 1
        clock.tick(frame_limit)

    pygame.quit()
    return current_frame

def find_loop(info: RobotInfo, x_limit: int, y_limit: int) -> int:
    x_repeat = None
    y_repeat = None
    x_start = info[0][0]
    y_start = info[0][1]
    current = info

    count = 0
    while x_repeat is None or y_repeat is None:
        current = update_position(current, 1, x_limit, y_limit)
        count += 1
        if x_repeat is None and current[0][0] == x_start:
            x_repeat = count
        if y_repeat is None and current[0][1] == y_start:
            y_repeat = count

    return math.lcm(x_repeat, y_repeat)

def find_loops(data: ParsedInput, x_limit, y_limit):
    repeats = []
    for info in data:
        repeats.append(find_loop(info, x_limit, y_limit))
    loop = math.lcm(*repeats)
    return loop

if __name__ == "__main__":
    filename = "data"
    x_limit = 101
    y_limit = 103
    start = timer()
    data = parse_input("data")
    print(f'Parsing Input: {(timer() - start) * 1000:.2f} ms')
    start = timer()
    r = part1(data.copy(), 100, 101, 103)
    print(f'Part 1: {r} in {(timer() - start) * 1000:.2f} ms')

    start = timer()
    r = find_loops(data, x_limit, y_limit)
    print(f'It took {(timer() - start) * 1000:.2f} ms to find loop {r}')
    part2(data.copy(), x_limit, y_limit)