package main

import (
    "log"
    "crypto/md5"
    "fmt"
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

    initials := make([]Point, 4)
    copy(initials, points)

    ixs := make([]int, 0)
    iys := make([]int, 0)
    izs := make([]int, 0)

    ixs = append(ixs, points[0].x, points[1].x, points[2].x, points[3].x)
    ixs = append(ixs, points[0].velocity.x, points[1].velocity.x, points[2].velocity.x, points[3].velocity.x)
    iys = append(iys, points[0].y, points[1].y, points[2].y, points[3].y)
    iys = append(iys, points[0].velocity.y, points[1].velocity.y, points[2].velocity.y, points[3].velocity.y)
    izs = append(izs, points[0].z, points[1].z, points[2].z, points[3].z)
    izs = append(izs, points[0].velocity.z, points[1].velocity.z, points[2].velocity.z, points[3].velocity.z)

    xcycle := 0
    ycycle := 0
    zcycle := 0

    count := 1
    for {
        points = stepTime(points)

        xs := make([]int, 0)
        ys := make([]int, 0)
        zs := make([]int, 0)

        xs = append(xs, points[0].x, points[1].x, points[2].x, points[3].x)
        xs = append(xs, points[0].velocity.x, points[1].velocity.x, points[2].velocity.x, points[3].velocity.x)
        ys = append(ys, points[0].y, points[1].y, points[2].y, points[3].y)
        ys = append(ys, points[0].velocity.y, points[1].velocity.y, points[2].velocity.y, points[3].velocity.y)
        zs = append(zs, points[0].z, points[1].z, points[2].z, points[3].z)
        zs = append(zs, points[0].velocity.z, points[1].velocity.z, points[2].velocity.z, points[3].velocity.z)

        if xcycle == 0 {
            if sliceEq(xs, ixs) {
                xcycle = count
            }
        }

        if ycycle == 0 {
            if sliceEq(ys, iys) {
                ycycle = count
            }
        }

        if zcycle == 0 {
            if sliceEq(zs, izs) {
                zcycle = count
            }
        }

        if xcycle > 0 && ycycle > 0 && zcycle > 0 {
            log.Println("Go online to calculate the LCM of these 3:", xcycle, ycycle, zcycle)
            break
        }

        count += 1
    }
}

func sliceEq(a, b []int) bool {
    if (a == nil) != (b == nil) { 
        return false; 
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
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