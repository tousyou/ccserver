package data

type AdInfo struct{
    adid                uint64  `json:"adid"`
    bid                 uint32  `json:"bid"`
    cost                uint32  `json:"cost"`
    target_score        uint32  `json:"taget_score"`
    rank_score          uint32  `json:"rank_score"`
}

type AdList []*AdInfo
