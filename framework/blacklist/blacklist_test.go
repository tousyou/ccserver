package blacklist
import (
    "testing"
    "../util"
)

func TestBlacklist(t *testing.T){
    black := &blacklist{
        data:     util.NewSet(),
    }
    util.Assert(t,black.Size(),0)
    util.Assert(t,black.Find("1"),false)
    black.Insert("1")
    util.Assert(t,black.Find("1"),true)
    util.Assert(t,black.Size(),1)
    black.Insert("2")
    util.Assert(t,black.Size(),2)
    black.Insert("3")
    black.Insert("4")
    black.Insert("5")
    util.Assert(t,black.Erase("6"),false)
    util.Assert(t,black.Find("3"),true)
    util.Assert(t,black.Erase("3"),true)
    util.Assert(t,black.Find("3"),false)
    util.Assert(t,black.Size(),4)
    util.Assert(t,black.Find("4"),true)
    black.Clear()
    util.Assert(t,black.Find("4"),false)
    util.Assert(t,black.Size(),0)
}
