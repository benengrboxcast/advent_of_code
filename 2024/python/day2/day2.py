def parse_data(filename):
    result = []
    with open(filename) as f:
        for line in f:
            data_set = []
            points = line.split(' ')
            if len(points) > 0:
                for p in points:
                    data_set.append(int(p))
                result.append(data_set)
    return result

def is_safe(line):
    if len(line) < 2:
        return False
    asc = line[1] > line[0]
    for i in range(1, len(line)):
        delta = line[i] - line[i-1]
        if abs(delta) > 3 or abs(delta) == 0:
            return False
        if delta > 0 and not asc:
            return False
        if delta < 1 and asc:
            return False
    return True

def part1(data):
    count = 0
    for line in data:
        if is_safe(line):
            count += 1
    return count

def part2(data):
    count = 0
    for line in data:
        if is_safe(line):
            count += 1
            continue
        for i in range(0, len(line)):
            temp = line.pop(i)
            if is_safe(line):
                count += 1
                break
            line.insert(i, temp)
    return count


if __name__ == '__main__':
    d = parse_data('data')
    print(part1(d))
    print(part2(d))

