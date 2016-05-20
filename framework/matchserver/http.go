package matchserver

import (
    "fmt"
	"net/http"
	"net/http/pprof"
    "strconv"
	"github.com/julienschmidt/httprouter"
)

type httpServer struct {
	ctx    *Context
	router http.Handler
}

func newHTTPServer(ctx *Context) *httpServer {
	log := Log(ctx.matchserver.opts.Log)

	router := httprouter.New()
	router.HandleMethodNotAllowed = true
	router.PanicHandler = LogPanicHandler(ctx.matchserver.opts.Log)
	router.NotFound = LogNotFoundHandler(ctx.matchserver.opts.Log)
	router.MethodNotAllowed = LogMethodNotAllowedHandler(ctx.matchserver.opts.Log)
	s := &httpServer{
		ctx:    ctx,
		router: router,
	}

	router.Handle("GET", "/ping", Decorate(s.pingHandler, log, PlainText))
	router.Handle("GET", "/addpeer", Decorate(s.addPeer, log, PlainText))

	// debug
	router.HandlerFunc("GET", "/debug/pprof", pprof.Index)
	router.HandlerFunc("GET", "/debug/pprof/cmdline", pprof.Cmdline)
	router.HandlerFunc("GET", "/debug/pprof/symbol", pprof.Symbol)
	router.HandlerFunc("POST", "/debug/pprof/symbol", pprof.Symbol)
	router.HandlerFunc("GET", "/debug/pprof/profile", pprof.Profile)
	router.Handler("GET", "/debug/pprof/heap", pprof.Handler("heap"))
	router.Handler("GET", "/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handler("GET", "/debug/pprof/block", pprof.Handler("block"))
	router.Handler("GET", "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	return s
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *httpServer) pingHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	return "OK", nil
}
func (s *httpServer) addPeer(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
    fmt.Println("start addPeer")
    reqParams, err := NewReqParams(req)
    if err != nil {
        return nil, Err{400, "INVALID_REQUEST"}
    }
    host, err := reqParams.Get("host")
    if err != nil {
        return nil, Err{400, "MISSING_ARG_HOST"}
    }
    port, err := reqParams.Get("port")
    if err != nil {
        return nil, Err{400, "MISSING_ARG_PORT"}
    }

    port_int,_ := strconv.Atoi(port)
    p := &PeerInfo{
        Host:    host,
        Port:    port_int,
    }
    bl := s.ctx.matchserver.db.Add(p)
    if bl {
        return "OK",nil
    }else{
        return "NG",nil
    }
}

