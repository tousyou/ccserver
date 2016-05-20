package abtest
import (
    "fmt"
    "testing"
)

func Test_Init(t *testing.T){
    ab := ABtest{}
    ab.Init("select.xml")
    targetname := new(string)
    rankname := new(string)
    ab.Select("1",targetname,rankname)
    fmt.Println(*targetname)
    fmt.Println(*rankname)
}
