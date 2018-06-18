package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name: "Alice", age: 30})
    fmt.Println(person{name: "Fred"})
    fmt.Println(&person{name: "Ann", age: 40})
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
    sp := &s
    fmt.Println(sp.age)
    sp.age = 51
    fmt.Println(sp.age)
}

/*import "math"

type circle struct {
    r float64
}

func (c *circle) area() float64 {
    return math.Pow(c.r, 2) * math.Pi
}

func (c *circle) circumference() float64 {
    return 2 * math.Pi * c.r
}

func (c *circle) diameter() float64 {
    return 2 * c.r
}

func main() {
    var c circle
    c.r = 2
    fmt.Println(c.area())
} what I remember */
