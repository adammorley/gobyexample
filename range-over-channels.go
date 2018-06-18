package main

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    select {
    case elem := <-queue:
        fmt.Println(elem)
    default:
        fmt.Println("default")
    }
    /*for elem := range queue {
        fmt.Println(elem)
    }*/
}
