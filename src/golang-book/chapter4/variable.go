package main

import "fmt"

var dogName string = "Max"

func main() {
    x := "hello world"
    fmt.Printf("Go compiler is able to infer the type based on the literal value you assign the variable\n")
    fmt.Printf("x=%s \n", x);
    
    fmt.Printf("dogName = %s\n", dogName)
}
