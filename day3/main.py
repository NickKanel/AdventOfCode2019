def main():
    wire1 = set()
    wire2 = set()

    with open('input', 'r') as f:
        wire1 = build_wire(f.readline())
        wire2 = build_wire(f.readline())

    print(len(wire1))
    print(len(wire2))

    points = list()
    for point in wire1:
        if point in wire2:
            points.append(point)

    lowest = points[0]
    for point in points:
        if point == (0, 0):
            continue
        if m_distance(point) < m_distance(lowest):
            lowest = point
    print(points)
    print(lowest)

def m_distance(point):
    return abs(point[0]) + abs(point[1])

def build_wire(line):
    print(sum(list(map(lambda x: int(x[1:]), line.split(',')))))

    current_point = (0, 0)
    wire = set()
    wire.add(current_point)

    paths = line.split(',')
    for path in paths:
        direction = path[0]
        distance = int(path[1:])

        if direction == 'R':
            for i in range(1, distance + 1):
                current_point = (current_point[0] + 1, current_point[1])
                wire.add(current_point)
        if direction == 'L':
            for i in range(1, distance + 1):
                current_point = (current_point[0] - 1, current_point[1])
                wire.add(current_point)
        if direction == 'U':
            for i in range(1, distance + 1):
                current_point = (current_point[0], current_point[1] + 1)
                wire.add(current_point)
        if direction == 'D':
            for i in range(1, distance + 1):
                current_point = (current_point[0], current_point[1] - 1)
                wire.add(current_point)

    return wire

if __name__ == '__main__':
    main()
