import copy
import math
import functools
import numpy as np

home = (22,19)

def main():
    asteroids = set()
    line_num = 0
    with open('input', 'r') as f:
        for line in f:
            for i, c in enumerate(line.strip()):
                if c == '#':
                    asteroids.add((i,line_num))
            line_num += 1

    max_visible = 0
    for asteroid in asteroids:
        points = get_visible_points(asteroid, asteroids)
        if len(points) > max_visible:
            max_visible = len(points)

    print("part 1", max_visible - 1)
    
    asteroids.remove(home)
    asteroids = get_visible_points(home, asteroids)
    asteroids_relative = list(map(lambda x: (x[0] - home[0], x[1] - home[1]), asteroids))
    asteroids = sort_radially(asteroids_relative)
    asteroids = list(map(lambda target: (target[0] + home[0], target[1] + home[1]), asteroids))
    target = asteroids[199]
    print("part 2", target)

def angle(p2):
    ang1 = np.arctan2(-1, 0)
    ang2 = np.arctan2(p2[1], p2[0])
    return np.rad2deg((ang2 - ang1) % (2 * np.pi))

def sort_radially(asteroids):
    return sorted(asteroids, key=angle)

def get_visible_points(source, others):
    others_copy = copy.deepcopy(others)
    for a in others:
        if a == source:
            continue
        for b in others:
            if b == source:
                continue
            if get_slope(source, a) == get_slope(source, b):
                if quadrant(source, a) != quadrant(source, b):
                    continue
                if get_distance(source, a) > get_distance(source, b):
                    others_copy.remove(a)
                    break
    return others_copy

def quadrant(source, dest):
    if source[0] - dest[0] > 0:
        if source[1] - dest[1] < 0:
            return 1
        else:
            return 2
    else:
        if source[1] - dest[1] < 0:
            return 3
        else:
            return 4

def get_slope(source, dest):
    slope = (source[1] - dest[1], source[0] - dest[0])
    gcd = math.gcd(slope[0], slope[1])
    slope = (int(slope[0]/gcd), int(slope[1]/gcd))
    return slope

def get_distance(source, dest):
    return abs(source[1] - dest[1]) + abs(source[0] - dest[0])

if __name__ == '__main__':
    main()
