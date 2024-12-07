def parse_input(filename):
    result = []
    with open(filename) as f:
        for line in f.readlines():
            result.append(line.strip())
    return result

def check_up(data, row, col):
    test = 'MAS'
    if row < len(test):
        return False
    for i in range(len(test)):
        if data[row - i - 1][col] != test[i]:
            return False
    return True

def check_left_up(data, row, col):
    test = 'MAS'
    if row < len(test) or col < len(test):
        return False
    for i in range(len(test)):
        if data[row - i - 1][col - i - 1] != test[i]:
            return False
    return True

def check_left(data, row, col):
    test = 'MAS'
    if col < len(test):
        return False
    for i in range(len(test)):
        if data[row][col - i - 1] != test[i]:
            return False
    return True

def check_left_down(data, row, col):
    test = 'MAS'
    if row + len(test) >= len(data) or col < len(test):
        return False
    for i in range(len(test)):
        if data[row + i + 1][col - i - 1] != test[i]:
            return False
    return True

def check_down(data, row, col):
    test = 'MAS'
    if row + len(test) >= len(data):
        return False
    for i in range(len(test)):
        if data[row + i + 1][col] != test[i]:
            return False
    return True

def check_right_down(data, row, col):
    test = 'MAS'
    if row + len(test) >= len(data) or col + len(test) >= len(data[row]):
        return False
    for i in range(len(test)):
        if data[row + i + 1][col + i + 1] != test[i]:
            return False
    return True

def check_right(data, row, col):
    test = 'MAS'
    if col + len(test) >= len(data[row]):
        return False
    for i in range(len(test)):
        if data[row][col + i + 1] != test[i]:
            return False
    return True

def check_right_up(data, row, col):
    test = 'MAS'
    if row < len(test) or col + len(test) >= len(data[row]):
        return False
    for i in range(len(test)):
        if data[row - i - 1][col + i + 1] != test[i]:
            return False
    return True


def part1(data):
    count = 0
    for row in range(len(data)):
        line = data[row]
        for col in range(len(line)):
            if data[row][col] == 'X':
                # Check Straight Up
                if check_up(data, row, col):
                    count += 1
                if check_left_up(data, row, col):
                    count += 1
                if check_left(data, row, col):
                    count += 1
                if check_left_down(data, row, col):
                    count += 1
                if check_down(data, row, col):
                    count += 1
                if check_right_down(data, row, col):
                    count += 1
                if check_right(data, row, col):
                    count += 1
                if check_right_up(data, row, col):
                    count += 1
    return count

def part2_check_diag(data, row, col):
    c1 = data[row - 1][col - 1]
    c2 = data[row + 1][col + 1]
    if not ((c1 == 'M' and c2 == 'S') or (c1 == 'S' and c2 == 'M')):
        return False
    c1 = data[row + 1][col - 1]
    c2 = data[row - 1][col + 1]
    if not ((c1 == 'M' and c2 == 'S') or (c1 == 'S' and c2 == 'M')):
        return False
    return True

def part2_check_cross(data, row, col):
    c1 = data[row][col - 1]
    c2 = data[row][col + 1]
    if not ((c1 == 'M' and c2 == 'S') or (c1 == 'S' and c2 == 'M')):
        return False
    c1 = data[row - 1][col]
    c2 = data[row + 1][col]
    if not ((c1 == 'M' and c2 == 'S') or (c1 == 'S' and c2 == 'M')):
        return False
    return True

def part2(data):
    count = 0
    for row in range(1, len(data) - 1):
        for col in range(1, len(data[row]) - 1):
            if data[row][col] == 'A':
                if part2_check_diag(data, row, col):
                    count += 1
    return count

if __name__ == '__main__':
    d = parse_input('data')
    print(part1(d))
    print(part2(d))