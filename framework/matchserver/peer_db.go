package matchserver

type PeerInfo struct {
    Host string  `json:"host"`
    Port int   `json:"port"`
    Desc string  `json:"desc"`
}

type Peers []*PeerInfo

type PeerDB interface{
    Add(p *PeerInfo)bool
    Del(p *PeerInfo)bool
    GetAll()Peers
}

