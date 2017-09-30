#!/usr/bin/env python

from enum import IntEnum
from abc import ABC, abstractmethod
from inspect import getouterframes, currentframe

program = []  # List of instructions
labels = {}  # name -> location
instructions = {}  # name -> class
num_regs = 8


def instruction(cls):
    """Decorator to register instruction"""
    instructions[cls.__name__] = cls
    return cls


def line_info(depth=3):
    """Get file and line number in source file"""
    try:
        return getouterframes(currentframe())[depth][2]
    except:
        return 0


class ASMError(Exception):
    def __init__(self, message, line=0):
        super().__init__(message)
        self.line = line


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


class Shifts(IntEnum):
    Code = 12
    Slot0 = 8
    Slot1 = 4
    Slot2 = 0


class REG:
    def __init__(self, code):
        self.code = code

    def __repr__(self):
        name = self.__class__.__name__
        return f'{name}({self.code!r})'


class ASM(ABC):
    def __init__(self):
        self.line = line_info()
        self.name = self.__class__.__name__
        program.append(self)

    def slot_bits(self, slot):
        if isinstance(slot, int):
            return slot & 0xF
        return slot.code

    def code(self, opcode, slot1, slot2=0, slot3=0):
        return \
            (opcode << Shifts.Code) | \
            (self.slot_bits(slot1) << Shifts.Slot0) | \
            (self.slot_bits(slot2) << Shifts.Slot1) | \
            (self.slot_bits(slot3) << Shifts.Slot2)

    def __str__(self):
        val = self.bits() & 0xFFFF
        return f'{val:016b}'

    @abstractmethod
    def bits(self):
        return 0


class ALU3(ASM):
    """ALU instruction with 3 operands"""
    def __init__(self, dest, src1, src2):
        super().__init__()
        self.dest = dest
        self.src1 = src1
        self.src2 = src2

    def bits(self):
        opcode = self.opcodes[(type(self.src1), type(self.src2))]
        return self.code(opcode, self.dest, self.src1, self.src2)

    def __repr__(self):
        return f'{self.name}({self.dest!r}), {self.src1!r}, {self.src2!r})'


@instruction
class ADD(ALU3):
    opcodes = {
        (int, int): OpCodes.ADD_II,
        (int, REG): OpCodes.ADD_IR,
        (REG, int): OpCodes.ADD_RI,
        (REG, REG): OpCodes.ADD_RR,
    }


@instruction
class SUB(ALU3):
    opcodes = {
        (int, int): OpCodes.SUB_II,
        (int, REG): OpCodes.SUB_IR,
        (REG, int): OpCodes.SUB_RI,
        (REG, REG): OpCodes.SUB_RR,
    }


@instruction
class JMP(ASM):
    opcode = OpCodes.JMP

    def __init__(self, target):
        super().__init__()
        self.target = target

    def bits(self):
        target = self.target
        if isinstance(target, str):
            target = labels[target]

        return self.code(self.opcode, target)

    def __repr__(self):
        return f'{self.name}({self.target!r})'


@instruction
class JMPE(JMP):
    opcode = OpCodes.JMPE


@instruction
class CMP(ASM):
    def __init__(self, lhs, rhs):
        super().__init__()
        self.lhs = lhs
        self.rhs = rhs

    def bits(self):
        opcode = OpCodes.CMP_I if isinstance(self.rhs, int) else OpCodes.CMP_R
        return self.code(opcode, self.lhs, self.rhs)

    def __repr__(self):
        return f'{self.name}({self.lhs!r}, {self.rhs!r})'


@instruction
def LABEL(name):
    if name in labels:
        line = line_info(depth=2)
        raise ASMError(f'duplicate label - {name!r}', line)
    labels[name] = len(program)


@instruction
class MOV(ASM):
    def __init__(self, dest, src):
        super().__init__()
        self.dest = dest
        self.src = src

    def bits(self):
        opcode = OpCodes.MOV_I if isinstance(self.src, int) else OpCodes.MOV_R
        return self.code(opcode, self.dest, self.src)

    def __repr__(self):
        return f'{self.name}({self.dest!r}, {self.src!r})'


def asm_compile(infile):
    program.clear()

    env = instructions.copy()
    for i in range(num_regs):
        env[f'R{i}'] = REG(i)

    code = infile.read()
    exec(code, env, {})
    return program


def show_debug(inst):
    return f'{inst.line:03d}: {inst!r}'


if __name__ == '__main__':
    from argparse import ArgumentParser, FileType

    parser = ArgumentParser(description='Tiny Assembler')
    parser.add_argument('infile', help='ASM input file', type=FileType('r'))
    parser.add_argument(
        '--out', help='output file', type=FileType('w'), default='-')
    parser.add_argument(
        '--debug', help='debug output', action='store_true', default=False)

    args = parser.parse_args()
    try:
        program = asm_compile(args.infile)
    except ASMError as err:
        raise SystemExit(f'error: {args.infile.name}:{err.line}: {err}')

    show = show_debug if args.debug else str

    for inst in program:
        print(show(inst), file=args.out)
