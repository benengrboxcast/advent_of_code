output = open('reformatted.txt', 'w')
with open('data') as f:
    for line in f:
        line = line.strip()
        for c in line:
            if c == 'Y':
                output.write(c)
            else:
                output.write('.')
        output.write('\n')
output.close()