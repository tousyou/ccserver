package matchserver
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

type PeerMc struct{
    Addr     string
    Timeout  time.Duration
    pool     *pools.ResourcePool
}

func NewPeerMc(addr string, timeout time.Duration) *PeerMc{
    p := pools.NewResourcePool(func() (pools.Resource, error) {
        c, err := memcache.Connect(addr, timeout)
        if err != nil {
            fmt.Println(err)
        }
        return ResourceConnMc{*c}, err
        }, 1, 2, time.Minute)
    return &PeerMc{
        Addr:   addr,
        Timeout:  timeout,
        pool:   p,
    }
}

func (db *PeerMc)Add(p *PeerInfo)bool{
    ctx := context.TODO()
    r, err := db.pool.Get(ctx)
    if err != nil {
        fmt.Println(err)
    }
    defer db.pool.Put(r)
    c := r.(ResourceConnMc)
    key := fmt.Sprintf("%s@%d",p.Host,p.Port)
    _, err = c.Set(key,0,0,[]byte("1"))
    if err != nil {
        fmt.Println(err)
    }
    return true
}
func (db *PeerMc)Del(p *PeerInfo)bool{
    return true
}
func (db *PeerMc)GetAll()Peers{
    return nil
}
