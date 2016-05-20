package ifc
import (
    "../data"
)

//the funcation of target interface must be thread safe
type IRank interface{
    Score(data.UserInfo,/*Feature,*/data.AdList)int32
    Sort(data.AdList)int32
    Bid(data.UserInfo,/*Feature,*/data.AdList)int32
}
