regs = {
    'r1': 0,
    'r2': 0,
}


def mov(dest, val):
    regs[dest] = val


def add(op1, op2, dest):
    regs[dest] = regs[op1] + regs[op2]


env = {
    'R1': 'r1',
    'R2': 'r2',
    'MOV': mov,
    'ADD': add,
}

code = '''
MOV(R1, 1)
MOV(R2, 2)
ADD(R1, R2, R2)
'''

exec(code, None, env)
print(regs)
