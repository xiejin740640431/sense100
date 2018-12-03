package request

type BasePageReq struct {
	PageIndex int  `json:"pageIndex"  description:"页面索引"`
	PageSize  int  `json:"pageSize" description:"页面数据个数"`
	IsPage    bool `json:"isPage" description:"是否分页"`
}
