package main

import (
        "fmt"
        "time"
)

func main() {
    n := 5
    requests := make(chan int, n)
    for i := 1; i <= n; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    //burstyLimiter := make(chan time.Time, 3)
    burstyLimiter := make(chan int, 3)

    for i := 0; i < 3; i++ {
        burstyLimiter <- i
    }

    go func() {
        //for t := range time.Tick(200 * time.Millisecond) {
        for range time.Tick(200 * time.Millisecond) {
            //burstyLimiter <- t
            burstyLimiter <- 1
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}

/*
    key learning here: you can use a channel being filled by a background thread (goroutine) to "tick" and rate
    limit the foreground process.

    time.Tick is a key realization; it ticks.  as is rate limiter channel being able to be any value as long as it's range'd across with a for loop.
*/
