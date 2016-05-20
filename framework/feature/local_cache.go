package feature
import (
    "fmt"
    "os"
    "io"
    "bufio"
    "path/filepath"
    "strings"
)


// get file modified time
func FileMTime(file string) (int64, error) {
    f, e := os.Stat(file)
    if e != nil {
        return 0, e
    }
    return f.ModTime().Unix(), nil
}

type LCache struct{
    path        string
    times       map[string]int64
    interval    int64
    Cache       cache
}

func NewLCache(path string,interval int64) *LCache{
    return &LCache{
        path:      path,
        times:     make(map[string]int64),
        interval:  interval,
    }
}

func (lc *LCache)load()bool{
    err := filepath.Walk(lc.path, func(path string, f os.FileInfo, err error) error {
        if ( f == nil  ) {return err}
        if f.IsDir() {return nil}
        lc.loadfile(path)
        return nil
    })
    if err != nil {
        fmt.Printf("filepath.Walk() returned %v\n", err)
        return false
    }
    return true
}

func (lc *LCache)loadfile(filename string)error{
    //fmt.Println(filename)
    mtime,err := FileMTime(filename)
    if err != nil {
        fmt.Println(err)
        return err
    }
    last_time,ok := lc.times[filename]
    if ok {
        if (mtime - last_time) < lc.interval {
            return nil
        }
    }else{
        lc.times[filename]=mtime
    }
    _,key1 := filepath.Split(filename)
    f, err := os.OpenFile(filename,os.O_RDONLY,0660)
    defer f.Close()
    if err != nil {
        fmt.Println(err)
        return err
    }
    r := bufio.NewReader(f)
    line, err := r.ReadString('\n')
    for err == nil {
        item := strings.Trim(line,"\n")
        //fmt.Println(item)
        subs := strings.Split(item," ")
        if len(subs) == 2 {
            key := key1 + "@" + subs[0]
            //fmt.Println(key)
            bv := ByteView{s: subs[1]}
            lc.Cache.add(key,bv)
        }
        line, err = r.ReadString('\n')
    }
    if err != io.EOF {
        fmt.Println(err)
        return err
    }
    return err
}

func (db *LCache)mget(keys ...string)([]string){
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

func (db *LCache)get(key string)(string,bool){
    val,bl := db.Cache.get(key)
    return val.String(),bl
}

