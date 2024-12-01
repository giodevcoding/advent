def get(depth, a, b):
    for c, d, dst in map_ranges[depth]:
        if c <= a < d:
            if b < d:
                new_a, new_b = a - c + dst, b - c + dst
                return new_a if depth == 6 else get(depth + 1, new_a, new_b)
            else:
                return min(get(depth, a, d - 1), get(depth, d, b))
    return a if depth == 6 else get(depth + 1, a, b)


with open("input.txt") as f:
    ls = f.read().split("\n\n")

seeds, maps = [int(n) for n in ls[0].split(": ")[1].split()], [
    [[int(m) for m in n.split()] for n in l.split(":\n")[1].splitlines()]
    for l in ls[1:]
]

map_ranges = [[] for _ in range(7)]
for depth, m in enumerate(maps):
    for dst, src, l in m:
        map_ranges[depth].append([src, src + l, dst])

print(min([get(0, seeds[i], seeds[i] + seeds[i + 1]) for i in range(0, len(seeds), 2)]))
