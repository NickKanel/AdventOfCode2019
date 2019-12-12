package main

import (
    "log"
    "crypto/md5"
    "fmt"
    // "math"
)

type Point struct {
    x int
    y int
    z int
    velocity Velocity
}

type Velocity struct {
    x int
    y int
    z int
}

func getInitialPoints() []Point {
    return []Point {
        Point {
            13, 9, 5, Velocity {},
        },
        Point {
            8, 14, -2, Velocity {},
        },
        Point {
            -5, 4, 11, Velocity {},
        },
        Point {
            2, -6, 1, Velocity {},
        },
    }
}

func main() {
    points := getInitialPoints()

    for i := 0; i < 1000; i++ {
        points = stepTime(points)
    }
    sum := 0
    for i := 0; i < len(points); i++ {
        sum += getEnergy(points[i])
    }
    log.Println(sum)
}

func stepTime(points []Point) []Point {
    for i := 0; i < len(points); i++ {
        point := points[i]
        for j := 0; j < len(points); j++ {
            otherPoint := points[j]
            if point.x < otherPoint.x {
                point.velocity.x += 1
                otherPoint.velocity.x -= 1
            } else if point.x > otherPoint.x {
                point.velocity.x -= 1
                otherPoint.velocity.x += 1
            }

            if point.y < otherPoint.y {
                point.velocity.y += 1
                otherPoint.velocity.y -= 1
            } else if point.y > otherPoint.y {
                point.velocity.y -= 1
                otherPoint.velocity.y += 1
            }
            
            if point.z < otherPoint.z {
                point.velocity.z += 1
                otherPoint.velocity.z -= 1
            } else if point.z > otherPoint.z {
                point.velocity.z -= 1
                otherPoint.velocity.z += 1
            }
            points[j] = otherPoint
        }
    }

    for i := 0; i < len(points); i++ {
        point := points[i]
        point.x += point.velocity.x
        point.y += point.velocity.y
        point.z += point.velocity.z
        points[i] = point
    }

    return points
}

func toString(points []Point) string {
    sum := md5.Sum([]byte(fmt.Sprintf("%+v", points)))
    return fmt.Sprintf("%x", sum)
}

func getEnergy(point Point) int {
    return getPotential(point) * getKinetic(point)
}

func getPotential(point Point) int {
    return Abs(point.x) + Abs(point.y) + Abs(point.z)
}

func getKinetic(point Point) int {
    return Abs(point.velocity.x) + Abs(point.velocity.y) + Abs(point.velocity.z)
}   

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}