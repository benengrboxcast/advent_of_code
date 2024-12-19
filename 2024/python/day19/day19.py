from sortedcontainers import SortedList
from typing import TypedDict
from timeit import default_timer as timer
import functools

class PuzzleInput(TypedDict):
    available: SortedList[str]
    desired: list[str]

def parse_input(filename):
    available: SortedList[str] = SortedList(key=lambda s: (s, len(s)))
    desired: list[str] = []

    parsing_available = True
    with open(filename) as f:
        for line in f:
            if len(line.strip()) == 0:
                parsing_available = False
            elif parsing_available:
                towels = line.strip().split(',')
                for t in towels:
                    available.add(t.strip())
                pass
            else:
                desired.append(line.strip())
            pass
        pass
    return {
        "available": available,
        "desired": desired,
    }


def can_form_target(target, words):
    """
    Chat-GPT gave me this.

    We keep track to see if any combination can get each index of the target word.  If we can get to the end of the
    target we can make the word otherwise we cannot.

    NOTE: it may be possible to make the target and not a substring, so not every element that we keep track of may
    be true.

    :param target: The word we want to make
    :param words: The substring that we can use to make the work
    :return:
    """
    n = len(target)
    dp = [False] * (n + 1) # We create an array of each position (starting with an empty string which is why n+1
    dp[0] = True  # Base case: empty string can always be formed

    for i in range(n):
        if dp[i]:  # If position i is reachable, mark all further positions reachable from this position
            for word in words:
                word_len = len(word)
                if target[i:i + word_len] == word: # This substring matches from this position, so mark the end of the substring true
                    dp[i + word_len] = True # If word_len == n we can return true here, but that didn't really help performance.
    return dp[n]


def count_map(target: int, m: dict[int, list[int]]):
    if target == 0:
        return 1
    if target not in m:
        return 0
    count = 0
    for prev in m[target]:
        count += count_map(prev, m)
    return count


def ways_to_form_target(target, words):
    """
    Squishy brain adapted from the above algorithm

    Assumption: there are no repeated words in the word list.  That means that for each position, there is at MOST one
    way to reach the end.  More generally there is at MOST one way to reach index i from index j.

    Theory: The number of ways we can make a word is the sum of the number of ways we can reach then end times the
    number of ways we can get to each index that can reach the end.

    We need to make a mapping of index -> [indexes that can reach this index]

    :param target: The word we want to make
    :param words: The substring that we can use to make the work
    :return:
    """
    n = len(target)
    dp = [0] * (n + 1) # We create an array of each position (starting with an empty string which is why n+1
    dp[0] = 1  # Base case: empty string can always be formed

    for i in range(n):
        if dp[i] > 0:  # If position i is reachable, mark all further positions reachable from this position
            for word in words:
                word_len = len(word)
                if target[i:i + word_len] == word: # This substring matches from this position, so mark the end of the substring true
                    dp[i + word_len]+= dp[i] # The number of ways to get to the target index is increased by the number of ways to get to this index.
    return dp[n]


def find_match(desired: str, available: SortedList[str]) -> int:

    #start = available.bisect_left(desired[0])
    # There are no 'z's in the data, so we can do this
    #end = available.bisect_right(chr(ord(desired[0]) + 1))
    # for idx in range(0, len(available)):
    for t in available:
        # t = available[idx]
        if len(t) > len(desired):
            return 0
        elif len(t) == len(desired) and (t == desired):
            return 1
        elif len(t) < len(desired) and t == desired[:len(t)]:
            next = find_match(desired[len(t):], available)
            if next > 0:
                return 1 + next
    return 0


def part1(data: PuzzleInput) -> int:
    """
    Given a list of desired patterns and available patterns, determine which desired patterns can be made by any
    arrangement of the available patterns.

    :param data: The puzzle input
    :return: The number of patterns that can be made
    """
    count = 0
    for towel in data['desired']:
        this_count = can_form_target(towel, data['available'])
        if this_count:
            count += 1
        pass
    return count

def part2(data: PuzzleInput) -> int:
    """
    Given the same data as part 1, return the number of possible combinations that can be made for each input

    :param data: The puzzle input
    :return: The sum of how many different ways it is possible to create each pattern.
    """
    count = 0
    for towel in data['desired']:
        this_count = ways_to_form_target(towel, data['available'])
        print(f'{towel}: {this_count}')
        count += this_count
    return count

if __name__ == '__main__':
    start = timer()
    d = parse_input('data')
    print(f'It took {(timer() - start) * 1000:.2f}ms to parse the input')

    start = timer()
    result = part1(d)
    print(f'Part 1 result is {result} in {(timer() - start) * 1000:.2f}ms')

    start = timer()
    result = part2(d)
    print(f'Part 2 result is {result} in {(timer() - start) * 1000:.2f}ms')