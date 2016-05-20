package main
import (
    "fmt"
    "../feature"
    "time"
)

func  GetKey(cnt int) {
    feat := feature.NewFeature()
    j := 0
    i := 0
    for i<cnt {
        if feat.LoadRemoteKey("key1","key2","key3","key4","key5") {
            j++
        }
        i++
    }
    fmt.Printf("get success cnt = %d\n",j)
}

func  GetKeyByName(cnt int, con int) <- chan string{
    c := make(chan string)
    k := 0
    for k< con {
        go func(){
            b := k
            feat := feature.NewFeature()
            j := 0
            i := 0
            for i<cnt/con {
                if feat.LoadRemoteKey("key1","key2","key3","key4","key5") {
                    j++
                }
                i++
            }
            c <- fmt.Sprintf("%d: get success cnt = %d\n",b,j)
        }()
        k++
    }
    return c
}
func Single(cnt int) {
    fmt.Println(time.Now())
    GetKey(cnt)
    fmt.Println(time.Now())
}

func Conn(cnt int, con int){
    c := GetKeyByName(cnt,con)
    t := time.NewTimer(time.Second * 1000)
    fmt.Println(time.Now())
    i := 0
    for {
        select{
        case s := <- c:
            fmt.Println(s)
            i++
            if i == con { goto END }
        case <- t.C:
            fmt.Println("end")
            return
        }
    }
    END:
    fmt.Println(time.Now())
}

func main() {
    Conn(100000,2)
}
