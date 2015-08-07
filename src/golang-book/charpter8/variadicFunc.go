package main

import "fmt"

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    fmt.Println(add(1,2,3))

    xs := [] int{1, 2, 3}
    
    fmt.Println("\nsum of xs = ")
    fmt.Println(add(xs...))

    // Closure:  create function inside of function
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println("add(1,1) = ")
    fmt.Println(add(1,1))
}

