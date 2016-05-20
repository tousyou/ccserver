package ifc
import (
    "../data"
)

//the funcation of target interface must be thread safe
type ITarget interface{
    Score(data.UserInfo,/*Feature,*/data.AdList)int32
    Sort(data.AdList)int32
}
