def main():
    wire1str = ''
    wire2str = ''

    with open('input', 'r') as f:
        wire1str = f.readline()
        wire2str = f.readline()
    
    wire1 = build_wire(wire1str)
    wire2 = build_wire(wire2str)

    print(len(wire1))
    print(len(wire2))

    points = list()
    for point in wire1:
        if point in wire2:
            points.append(point)

    lowest_total_distance = len(wire1) + len(wire2)
    for point in points:
        if point == (0, 0):
            continue

        wire1_distance = get_distance(point, wire1str)
        wire2_distance = get_distance(point, wire2str)

        if wire1_distance + wire2_distance < lowest_total_distance:
            lowest_total_distance = wire2_distance + wire1_distance
    
    print(points)
    print(lowest_total_distance)

def m_distance(point):
    return abs(point[0]) + abs(point[1])

def get_distance(point, wirestr):
    if point == (0,0):
        return 0

    current_point = (0,0)

    traveled_distance = 0
    paths = wirestr.split(',')
    for path in paths:
        direction = path[0]
        distance = int(path[1:])

        if direction == 'R':
            for i in range(1, distance + 1):
                traveled_distance += 1
                current_point = (current_point[0] + 1, current_point[1])
                if current_point == point:
                    return traveled_distance
        if direction == 'L':
            for i in range(1, distance + 1):
                traveled_distance += 1
                current_point = (current_point[0] - 1, current_point[1])
                if current_point == point:
                    return traveled_distance
        if direction == 'U':
            for i in range(1, distance + 1):
                traveled_distance += 1
                current_point = (current_point[0], current_point[1] + 1)
                if current_point == point:
                    return traveled_distance
        if direction == 'D':
            for i in range(1, distance + 1):
                traveled_distance += 1
                current_point = (current_point[0], current_point[1] - 1)
                if current_point == point:
                    return traveled_distance

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
