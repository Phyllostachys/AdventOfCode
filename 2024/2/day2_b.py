import sys

INCREASE = -1
DECREASE = 1

def is_safe(data: list[int]) -> bool:
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

def is_safe_with_dampening(data):
    # check if the data is safe without dampening
    if is_safe(data):
        return True

    # check every version of the data where we remove one value
    for i in range(len(data)):
        if is_safe(data[:i] + data[i+1:]):
            return True

    return False

count = 0
for line in sys.stdin:
    data = list(map(int, line.strip().split()))
    res = is_safe_with_dampening(data)
    # print(res)
    if res:
        count += 1
print(count)
