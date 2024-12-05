import sys

INCREASE = -1
DECREASE = 1

def process_line(data):
    last = data[0]
    direction = 0
    for num in data[1:]:
        diff = abs(num - last)
        if diff == 0 or diff > 3:
            return False
        if direction == INCREASE:
            if num < last:
                return False
        elif direction == DECREASE:
            if num > last:
                return False
        else:
            if num < last:
                direction = DECREASE
            elif num > last:
                direction = INCREASE
            else:
                return False # this shouldn't happen cause the diff is detected above
        last = num

    return True


count = 0
for line in sys.stdin:
    data = list(map(int, line.strip().split()))
    res = process_line(data)
    print(res)
    if res:
        count += 1
print(count)
