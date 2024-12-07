from sortedcontainers import SortedDict, SortedList
from timeit import default_timer as timer

def parse_input(filename):
    node_list = []
    for i in range(0, 100):
        node = {
            'id': i,
            'prev': SortedList(),
            'next': SortedList(),
        }
        node_list.append(node)
    print_sequences = []
    doing_sequences = False
    with open(filename, 'r') as f:
        for line in f:
            line = line.strip()
            if len(line) == 0:
                doing_sequences = True
            elif doing_sequences:
                seq = []
                values = line.split(',')
                for value in values:
                    seq.append(int(value))
                print_sequences.append(seq)
            else:
                values = line.split('|')
                first = int(values[0])
                last = int(values[1])
                node_list[first]['next'].add(last)
                node_list[last]['prev'].add(first)
    return node_list, print_sequences

def order_sequence(rules, seq):
    changed = False
    for i in range(len(seq) // 2 + 1):
        for j in range(i + 1, len(seq)):
            if rules[seq[j]]['next'].count(seq[i]) > 0:
                temp = seq.pop(j)
                seq.insert(i, temp)
                j += 1
                changed = True
    return seq[len(seq) // 2], changed

def day5(rules, sequences):
    part1 = 0
    part2 = 0
    for seq in sequences:
        ok = True
        for x in range(len(seq)):
            for y in range(len(seq)):
                if y < x:
                    if rules[seq[x]]['next'].count(seq[y]) > 0:
                        ok = False
                        break
                elif y > x:
                    if rules[seq[x]]['prev'].count(seq[y]) > 0:
                        ok = False
                        break
            if not ok:
                continue
        if ok:
            middle_index = len(seq) // 2
            part1 += seq[middle_index]
        else:
            result, changed = order_sequence(rules, seq)
            part2 += result
    return part1, part2


def sort_all(rules, sequences):
    part1 = 0
    part2 = 0
    for s in sequences:
        result, changed = order_sequence(rules, s)
        if changed:
            part2 += result
        else:
            part1 += result
    return part1, part2


if __name__ == '__main__':
    start = timer()
    nodes, prints = parse_input('data')
    print(f'It took {timer() - start:.6f}s to parse the input')

    start = timer()
    print(day5(nodes, prints))
    print(timer() - start)

    nodes, prints = parse_input('data')
    start = timer()
    print(sort_all(nodes, prints))
    print(timer() - start)