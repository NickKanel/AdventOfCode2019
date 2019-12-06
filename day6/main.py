class Object:
    def __init__(self, name):
        self.name     = name
        self.orbiters = set()
        self.orbiting = None

objects = dict()

def main():
    with open('input', 'r') as f:
        for line in f:
            center = None
            orbiter = None

            object_names = line.split(')')
            first = object_names[0].strip()
            second = object_names[1].strip()
            if first in objects:
                center = objects[first]
            else:
                center = Object(first)
                objects[center.name] = center

            if second in objects:
                orbiter = objects[second]
            else:
                orbiter = Object(second)
                objects[orbiter.name] = orbiter

            orbiter.orbiting = center
            center.orbiters.add(orbiter)

    print(distance_from_to(objects['YOU'].orbiting, objects['SAN'].orbiting))

def get_distance_to_com(obj):
    if obj.name == 'COM':
        return 0
    return 1 + get_distance_to_com(obj.orbiting)

def distance_from_to(source, dest):
    visited = set()
    distances = dict()
    for _, obj in objects.items():
        distances[obj] = 1000000000
    distances[source] = 0

    current = source
    while len(visited) < len(objects):
        visited.add(current)
        for child in get_nodes(current):
            if 1 + distances[current] <= distances[child]:
                distances[child] = 1 + distances[current]
        
        next_node = None
        lowest_score = 1000000000
        for child in get_nodes(current):
            if child in visited:
                continue
            if distances[child] <= lowest_score:
                next_node = child
                lowest_score = distances[child]
        for obj in visited:
            for child in get_nodes(obj):
                if child in visited:
                    continue
                if distances[child] <= lowest_score:
                    next_node = child
                    lowest_score = distances[child]
        current = next_node

    return distances[dest]

def get_nodes(obj):
    if obj.orbiting == None:
        return list(obj.orbiters)
    return [obj.orbiting] + list(obj.orbiters)

if __name__ == '__main__':
    main()
