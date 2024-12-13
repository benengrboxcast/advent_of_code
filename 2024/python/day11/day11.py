from timeit import default_timer as timer
from math import log10
from memory_profiler import  memory_usage
from functools import cache

def parse_input(filename):
    with open(filename) as f:
        return f.readline().strip().split(' ')

def stone_iterate(stone: str) -> list[str]:
    if stone == '0':
        return ['1']
    elif len(stone) % 2 == 0:
        sp = len(stone) // 2
        return [str(int(stone[:sp])), str(int(stone[sp:]))]
    else:
        i = int(stone)
        return [str(i * 2024)]

def digit_iterate(d: int) -> list[int]:
    if d == 0:
        return [1]
    digits = int(log10(d)) + 1
    if digits % 2 == 0:
        p = 10 ** (digits // 2)
        return [d // p, d % p]
    else:
        return [d * 2024]

@cache
def count_stones_with_cache(digit: int, current: int) -> int:
    if current < 1:
        return 1

    count = 0
    digits = digit_iterate(digit)
    for d in digits:
        count += count_stones_with_lut(d, current - 1)

    return count

# lut: dict[int, dict[int, int]] = {}
lut: dict[tuple[int, int], int] = {}

def count_stones_with_lut(digit: int, current: int) -> int:
    # If we are not going to iterate anymore, it's just this stone
    if current < 1:
        return 1

    # If we have already seen this stone and iteration, just return the value.
    if (digit, current) in lut:
        return lut[(digit, current)]

    # The number of stones after x iterations is the sum of the number of stones
    # cause by x-1 iterations for each of the stones created by iterating once.
    count = 0
    digits = digit_iterate(digit)
    for d in digits:
        count += count_stones_with_lut(d, current - 1)

    # Now that we have found the number of stones needed for this iteration,
    # store it in the lut.
    lut[(digit, current)] = count

    # Return the number of stones needed
    return count

def count_stones(data: list[str], current: int, limit: int) -> int:
    if len(data) == 0:
        return 0

    if current > limit:
        return len(data)

    count = 0
    if len(data) > 1:
        for d in data:
            count += count_stones([d], current, limit)
    elif len(data) == 0:
        return 0
    else:
        next_stones = stone_iterate(data[0])
        count += count_stones(next_stones, current + 1, limit)
    return count

def part1(data):
    start = timer()
    r = count_stones(data, 1, 25)
    print(f'It took {timer() - start:.3f} seconds to count the stones {r}')

def part2(data, count):
    start = timer()
    r = 0
    for d in data:
        r += count_stones_with_lut(int(d), count)
    print(f'It took {timer() - start:.3f} seconds to count the stones (part 2) {r}')

def part3(data, count):
    start = timer()
    r = 0
    for d in data:
        r += count_stones_with_cache(int(d), count)
    print(f'It took {timer() - start:.3f} seconds to count the stones (part 3) {r}')

if __name__ == '__main__':
    input_data = parse_input('data')

    part1(input_data)
    part2(input_data, 25)
    part2(input_data, 75)
    part3(input_data, 75)
    usage = memory_usage((part1, (input_data,), {}))
    print(f'Mem usage: {1}')
    print(f'Max Usage: {max(usage)}')
    print()

    part2(input_data, 25)
    usage = memory_usage((part2, (input_data,75), {}))
    print(f'Mem usage: {usage}')
    print(f'Max Usage: {max(usage)}')
    print()

    usage = memory_usage((part3, (input_data, 75), {}))
    print(f'Mem usage: {usage}')
    print(f'Max Usage: {max(usage)}')
    print()


