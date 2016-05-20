package ifc
type AbstractFactory interface{
    CreateStrategy(string)IStrategy
    CreateTarget(string)ITarget
    CreateRank(string)IRank
}
