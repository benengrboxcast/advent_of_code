from timeit import default_timer as timer
from typing import TypedDict


class FileInfo(TypedDict):
    location: int
    size: int
    id: int

def parse_input(filename):
    count = 0
    data: list = []
    blank: bool = False
    id = 0

    with open(filename) as f:
        for line in f:
            line = line.strip()
            for c in line:
                current = int(c)
                fill_id = '.'
                if not blank:
                    fill_id = id
                    id += 1
                for i in range(current):
                    data.append(fill_id)
                blank = not blank
    print(count)
    return data

def sort_data(data):
    front_idx = 0
    back_idx = len(data) - 1
    while front_idx < back_idx:
        if data[front_idx] == '.':
            while data[back_idx] == '.' and back_idx > front_idx:
                back_idx -= 1
            if back_idx < front_idx:
                return data
            data[front_idx] = data[back_idx]
            data[back_idx] = '.'
        front_idx += 1
    return data

def checksum_data(data):
    front_idx = 0
    result = 0
    while front_idx < len(data) and data[front_idx] != '.':
        result += data[front_idx] * front_idx
        front_idx += 1
    return result

def part1(data):
    sort_data(data)
    return checksum_data(data)

def parse_part2(filename):
    file_list: list[FileInfo] = []
    blank_list: list[FileInfo] = []
    id = 0
    idx = 0
    blank: bool = False
    with open(filename) as f:
        for line in f:
            line = line.strip()
            for c in line:
                current = int(c)
                if blank:
                    blank_list.append({
                        'id': -1,
                        'size': current,
                        'location': idx,
                    })
                else:
                    file_list.append({
                        'id': id,
                        'size': current,
                        'location': idx,
                    })
                    id += 1
                idx += current
                blank = not blank
    return file_list, blank_list


def sort_whole_files(file_list: list[FileInfo], blanks: list[FileInfo]):
    file_index = len(file_list) - 1
    while file_index >= 0:
        file = file_list[file_index]
        for blank_index in range(len(blanks)):
            blank = blanks[blank_index]
            if blank['location'] >= file['location']:
                break
            if blank['size'] >= file['size']:
                file['location'] = blank['location']
                blank['size'] = blank['size'] - file['size']
                blank['location'] = blank['location'] + file['size']
        file_index -= 1

def sort_whole_files_2(file_list: list[FileInfo], blanks: list[FileInfo]):
    file_index = len(file_list) - 1

    while file_index >= 0:
        file = file_list[file_index]
        for blank_index in range(len(blanks)):
            blank = blanks[blank_index]
            if blank['location'] >= file['location']:
                break
            if blank['size'] >= file['size']:
                file['location'] = blank['location']
                blank['size'] = blank['size'] - file['size']
                blank['location'] = blank['location'] + file['size']
        file_index -= 1


def checksum_files(files: list[FileInfo]):
    result = 0
    for file in files:
        for i in range(file['size']):
            result += file['id'] * (i + file['location'])
    return result

def part2(files: list[FileInfo], blanks: list[FileInfo]):
    sort_whole_files(files, blanks)
    return checksum_files(files)

if __name__ == '__main__':
    start = timer()
    d = parse_input('data')

    print(f'It took {timer() - start:.3f} to parse the input')
    start = timer()
    r = part1(d)
    print(f'Part 1 took {timer() - start:.3f}: {r}')

    start = timer()
    f, b = parse_part2('data')
    r = part2(f, b)
    print(f'Part 2 took {timer() - start:.3f}: {r}')
