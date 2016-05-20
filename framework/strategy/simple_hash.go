package strategy
import (
    "strings"
    "strconv"
)

type SimpleHash struct{
    plugins    map[int]string
}
//the format of config parameter
//example:      30:target1;10:target2;60:target3
func (s *SimpleHash)LoadCfg(config string){
    s.plugins = make(map[int]string)
    ps := make(map[int]string)
    targets := strings.Split(config,";")
    for _,target := range targets {
        kv := strings.Split(target,":")
        if len(kv) == 2 {
            key,err := strconv.Atoi(kv[0])
            if (err == nil) && (key > 0) && (key <= 100) {
                ps[key] = kv[1]
            }
        }
    }
    start := 0
    for partition,pluginname := range ps{
        start += partition
        s.plugins[start] = pluginname
    }
}
func (s *SimpleHash)SelectPlugin(uid string)string{
    uidint :=0
    if len(uid) >= 10 {
        uidint,_ = strconv.Atoi(uid[0:8])
    }else{
        uidint,_ = strconv.Atoi(uid)
    }
    part := uidint % 100
    for key,value := range s.plugins{
        if part <= key {
            return value
        }
    }
    return ""
}
