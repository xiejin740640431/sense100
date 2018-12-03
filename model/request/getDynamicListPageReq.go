package request

type GetDynamicListPageReq struct {
	BasePageReq
	ProgramId int64 `json:"programId" description:"小程序Id"`
	CompanyId int64 `json:"companyId" description:"企业Id"`
	FacadeId  int64 `json:"facadeId" description:"门面id"`
	UserId    int64 `json:"userId" description:"用户Id，没有可以不传"`
}

func (req GetDynamicListPageReq) Check() bool {
	if req.CompanyId == 0 && req.ProgramId == 0 && req.FacadeId == 0 {
		return false
	}
	return true
}
