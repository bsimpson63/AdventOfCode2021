import math

def explode(s, i):
    assert s[i] == "["

    j = i+1
    while s[j] != "]":
        assert s[j] == "," or s[j]
        j += 1

    pair = s[i+1:j].split(',')
    left, right = int(pair[0]), int(pair[1])

    e = ""
    for i_left in range(i-1, 0, -1):
        if s[i_left].isdigit():
            # we found a number to the left
            if s[i_left-1].isdigit():
                left_neighbor = int(s[i_left-1:i_left+1])
                left_neighbor += left
                e = s[:i_left-1]
                e += str(left_neighbor)
                e += s[i_left+1:i]
            else:
                left_neighbor = int(s[i_left])
                left_neighbor += left
                e = s[:i_left]
                e += str(left_neighbor)
                e += s[i_left+1:i]
            break
    else:
        e = s[:i]

    e += "0"

    for i_right in range(j, len(s)):
        if s[i_right].isdigit():
            if s[i_right+1].isdigit():
                right_neighbor = int(s[i_right:i_right+2])
                right_neighbor += right
                e += s[j+1:i_right]
                e += str(right_neighbor)
                e += s[i_right+2:]
            else:
                right_neighbor = int(s[i_right])
                right_neighbor += right
                e += s[j+1:i_right]
                e += str(right_neighbor)
                e += s[i_right+1:]
            break
    else:
        e += s[j+1:]

    return e

def split(s, i):
    n = int(s[i:i+2])
    l = math.floor(n/2)
    r = math.ceil(n/2)

    sp = s[:i]
    sp += f"[{l},{r}]"
    sp += s[i+2:]
    return sp

def reduce(s):
    open_count = 0
    for i, c in enumerate(s):
        if c == '[':
            open_count += 1
            if open_count == 5:
                #print(f"exploding")
                #print(s)
                #print(" "*i + "^")
                #print("*"*10)
                s = explode(s, i)
                #return s
                return reduce(s)
        elif c == ']':
            open_count -= 1

    for i in range(len(s)-1):
        try:
            n = int(s[i:i+2])
        except ValueError:
            continue

        #print(f"splitting")
        #print(s)
        #print(" "*i + "^")
        #print("*"*10)

        s = split(s, i)
        return reduce(s)
    
    return s

def add(s1, s2):
    return f"[{s1},{s2}]"

def get_magnitude(s):
    # The magnitude of a pair is 3 times the magnitude of its left element
    # plus 2 times the magnitude of its right element. The magnitude of a
    # regular number is just that number.

    try:
        val = int(s)
        return val
    except ValueError:
        pass

    open_count = 0
    s = s[1:-1]
    for i, c in enumerate(s):
        if c == "[":
            open_count += 1
        elif c == "]":
            open_count -= 1
        elif c == "," and open_count == 0:
            l = s[:i]
            r = s[i+1:]
            lval = get_magnitude(l)
            rval = get_magnitude(r)
            return 3*lval + 2*rval


def main():
    with open("./input.txt") as f:
        txt = f.read()
 
    numbers = txt.split('\n')
    sums = {}
    for n1 in numbers:
        for n2 in numbers:
            if n1 == n2:
                continue
            s = add(n1, n2)
            s = reduce(s)
            m = get_magnitude(s)
            sums[(n1, n2)] = m

    tups = sorted(sums.items(), key=lambda t: t[1], reverse=True)
    nums, m = tups[0]
    print(f"{nums}: has highest magnitude {m}")


if __name__ == "__main__":
    main()