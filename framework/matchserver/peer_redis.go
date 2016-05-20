package matchserver
import (
    "github.com/garyburd/redigo/redis"
    "github.com/youtube/vitess/go/pools"
    "golang.org/x/net/context"
    "fmt"
    "time"
)

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceConn struct {
    redis.Conn
}

func (r ResourceConn) Close() {
    r.Conn.Close()
}

type PeerRedis struct{
    Host     string
    Proto    string
    pool     *pools.ResourcePool
}

func NewPeerRedis(proto string, host string) *PeerRedis{
    p := pools.NewResourcePool(func() (pools.Resource, error) {
        c, err := redis.Dial(proto, host)
        return ResourceConn{c}, err
        }, 1, 2, time.Minute)
    return &PeerRedis{
        Host:   host,
        Proto:  proto,
        pool:   p,
    }
}

func (db *PeerRedis)Add(p *PeerInfo)bool{
    ctx := context.TODO()
    r, err := db.pool.Get(ctx)
    if err != nil {
        fmt.Println(err)
    }
    defer db.pool.Put(r)
    c := r.(ResourceConn)
    key := fmt.Sprintf("%s@%d",p.Host,p.Port)
    _, err = c.Do("SET",key,"1")
    if err != nil {
        fmt.Println(err)
    }
    return true
}
func (db *PeerRedis)Del(p *PeerInfo)bool{
    return true
}
func (db *PeerRedis)GetAll()Peers{
    return nil
}
