package main

import (
    "fmt"
    "math/rand"
    "time"

)


func boring(msg string) <-chan string {
    c := make(chan string)
    go func() {
        for i := 0; ; i++ {
            c <- fmt.Sprintf("%s %d", msg, i)
            time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

        }
    }()
    return c
}

func main() {
    c1 := boring("Joe")
    c2 := boring("Bob")
    t1 := time.NewTimer(time.Second * 1)
    t := time.Now()
    fmt.Println(t)
    i := 0
    for {
        select {
        case s1 := <-c1:
            fmt.Println(s1)
            i++
            if i == 5 {
                fmt.Println("i=%d",i)
                return
            }
        case s2 := <-c2:
            fmt.Println(s2)
            i++
            if i == 5 {
                fmt.Println("i=%d",i)
                return
            }
        case <-t1.C:
            fmt.Println("You're too slow.")
            fmt.Println(time.Now())
            return
        }
    }
}

