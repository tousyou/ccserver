package matchserver

import (
    "fmt"
    "net"
    "os"
)

type MatchServer struct {
    opts         *Options
    httpListener net.Listener
    db           PeerDB
}

func New(opts *Options) *MatchServer {
    n := &MatchServer{
        opts: opts,
    }
    n.logf("MatchServer")
    return n

}

func (l *MatchServer) logf(f string, args ...interface{}) {
    if l.opts.Log == nil {
        return
    }
    l.opts.Log.Output(2, fmt.Sprintf(f, args...))
}

func (l *MatchServer) Main() {
    fmt.Println("MatchServer Starting ...")
    ctx := &Context{l}
    //db,err := NewPeerMysql("localhost",3306,"root","12321")
    //db := NewPeerRedis("tcp","127.0.0.1:7777")
    db := NewPeerMc("127.0.0.1:8888",1000000)
    if db == nil {
        l.logf("FATAL: mysql connect failed")
        os.Exit(1)
    }
    l.db = db
    httpListener, err := net.Listen("tcp", l.opts.HTTPAddress)
    if err != nil {
        l.logf("FATAL: listen (%s) failed - %s", l.opts.HTTPAddress, err)
        os.Exit(1)
    }
    l.httpListener = httpListener
    httpServer := newHTTPServer(ctx)
    Serve(l.httpListener, httpServer, "HTTP", l.opts.Log)
}

func (m *MatchServer) Exit() {
    fmt.Println("MatchServer exiting ...")
}
