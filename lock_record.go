package ttlock

type LockRecordListResp struct {
	Err
	List []*LockRecordItem `json:"list"`
	PageNo int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

type LockRecordItem struct {
	LockId int `json:"lockId"`
	RecordType int `json:"recordType"`
	Success int `json:"success"`
	Username string `json:"username"`
	KeyboardPwd string `json:"keyboardPwd"`
	LockDate int64 `json:"lockDate"`
	ServerDate int64 `json:"serverDate"`
}

//获取开锁记录列表
//https://open.ttlock.com/doc/api/v3/lockRecord/list
func (s *Service) LockRecordList(req *ReqParams) (resp *LockRecordListResp) {
	var (
		u   = V3_URL + "/lockRecord/list"
		err error
	)
	resp = &LockRecordListResp{}
	err = req.CheckKeys("accessToken", "lockId","pageNo","pageSize")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//上传锁记录
//https://open.ttlock.com/doc/api/v3/lockRecord/upload
func (s *Service) LockRecordUpload(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lockRecord/upload"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "lockId","records")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}
