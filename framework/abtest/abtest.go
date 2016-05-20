package abtest
import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "../ifc"
)

type Strategy struct {
    Name string           `xml:"Strategy>Name"`
    Config  string        `xml:"Strategy>Config"`
    selectPlugin    ifc.IStrategy
}
type TargetLevel struct {
    Name string           `xml:"Name"`
    Strategy
}

type ABConfig struct {
    XMLName xml.Name      `xml:"ABtest"`
    Strategy
    Target  []TargetLevel
}

type ABtest struct {
    filename      string
    cfg           ABConfig
}
func (p *ABtest)Init(filename string)int{
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("error: %v", err)
        return 1
    }
    err = xml.Unmarshal([]byte(buf), &p.cfg)
    if err != nil {
        fmt.Printf("error: %v", err)
        return 2
    }

    fp := FactoryProducer{}
    sf := fp.GetFactory("Strategy")
    p.cfg.selectPlugin = sf.CreateStrategy(p.cfg.Name)
    p.cfg.selectPlugin.LoadCfg(p.cfg.Config)

    for i,value := range p.cfg.Target {
        p.cfg.Target[i].selectPlugin = sf.CreateStrategy(value.Strategy.Name)
        p.cfg.Target[i].selectPlugin.LoadCfg(value.Strategy.Config)
    }
    return 0
}

func (p *ABtest)Select(uid string,targetname *string, rankname *string){
    *targetname = p.cfg.selectPlugin.SelectPlugin(uid)
    for _,value := range p.cfg.Target {
        if value.Name == *targetname {
            *rankname = value.selectPlugin.SelectPlugin(uid)
        }
    }
}
