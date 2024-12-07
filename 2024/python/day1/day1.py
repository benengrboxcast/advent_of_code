def process_data(filename):
    col1 = []
    col2 = []
    with open(filename, 'r') as f:
        for line in f:
            temp = line.split()
            col1.append(int(temp[0]))
            col2.append(int(temp[1]))
    return col1, col2

def problem1(data):
    col1, col2 = data[0], data[1]
    col1.sort()
    col2.sort()
    result = 0
    for i in range(len(col1)):
        result += abs(col1[i] - col2[i])
    return result

def problem2(data):
    col1, col2 = data[0], data[1]
    col1.sort()
    col2.sort()
    ri = 0
    current = 0
    result = 0
    for li in range(len(col1)):
        example = col1[li]
        if li > 0 and col1[li-1] == example:
            result += current
        else:
            while ri < len(col2) and col2[ri] < example:
                ri += 1
            if ri >= len(col2):
                return result
            else:
                count = 0
                while ri < len(col2) and col2[ri] == example:
                    count += 1
                    ri += 1
                current = example * count
                result += current
    return result

if __name__ == '__main__':
    a, b = process_data('day1.data')
    print(problem1([a, b]))
    print(problem2([a, b]))
