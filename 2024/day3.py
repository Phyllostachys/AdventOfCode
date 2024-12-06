import sys

SCANNING = 0
FOUND_M = 1
FOUND_U = 2
FOUND_L = 3
FOUND_LPAREN = 4
FOUND_COMMA = 5
FOUND_RPAREN = 6

def parse1(line: str) -> int:
    value = 0
    state = SCANNING
    first_str = ''
    first_num = 0
    second_str = ''
    second_num = 0
    for c in line:
        if state == FOUND_M:
            if c == 'u':
                state = FOUND_U
                continue
        elif state == FOUND_U:
            if c == 'l':
                state = FOUND_L
                continue
        elif state == FOUND_L:
            if c == '(':
                state = FOUND_LPAREN
                continue
        elif state == FOUND_LPAREN:
            if 48 <= ord(c) <= 57:
                first_str += c
                continue
            elif c == ',':
                first_num = int(first_str)
                state = FOUND_COMMA
                continue
        elif state == FOUND_COMMA:
            if 48 <= ord(c) <= 57:
                second_str += c
                continue
            elif c == ')':
                second_num = int(second_str)
                value += first_num * second_num
                print(first_num, "*", second_num, "=", value)
        else:
            if c == 'm':
                state = FOUND_M
                continue

        state = SCANNING
        first_str = ''
        second_str = ''
    return value

def parse2(line: str) -> int:
    index = 0
    value = 0
    enabled = True
    while index < len(line):
        print(index)
        c = line[index]
        if c == 'm' and line[index+1] == 'u' and line[index+2] == 'l' and line[index+3] == '(':
            subindex = 4
            comma = 0
            right_paren = 0
            while subindex < len(line):
                print("  ", index + subindex)
                new_c = line[index + subindex]
                if new_c == ',':
                    comma = subindex
                    subindex += 1
                elif new_c == ')':
                    right_paren = subindex
                    break
                elif 48 <= ord(new_c) <= 57:
                    subindex += 1
                    continue
                else:
                    break
            if right_paren != 0 and enabled:
                first = int(line[index + 4:index + comma])
                second = int(line[index + comma + 1:index + right_paren])
                print(line[index:index+right_paren])
                print(first, second)
                value += (first * second)
                index += right_paren+1
                continue

        elif c == 'd' and line[index+1] == 'o':
            print('detected a do!')
            if line[index+2] == '(' and line[index+3] == ')':
                enabled = True
                print("enabled!")
                index += 4
                continue
            elif line[index+2] == 'n' and line[index+3] == '\'' and line[index+4] == 't' and line[index+5] == '(' and line[index+6] == ')':
                enabled = False
                print("disabled!")
                index += 7
                continue

        index += 1
    return value

lines = []
results_per_line = []
for line in sys.stdin:
    res = parse1(line)
    results_per_line.append(res)

    lines.append(line)
print("part 1:", sum(results_per_line))
input = "".join(lines)
res = parse2(input)
print("part 2:", res)
