#!/usr/bin/env python

from enum import IntEnum
from inspect import getouterframes, currentframe

program = []  # List of instructions
labels = {}  # name -> location
num_regs = 8


def line_info():
    """Get file and line number in source file"""
    return getouterframes(currentframe())[2][2]


class ASMError(Exception):
    def __init__(self, message, lineno=0):
        super().__init__(message)
        self.lineno = lineno


class OpCodes(IntEnum):
    ADD_II = 0
    ADD_IR = 1
    ADD_RI = 2
    ADD_RR = 3
    SUB_II = 4
    SUB_IR = 5
    SUB_RI = 6
    SUB_RR = 7
    MOV_I = 8
    MOV_R = 9
    JMP = 10
    JMPE = 11
    CMP_I = 12
    CMP_R = 13


class Shift(IntEnum):
    opcode = 12
    slot0 = 8
    slot1 = 4
    slot2 = 0


class Instruction:
    def __init__(self, opcode, slot0, slot1, slot2, lineno):
        self.opcode, self.slot0, self.slot1, self.slot2, self.lineno = \
            opcode, slot0, slot1, slot2, lineno
        program.append(self)


def slot_bits(slot):
    # 1 -> 1, 'R2' -> 2
    return slot & 0xF if isinstance(slot, int) else int(slot[-1])


def is_jump(inst):
    return inst.opcode == OpCodes.JMP or inst.opcode == OpCodes.JMPE


def encode(inst):
    if is_jump(inst) and isinstance(inst.slot0, str):
        try:
            inst.slot0 = labels[inst.slot0]
        except KeyError:
            raise ASMError(f'Unkown label {inst.slot0!r}', inst.lineno)

    val = (inst.opcode << Shift.opcode) | \
          (slot_bits(inst.slot0) << Shift.slot0) | \
          (slot_bits(inst.slot1) << Shift.slot1) | \
          (slot_bits(inst.slot2) << Shift.slot2)
    return val & 0xFFFF


def label(name):
    if name in labels:
        lineno = line_info()
        raise ASMError(f'duplicate label - {name!r}', lineno)
    labels[name] = len(program)


def add(dest, src1, src2):
    opcode = {
        (int, int): OpCodes.ADD_II,
        (int, str): OpCodes.ADD_IR,
        (str, int): OpCodes.ADD_RI,
        (str, str): OpCodes.ADD_RR,
    }[(type(src1), type(src2))]

    return Instruction(opcode, dest, src1, src2, line_info())


def sub(dest, src1, src2):
    opcode = {
        (int, int): OpCodes.SUB_II,
        (int, str): OpCodes.SUB_IR,
        (str, int): OpCodes.SUB_RI,
        (str, str): OpCodes.SUB_RR,
    }[(type(src1), type(src2))]

    return Instruction(opcode, dest, src1, src2, line_info())


def jmp(target):
    return Instruction(OpCodes.JMP, target, 0, 0, line_info())


def jmpe(target):
    return Instruction(OpCodes.JMPE, target, 0, 0, line_info())


def cmp(lhs, rhs):
    opcode = OpCodes.CMP_I if isinstance(rhs, int) else OpCodes.CMP_R
    return Instruction(opcode, lhs, rhs, 0, line_info())


def mov(dest, src):
    opcode = OpCodes.MOV_I if isinstance(src, int) else OpCodes.MOV_R
    return Instruction(opcode, dest, src, 0, line_info())


env = {
    'LABEL': label,
    'JMP': jmp, 'JMPE': jmpe,
    'ADD': add, 'SUB': sub,
    'MOV': mov, 'CMP': cmp,
}
env.update({f'R{i}': f'R{i}' for i in range(num_regs)})


def asm_compile(infile):
    """Compile code from input file, return list of instructions"""
    code = infile.read()
    program.clear()
    exec(code, env, {})
    return program


def show(inst):
    return f'{encode(inst):016b}'


if __name__ == '__main__':
    from argparse import ArgumentParser, FileType

    parser = ArgumentParser(description='Tiny Assembler')
    parser.add_argument('infile', help='ASM input file', type=FileType('r'))
    parser.add_argument(
        '--out', help='output file', type=FileType('w'), default='-')
    parser.add_argument(
        '--debug', help='debug output', action='store_true', default=False)

    args = parser.parse_args()

    show = '{0.lineno:03d}: {0.__dict__!r}'.format if args.debug else show
    try:
        program = asm_compile(args.infile)
        for inst in program:
            print(show(inst), file=args.out)
    except (SyntaxError, ASMError) as err:
        raise SystemExit(f'error: {args.infile.name}:{err.lineno}: {err}')
