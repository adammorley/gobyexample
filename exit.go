package main

import (
        "fmt"
        "os"
)

func main() {
    defer fmt.Println("!")
    fmt.Println("hi")

    os.Exit(3)
}
