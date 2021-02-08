import operator
from functools import reduce
from string import digits
from typing import Tuple

operators_dict = {
    '+': operator.add,
    '*': operator.mul,
}

def solve(expression: str) -> Tuple[int, int]:
    chars_count = inner_chars = partial = 0
    op = operator.add

    for index, char in enumerate(expression):
        chars_count += 1
        if inner_chars > 0:
            inner_chars -= 1
            continue
        if char == '(':
            inner, inner_chars = solve(expression[index + 1:])
            partial = op(partial, inner)
        elif char == ')':
            return partial, chars_count
        elif char in digits:
            partial = op(partial, int(char))
        else:
            op = operators_dict[char]

    return partial, chars_count


def solve_add_first(expression: str) -> Tuple[int, int]:
    chars_count = inner_chars = partial = 0
    numbers = []
    add = False

    for index, char in enumerate(expression):
        chars_count += 1
        if inner_chars > 0:
            inner_chars -= 1
            continue
        if char == '(':
            inner, inner_chars = solve_add_first(expression[index + 1:])
            if add:
                numbers[-1] += inner
                add = False
            else:
                numbers.append(inner)
        elif char == ')':
            break
        elif char in digits:
            if add:
                numbers[-1] += int(char)
                add = False
            else:
                numbers.append(int(char))
        elif char == '+':
            add = True

    return reduce(lambda x, y: x*y, numbers), chars_count

total = 0
total_add_first = 0

with open('input.txt') as f:
    for line in f.readlines():
        line = line.strip().replace(' ', '')
        solution, _ = solve(line)
        solution_add_first, _ = solve_add_first(line)
        # print(f'{line} = {solution}')
        total += solution
        total_add_first += solution_add_first

print(total)
print(total_add_first)
