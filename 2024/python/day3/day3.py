def parse_input(filename):
    with open(filename) as f:
        return f.read()


def get_number(data, index):
    result = 0
    while index < len(data) and data[index].isdigit():
        result *= 10
        result += int(data[index])
        index += 1
    return result, index

def part1(data):
    result = 0
    index = 0
    while index < len(data):
        index = data.find("mul(", index)
        if index == -1:
            return result
        index += 4
        x, index = get_number(data, index)
        if index >= len(data):
            return result
        elif data[index] != ',':
            index += 1
            continue
        else:
            index += 1
            y, index = get_number(data, index)
            if index >= len(data):
                return result
            elif data[index] != ')':
                index += 1
                continue
            else:
                index += 1
                result += x * y
    return result


def handle_m(data, index):
    if index + 8 >= len(data):
        return 0, index + 8
    if data[index + 1] == 'u' and data[index + 2] == 'l' and data[index + 3] == '(':
        index += 4
        x, index = get_number(data, index)
        if index >= len(data):
            return 0, index
        elif data[index] != ',':
            index += 1
            return 0, index
        else:
            index += 1
            y, index = get_number(data, index)
            if index >= len(data):
                return 0, index
            elif data[index] != ')':
                index += 1
                return 0, index
            else:
                index += 1
                return x * y, index
    return 0, index + 1

def handle_d(data, index):
    # can this be either?
    if index + 3 >= len(data):
        return None, index
    index += 1
    if data[index] == 'o' and data[index + 1] == '(' and data[index+2] == ')':
        return True, index + 3
    if index + 5 >= len(data):
        return None, index
    if data[index] == 'o' and data[index + 1] == 'n' and data[index + 2] == '\'' and data[index + 3] == 't' and data[index + 4] == '(' and data[index + 5] == ')':
        return False, index + 6
    return None, index

def part2(data):
    result = 0
    enabled = True
    index = 0
    while index < len(data):
        if data[index] == 'm':
            current, index = handle_m(data, index)
            if enabled:
                result += current
        elif data[index] == 'd':
            current, index = handle_d(data, index)
            if current is True:
                enabled = True
            elif current is False:
                enabled = False
        else:
            index += 1
    return result

if __name__ == '__main__':
    d = parse_input('data')
    print(part1(d))
    print(part2(d))