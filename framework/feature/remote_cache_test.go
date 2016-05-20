package feature
import (
    "fmt"
    "testing"
)

func TestRCacheLoadMget(t *testing.T){
    fmt.Println("remote-mget begin")
    rc := NewRCache("127.0.0.1:11211",1000000)
    rc.load("key1","key2","key3","key4","key5")
    fmt.Println(rc.mget("key1","key2","key3"))
    fmt.Println("remote-mget end")
}

func TestRCacheLoadGet(t *testing.T){
    fmt.Println("remote-get begin")
    rc := NewRCache("127.0.0.1:11211",1000000)
    rc.load("key1","key2","key3","key4","key5")
    fmt.Println(rc.get("key1"))
    fmt.Println("remote-get end")
}

