import fileinput

for line in fileinput.input():
    name = fileinput.filename()
    lnum = fileinput.lineno()
    count = len(line.split())
    print(f'{name}:{lnum}: {count}')
