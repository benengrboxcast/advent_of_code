from enum import Enum
import re

class Instruction(Enum):
    adv = 0
    bxl = 1
    bst = 2
    jnz = 3
    bxc = 4
    out = 5
    bdv = 6
    cdv = 7

class Handheld():
    def __init__(self, a, b, c, inst):
        self._registers = [a ,b, c]
        self._ip = 0
        self._code = inst
        self._output = []

    def step(self):
        instruction: Instruction = Instruction(self._code[self._ip])
        operand = self._code[self._ip + 1]

        if instruction == Instruction.adv:
            denom = self.get_combo(operand)
            result = self._registers[0] >> denom
            self._registers[0] = result
        elif instruction == Instruction.bxl:
            result = self._registers[1] ^ operand
            self._registers[1] = result
        elif instruction == Instruction.bst:
            combo = self.get_combo(operand)
            result = combo & 0x7
            self._registers[1] = result
        elif instruction == Instruction.jnz:
            if self._registers[0] != 0:
                self._ip = operand - 2 # We are going to increase by 2 at the end of this function
        elif instruction == Instruction.bxc:
            result = self._registers[1] ^ self._registers[2]
            self._registers[1] = result
        elif instruction == Instruction.out:
            combo = self.get_combo(operand)
            combo = combo & 0x7
            self._output.append(combo)
        elif instruction == Instruction.bdv:
            denom = self.get_combo(operand)
            result = self._registers[0] >> denom
            self._registers[1] = result
        elif instruction == Instruction.cdv:
            denom = self.get_combo(operand)
            result = self._registers[0] >> denom
            self._registers[2] = result

        self._ip += 2

    def get_combo(self, operand: int):
        if 0 <= operand <= 3:
            return operand
        elif 4 <= operand <= 6:
            return self._registers[operand - 4]
        raise ValueError(f"Invalid operand: {operand}")

    @property
    def halted(self):
        return self._ip > len(self._code) - 1

    @property
    def output(self):
        return self._output


def part1(a, b, c, inst):
    comp = Handheld(a, b, c, inst)
    while not comp.halted:
        comp.step()
    return comp.output

def find_a(program, a, b, c, prg_pos):
    if abs(prg_pos) > len(program): return a
    for i in range(8):
        c = Handheld(a * 8 + i, b, c, program)
        while not c.halted:
            c.step()
        first_digit = c.output[0]
        if first_digit == program[prg_pos]:
            e = find_a(program, a * 8 + i, b, c, prg_pos - 1)
            if e: return e

def part2(b, c, inst):
    return find_a(inst, 0, b, c, -1)

if __name__ == "__main__":
    with open("data") as f:
        a, b, c, *inst = list(map(int, re.findall(r'\d+', f.read())))
    print(part1(a, b, c, inst))
    print(part2(b, c, inst))