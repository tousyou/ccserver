package feature

type Feature struct{
    r_cache      *RCache
    l_cache      *LCache
}

func NewFeature() *Feature{
    return &Feature{
        r_cache:     NewRCache("127.0.0.1:11211",1000000),
        l_cache:     NewLCache("/Users/lids/Dev/ccserver/framework/test",100),
    }
}

func (f *Feature)LoadRemoteKey(keys ...string)bool{
    return f.r_cache.load(keys...)
}
func (f *Feature)LoadLocalKey()bool{
    return f.l_cache.load()
}
func (f *Feature)Get(key string)(string,bool){
    val,bl := f.l_cache.get(key)
    if !bl{
        val, bl = f.r_cache.get(key)
    }
    return val,bl
}
func (f *Feature)Mget(keys ...string)([]string){
    vals := make([]string,len(keys))
    for i,_ := range keys{
        val,_ := f.Get(keys[i])
        vals[i] = val
    }
    return vals
}

