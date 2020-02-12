package bean

//删除地点结构体
type DelAddr struct {
	AddrId   string `json:"addrId"`
	AddrType int    `json:"addrType"`
}

//地点结构体
type InsAddr struct {
	AddrId      string `json:"addrId"`
	Name        string `json:"name"`
	OrderNo     int32  `json:"orderNo"`
	IsLeaf      int32  `json:"isLeaf"`
	ParentId    string `json:"parentId"`
	Src         string `json:"src"`
	AddrType    int32  `json:"addrType"`
	SimpleName  string `json:"simpleName"`
	NickName    string `json:"nickName"`
	PushName    string `json:"pushName"`
	Description string `json:"description"`
	State       int32  `"state"`
}

//客户端传来的查询结构体
type QueAddr struct {
	//园区
	Garden string `json:"garden"`
	//数据中心
	Center string `json:"center"`
	//模组
	Model string `json:"model"`
	//楼层
	Floor string `json:"floor"`
	//房间
	Room string `json:"room"`
}

//地点参数
type AddrResp struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Addr *[]InsAddr `json:"addr"`
}
