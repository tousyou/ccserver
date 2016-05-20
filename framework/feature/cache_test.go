package feature
import(
    "fmt"
    "testing"
    "./lru"
)

func TestCache(t *testing.T){
    fmt.Println("cache-begin")
    ca := &cache{
        nbytes:    300000,
        lru:       lru.New(0),
    }
    value := ByteView{s:"hehe"}
    ca.add("key",value)
    if val,ok := ca.get("key"); ok{
        fmt.Println(val)
    }else{
        fmt.Println("no found")
    }
    fmt.Println("cache-end")
}
