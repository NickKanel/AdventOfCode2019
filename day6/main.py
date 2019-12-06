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

            # if orbiter.
            orbiter.orbiting = center
            center.orbiters.add(orbiter)

    summ = 0
    for key, obj in objects.items():
        #print(key, obj.name)
        summ += get_distance_to_com(obj)
    print(summ)

def get_distance_to_com(obj):
    # print('checking ' + obj.name)
    if obj.name == 'COM':
        return 0
    return 1 + get_distance_to_com(obj.orbiting)

# def distance_from_to(source, dest):
#     shorte

if __name__ == '__main__':
    main()
