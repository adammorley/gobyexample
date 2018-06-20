package main

import (
        "encoding/json"
        "fmt"
        "os"
)

type response1 struct {
    Page int
    Fruits []string
}
type response2 struct {
    Page    int         `json:"page"`
    Fruits  []string   `json:"fruits"`    
}

var p = fmt.Println

func main() {
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    p(string(fltB))

    strB, _ := json.Marshal("gopher")
    p(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    p(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    p(string(mapB))

    res1D := &response1{
                Page:   1,
                Fruits: []string{"apple", "peach", "pear"},
             }
    res1B, _ := json.Marshal(res1D)
    p(string(res1B))

    res2D := &response2{
                Page:   1,
                Fruits: []string{"apple", "peach", "pear"},
             }
    res2B, _ := json.Marshal(res2D)
    p(string(res2B))

    bytes := []byte(`{"num":6.13,"strs":["a","b"]}`)
    var dat map[string]interface{}

    if err := json.Unmarshal(bytes, &dat); err != nil {
        panic(err)
    }
    p(dat)

    n := dat["num"]
    fmt.Printf("%T\n", n)
    p(dat["num"])

    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    p(str1)

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    p(res)
    p(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
