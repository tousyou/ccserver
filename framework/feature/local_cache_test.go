package feature
import (
    "fmt"
    "testing"
)

func TestLCLoadMget(t *testing.T){
    fmt.Println("local-mget begin")
    lc := NewLCache("/Users/lids/Dev/ccserver/framework/test",100)
    lc.load()
    results := lc.mget("u@1","a@1","f@1","c@1")
    fmt.Println(results)
    fmt.Println("local-mget end")
}

func TestLCLoadGet(t *testing.T){
    fmt.Println("local-get begin")
    lc := NewLCache("/Users/lids/Dev/ccserver/framework/test",100)
    lc.load()
    results,_ := lc.get("u@1")
    fmt.Println(results)
    fmt.Println("local-get end")
}
