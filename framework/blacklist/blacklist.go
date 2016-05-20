package blacklist
import (
    "../util"
)

type BlackList interface{
    Insert(key string)
    Erase(key string)bool
    Find(key string)bool
    Size()int
    Clear()
}

type blacklist struct{
    data    util.Set
}

func (b *blacklist)Insert(key string){
    b.data.Add(key)
}

func (b *blacklist)Erase(key string)bool{
    return b.data.Remove(key)
}

func (b blacklist)Find(key string)bool{
    return b.data.Contains(key)
}

func (b blacklist)Size()int{
    return b.data.Len()
}

func (b *blacklist)Clear(){
    b.data = util.NewSet()
}


