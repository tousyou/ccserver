package match
import (
    "../blacklist"
    "../feature"
    //"../matchserver"
)

type Handle interface{
    //Handle(req * matchserver.ReqParams)string
    Handle(string)string
}

type match struct{
    black     blacklist.BlackList
    feat      feature.Feature
}

//func (m match)Handle(req * matchserver.ReqParams)string{
func (m match)Handle(req string)string{
    c := make(chan int)
    go func(){
        m.black.Find(req)
        c <- 1
    }
    go func(){
        m.feat.LoadRemoteKey(req)
        c <- 2
    }
    for {
        select {
        case s := <-c:
            if = 2 { break }
        case <-time.After(1 * time.Second):
            fmt.Println("You're too slow.")
            break
        }
    }
    return ""
}
