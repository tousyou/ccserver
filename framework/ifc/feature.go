package ifc

type IFeature interface{
    Get(key string)(string,bool)
    Mget(keys ...string)([]string)
}
