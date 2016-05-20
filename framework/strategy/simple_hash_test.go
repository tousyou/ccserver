package strategy
import (
    "testing"
    "fmt"
)

func Test_LoadCfg(t *testing.T){
    sh := SimpleHash{}
    sh.LoadCfg("30:t1;70:t2")
    fmt.Println(sh.SelectPlugin("0"))
    fmt.Println(sh.SelectPlugin("1"))
    fmt.Println(sh.SelectPlugin("29"))
    fmt.Println(sh.SelectPlugin("30"))
    fmt.Println(sh.SelectPlugin("31"))
    fmt.Println(sh.SelectPlugin("69"))
    fmt.Println(sh.SelectPlugin("70"))
    fmt.Println(sh.SelectPlugin("71"))
    fmt.Println(sh.SelectPlugin("90"))
}
