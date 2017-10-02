#!/usr/bin/env python

from asm import num_regs, OpCodes, Shifts

insts = {}  # opcode -> inst


class SimError(Exception):
    pass


class Simulator:
    def __init__(self, program, regs=None, *, trace=False):
        self.program = program
        self.trace = trace
        self.regs = [0] * num_regs
        for reg, val in (regs or {}).items():
            self.regs[reg] = val

        self.ip = -1
        self.cmp_flag = False

    def parse(self, inst):
        opcode = (inst >> Shifts.opcode) & 0xF
        slot0 = (inst >> Shifts.slot0) & 0xF
        slot1 = (inst >> Shifts.slot1) & 0xF
        slot2 = (inst >> Shifts.slot2) & 0xF

        return opcode, slot0, slot1, slot2

    def run(self):
        self.ip = -1
        while True:
            self.ip += 1
            if self.ip >= len(self.program):
                break

            opcode, slot0, slot1, slot2 = self.parse(self.program[self.ip])
            fn = insts.get(opcode)
            if fn is None:
                raise SimError(f'{self.ip:03d}: unknown opcode - {opcode}')

            if self.trace:
                print(','.join('{:02d}'.format(reg) for reg in self.regs))
                print(f'{self.ip:02d}: {fn.__name__} {slot0} {slot1} {slot2}')

            fn(self, slot0, slot1, slot2)


def inst(opcode):
    def dec(fn):
        insts[opcode] = fn
        return fn
    return dec


@inst(OpCodes.ADD_II)
def add_ii(sim, slot0, slot1, slot2):
    sim.regs[slot0] = slot1 + slot2


@inst(OpCodes.ADD_IR)
def add_ir(sim, slot0, slot1, slot2):
    sim.regs[slot0] = slot1 + sim.regs[slot2]


@inst(OpCodes.ADD_RI)
def add_ri(sim, slot0, slot1, slot2):
    sim.regs[slot0] = sim.regs[slot1] + slot2


@inst(OpCodes.ADD_RR)
def add_rr(sim, slot0, slot1, slot2):
    sim.regs[slot0] = sim.regs[slot1] + sim.regs[slot2]


@inst(OpCodes.SUB_II)
def sub_ii(sim, slot0, slot1, slot2):
    sim.regs[slot0] = slot1 - slot2


@inst(OpCodes.SUB_IR)
def sub_ir(sim, slot0, slot1, slot2):
    sim.regs[slot0] = slot1 - sim.regs[slot2]


@inst(OpCodes.SUB_RI)
def sub_ri(sim, slot0, slot1, slot2):
    sim.regs[slot0] = sim.regs[slot1] - slot2


@inst(OpCodes.SUB_RR)
def sub_rr(sim, slot0, slot1, slot2):
    sim.regs[slot0] = sim.regs[slot1] - sim.regs[slot2]


@inst(OpCodes.MOV_I)
def mov_i(sim, slot0, slot1, slot2):
    sim.regs[slot0] = slot1


@inst(OpCodes.MOV_R)
def mov_r(sim, slot0, slot1, slot2):
    sim.regs[slot0] = sim.regs[slot1]


@inst(OpCodes.JMP)
def jmp(sim, slot0, slot1, slot2):
    sim.ip = slot0 - 1


@inst(OpCodes.JMPE)
def jmpe(sim, slot0, slot1, slot2):
    if sim.cmp_flag:
        sim.ip = slot0 - 1


@inst(OpCodes.CMP_I)
def cmpi(sim, slot0, slot1, slot2):
    sim.cmp_flag = sim.regs[slot0] == slot1


@inst(OpCodes.CMP_R)
def cmpr(sim, slot0, slot1, slot2):
    sim.cmp_flag = sim.regs[slot0] == sim.regs[slot1]


if __name__ == '__main__':
    from argparse import ArgumentParser, FileType

    parser = ArgumentParser(description='Simulator')
    parser.add_argument(
        'infile', help='input file', type=FileType('r'), default='-',
        nargs='?')
    for i in range(num_regs):
        parser.add_argument(
            f'-R{i}', help=f'set value of R{i}', type=int, default=0)
    parser.add_argument(
        '--trace', help='print trace', action='store_true', default=False)
    args = parser.parse_args()

    program = [int(line, 2) & 0xFFFF for line in args.infile]
    regs = {i: getattr(args, f'R{i}') for i in range(num_regs)}
    sim = Simulator(program, regs, trace=args.trace)

    sim.run()

    for i, reg in enumerate(sim.regs):
        print(f'R{i}: {reg}')
