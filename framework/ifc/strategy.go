package ifc
type IStrategy interface{
    LoadCfg(config string)
    SelectPlugin(uid string)string
}
