package feature
import (
    "fmt"
    "testing"
)

func TestGet(t *testing.T){
    fmt.Println("feature-begin")
    feat := NewFeature()
    feat.LoadRemoteKey("key1","key2")
    fmt.Println(feat.Get("key1"))
    fmt.Printf("%q\n",feat.Mget("key2","key10"))
    feat.LoadLocalKey()
    fmt.Println(feat.Mget("u@1","a@1","f@1","c@1","key1","key2"))
    fmt.Println("feature-end")
}
