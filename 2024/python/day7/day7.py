from math import log10, floor
from typing import TypedDict
from timeit import default_timer as timer

class PuzzleInput(TypedDict):
    result: int
    numbers: list[int]

def parse_input(filename) -> list[PuzzleInput]:
    data = []
    with open(filename) as f:
        for line in f:
            l = line.split(':')
            result = int(l[0])
            numbers = []
            values = l[1].strip().split(' ')
            for v in values:
                numbers.append(int(v))
            data.append({
                'result': result,
                'numbers': numbers
            })
        pass
    return data

def concat_ints(x: int, y: int) -> int:
    # return int(str(x) + str(y))
    return x * (10 ** (int(log10(y)) + 1)) + y

def is_valid(result: int, numbers: list[int], use_concat: bool) -> bool:
    if len(numbers) == 0:
        return False

    if numbers[0] > result:
        return False

    if len(numbers) == 2:
        if numbers[0] + numbers[1] == result:
            return True
        elif numbers[0] * numbers[1] == result:
            return True
        elif use_concat:
            return concat_ints(numbers[0], numbers[1]) == result
        else:
            return False

    next_numbers = numbers[1:]
    if use_concat:
        next_numbers[0] = concat_ints(numbers[0], numbers[1])
        if is_valid(result, next_numbers, use_concat):
            return True
    next_numbers[0] = numbers[0] + numbers[1]
    if is_valid(result, next_numbers, use_concat):
        return True
    next_numbers[0] = numbers[0] * numbers[1]
    return is_valid(result, next_numbers, use_concat)

def part1(puzzle_input: list[PuzzleInput], use_concat) -> int:
    result = 0
    for line in puzzle_input:
        if is_valid(line['result'], line['numbers'], use_concat):
            result += line['result']
    return result

if __name__ == '__main__':
    start = timer()
    pi = parse_input('data')
    print(f'It took {timer() - start:.2f}s to parse input')
    start = timer()
    r = part1(pi, False)
    print(f'Part 1: {r} in {timer() - start:.2f}s')
    start = timer()
    r = part1(pi, True)
    print(f'Part 2: {r} in {timer() - start:.2f}s')
