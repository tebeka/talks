'''
0: drjo014a089.embratel.net.br
1: -
2: -
3: [02/Jul/1995:16:30:08
4: -0400]
5: "GET
6: /shuttle/resources/orbiters/enterprise.html
7: HTTP/1.0"
8: 200
9: 9732
'''


def parse_line(line):
    fields = line.split()
    size = 0 if fields[-1] == '-' else int(fields[-1])
    return {
        'origin': fields[0],
        'time': fields[3] + ' ' + fields[4],
        'method': fields[5][1:],  # Remove leading "
        'path': fields[6],
        'status_code': int(fields[-2]),
        'size': size,
    }
