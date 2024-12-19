from timeit import default_timer as timer
import re

type ButtonDef = tuple[int, int]
type Location = tuple[int, int]
type Machine = tuple[ButtonDef, ButtonDef, Location]
type ParsedInput = list[Machine]

def parse_input(filename: str) -> ParsedInput:
    parsed_input: ParsedInput = []
    button_a: ButtonDef = (0, 0)
    button_b: ButtonDef = (0, 0)
    prize_location: Location = (0, 0)

    button_regex = re.compile(r"X\+([0-9]*), Y\+([0-9]*)")
    location_regex = re.compile(r"X=(\d+), Y=(\d+)")
    with open(filename) as f:
        for line in f:
            if len(line) < 8:
                continue
            if line[7] == 'A':
                found = button_regex.search(line)
                button_a = (int(found.group(1)), int(found.group(2)))
                pass
            if line[7] == 'B':
                found = button_regex.search(line)
                button_b = (int(found.group(1)), int(found.group(2)))
                pass
            if line[0] == 'P':
                found = location_regex.search(line)
                prize_location = (int(found.group(1)), int(found.group(2)))
                parsed_input.append((button_a, button_b, prize_location))
                pass
    return parsed_input

def find_best_count(machine: Machine, press_limit) -> int | None:
    button_a: ButtonDef = machine[0]
    button_b: ButtonDef = machine[1]
    location: Location = machine[2]
    max_found = 4 * press_limit + 1
    found = max_found
    for a_press in range(press_limit):
        b_press = (location[0] - (button_a[0] * a_press)) // button_b[0]
        if a_press * button_a[0] + b_press * button_b[0] == location[0]:
            if a_press * button_a[1] + b_press * button_b[1] == location[1]:
                coin_count = 3 * a_press + b_press
                if coin_count < found:
                    found = coin_count
    if found == max_found:
        return None
    else:
        return found


def find_best_count_2(machine: Machine) -> int | None:
    """ Find the best count without a press limit

    The total number of coins is = 3 * a_press + b_press

    Let's only care about the x location and then if that works, verify the y location.

    The x location is button_a[0] * a_press + button_b[0] * b_press and this has to equal location[0]

    Solving for b_press
    b_press = (location[0] - (button_a[0] * a_press)) / button_b[0]

    coins = 3 * a_press + (location[0] - (button_a[0] * a_press) / button_b[0])
    coins = a_press * ( 3 * button_b[0] - button_a[0]) / button_b[0] + location[0] / button_b[0]

    This is a line.
    Slope is positive if button_b[0] < 3 * button_a[0], so we start with 0 and stop with the first match
    Slope is negative if button_b[0] > 3 * button_a[0], so we start with b = 0 and stop with the first match
    If they are equal it doesn't matter

    Simpler:
    To maximize the x distance for the least number of coins, if bx > 3 * ax, then we use b


    And the x location



    :param machine: The machine to solve
    :param offset: The offset of the locations
    :return:
    """
    button_a: ButtonDef = machine[0]
    button_b: ButtonDef = machine[1]
    location: Location = (machine[2][0] + 10000000000000, machine[2][1] + 10000000000000)

    a_limit = location[0] // button_a[0]
    a_range = range(0, a_limit)
    if button_b[0] > 3 * button_a[0]:
        a_range = reversed(a_range)

    for a_press in a_range:
        b_press = (location[0] - (button_a[0] * a_press)) // button_b[0]
        if a_press * button_a[0] + b_press * button_b[0] == location[0]:
            if a_press * button_a[1] + b_press * button_b[1] == location[1]:
                return 3 * a_press + b_press
    return None

def find_best_count_3(machine: Machine, offset: int = 10000000000000) -> int | None:
    """

    Let's use x and y to figure out the max presses for each button
    :param machine:
    :param offset:
    :return:
    """
    button_a: ButtonDef = machine[0]
    button_b: ButtonDef = machine[1]
    location: Location = (machine[2][0] + offset, machine[2][1] + offset)

    b_press = (location[1] * button_a[0] - location[0] * button_a[1]) / (button_a[0] * button_b[1] - button_a[1] * button_b[0])
    a_press = (location[0] - button_b[0] * b_press) / button_a[0]

    if int(a_press) == a_press and int(b_press) == b_press:
        return 3 * int(a_press) + int(b_press)
    return None

def part1(parsed_input: ParsedInput) -> int:
    total_count = 0
    for machine in parsed_input:
        this_count = find_best_count_3(machine, 0)
        if this_count is not None:
            total_count += this_count
    return total_count

def part2(parsed_input: ParsedInput) -> int:
    total_count = 0
    index = 0
    for machine in parsed_input:
        this_count = find_best_count_3(machine)
        if this_count is not None:
            total_count += this_count
        index += 1
    return total_count

if __name__ == "__main__":
    start = timer()
    data = parse_input("data")
    print(f'Parse Input: in {(timer() - start) * 1000:.3f} ms')
    start = timer()
    r = part1(data)
    print(f'Part 1: {r} in {(timer() - start):.3f} ms')
    start = timer()
    r = part2(data)
    print(f'Part 2: {r} in {(timer() - start):.3f} ms')