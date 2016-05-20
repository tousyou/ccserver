package feature
import (
    "github.com/youtube/vitess/go/memcache"
    "github.com/youtube/vitess/go/pools"
    "golang.org/x/net/context"
    "fmt"
    "time"
)

type ResourceConnMc struct {
    memcache.Connection
}

func (r ResourceConnMc) Close() {
    r.Connection.Close()
}

type RCache struct{
    addr        string
    timeout     time.Duration
    pool        *pools.ResourcePool
    Cache       cache
}

func NewRCache(addr string, timeout time.Duration) *RCache{
    p := pools.NewResourcePool(func() (pools.Resource, error) {
        c, err := memcache.Connect(addr, timeout)
        if err != nil {
            fmt.Println(err)
        }
        return ResourceConnMc{*c}, err
        }, 10, 20, time.Minute)
    return &RCache{
        addr:       addr,
        timeout:    timeout,
        pool:       p,
    }
}

func (db *RCache)load(keys ...string)bool{
    ctx := context.TODO()
    r, err := db.pool.Get(ctx)
    if err != nil {
        //fmt.Println(err)
        return false
    }
    defer db.pool.Put(r)
    c := r.(ResourceConnMc)
    results, err := c.Get(keys...)
    if err != nil {
        //fmt.Println(err)
        return false
    }
    for i,_ := range results{
        db.Cache.add(results[i].Key,ByteView{b: results[i].Value})
    }
    return true
}

func (db *RCache)mget(keys ...string)([]string){
    rets := make([]string,len(keys))
    for i,_ := range keys{
        val,bl := db.Cache.get(keys[i])
        if bl{
            rets[i] = val.String()
        }else{
            rets[i] = ""
        }
    }
    return rets
}

func (db *RCache)get(key string)(string,bool){
    val,bl := db.Cache.get(key)
    return val.String(),bl
}

