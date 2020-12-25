package ttlock

type KeySendResp struct {
	Err
	KeyId int `json:"keyId"`
}

//发送钥匙
//https://open.ttlock.com/doc/api/v3/key/send
func (s *Service) KeySend(req *ReqParams) (resp *KeySendResp) {
	var (
		u   = V3_URL + "/key/send"
		err error
	)
	resp = &KeySendResp{}
	err = req.CheckKeys("accessToken", "lockId", "receiverUsername", "keyName",
		"startDate", "endDate", "remarks", "remoteEnable")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

type KeyListResp struct {
	Err
	List []*KeyDetail `json:"list,omitempty"`
}

type KeyDetail struct {
	KeyId              int          `json:"keyId"`
	LockData           string       `json:"lockData"`
	LockId             int          `json:"lockId"`
	UserType           string       `json:"userType"`
	KeyStatus          string       `json:"keyStatus"`
	LockName           string       `json:"lockName"`
	LockAlias          string       `json:"lockAlias"`
	LockMac            string       `json:"lockMac"`
	NoKeyPwd           string       `json:"noKeyPwd"`
	DeletePwd          string       `json:"deletePwd"`
	ElectricQuantity   int          `json:"electricQuantity"`
	LockVersion        *LockVersion `json:"lockVersion"`
	StartDate          int64        `json:"startDate"`
	EndDate            int64        `json:"endDate"`
	Remarks            string       `json:"remarks"`
	KeyRight           int8         `json:"keyRight"`
	KeyboardPwdVersion int8         `json:"keyboardPwdVersion"`
	SpecialValue       int          `json:"specialValue"`
	RemoteEnable       int8         `json:"remoteEnable"`
}

//发送钥匙 如果是微信小程序调用则IsWxApp传true
//https://open.ttlock.com/doc/api/v3/key/list
func (s *Service) KeyList(req *ReqParams,IsWxApp bool) (resp *KeyListResp) {
	var (
		u   = V3_URL + "/key/list"
		err error
	)
	resp = &KeyListResp{}
	err = req.CheckKeys("accessToken", "pageNo", "pageSize")
	if err != nil {
		resp.failed(err)
		return
	}
	if IsWxApp{
		req.SetInt("sdkVersion",2)
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

type KeyGetResp struct {
	Err
	KeyDetail
}

//获取单把钥匙
//https://open.ttlock.com/doc/api/v3/key/get
func (s *Service) KeyGet(req *ReqParams) (resp *KeyGetResp) {
	var (
		u   = V3_URL + "/key/get"
		err error
	)
	resp = &KeyGetResp{}
	err = req.CheckKeys("accessToken", "pageNo", "pageSize")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//删除电子钥匙
//https://open.ttlock.com/doc/api/v3/key/delete
func (s *Service) KeyDelete(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/key/delete"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "keyId")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//修改钥匙有效期
//https://open.ttlock.com/doc/api/v3/key/changePeriod
func (s *Service) KeyChangePeriod(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/key/changePeriod"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "keyId","startDate","endDate")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}
