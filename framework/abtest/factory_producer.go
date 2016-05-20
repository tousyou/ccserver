package abtest
import (
    "../ifc"
    "../strategy"
)

type FactoryProducer struct{}
func (f *FactoryProducer)GetFactory(ptype string)ifc.AbstractFactory{
    switch(ptype){
    case "Strategy":
        return new(strategy.StrategyFactory)
    default:
        return nil
    }
}
