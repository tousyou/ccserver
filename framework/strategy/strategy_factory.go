package strategy
import (
    "../ifc"
)
type StrategyFactory struct{}
func (s *StrategyFactory)CreateStrategy(stype string)ifc.IStrategy{
    switch(stype){
    case "SimpleHash":
        return new(SimpleHash)
    default:
        return new(SimpleHash)
    }
}
func (s *StrategyFactory)CreateTarget(stype string)ifc.ITarget{ return nil }
func (s *StrategyFactory)CreateRank(stype string)ifc.IRank{ return nil }

