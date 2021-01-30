package ttlock

type LockInitResp struct {
	Err
	LockId int `json:"lockId"`
	KeyId  int `json:"keyId"`
}

//锁初始化
//https://open.ttlock.com/doc/api/v3/lock/initialize
func (s *Service) LockInit(req *ReqParams) (resp *LockInitResp) {
	var (
		u   = V3_URL + "/lock/initialize"
		err error
	)
	resp = &LockInitResp{}
	err = req.CheckKeys("accessToken", "lockData", "lockAlias")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

type LockListResp struct {
	Err
	List []*LockListItem `json:"list,omitempty"`
}
type LockListItem struct {
	LockId             int    `json:"lockId"`
	Date               int64  `json:"date"`
	LockName           string `json:"lockName"`
	LockAlias          string `json:"lockAlias"`
	LockMac            string `json:"lockMac"`
	ElectricQuantity   int    `json:"electricQuantity"`
	KeyboardPwdVersion int8   `json:"keyboardPwdVersion"`
	SpecialValue       int    `json:"specialValue"`
	HasGateway         int    `json:"hasGateway"`
	LockData           string `json:"lockData"`
}

//获取名下锁列表
//https://open.ttlock.com/doc/api/v3/lock/list
func (s *Service) LockList(req *ReqParams) (resp *LockListResp) {
	var (
		u   = V3_URL + "/lock/list"
		err error
	)
	resp = &LockListResp{}
	err = req.CheckKeys("accessToken", "pageNo", "pageSize")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(GetRequest(u, req, resp))
	return
}

type LockDetailResp struct {
	Err
	LockId             int          `json:"lockId"`
	LockName           string       `json:"lockName"`
	LockAlias          string       `json:"lockAlias"`
	LockMac            string       `json:"lockMac"`
	LockKey            string       `json:"lockKey"`
	LockFlagPos        int          `json:"lockFlagPos"`
	AdminPwd           string       `json:"adminPwd"`
	NoKeyPwd           int          `json:"noKeyPwd"`
	DeletePwd          string       `json:"deletePwd"`
	AesKeyStr          string       `json:"aesKeyStr"`
	LockVersion        *LockVersion `json:"lockVersion"`
	KeyboardPwdVersion int8         `json:"keyboardPwdVersion"`
	ElectricQuantity   int          `json:"electricQuantity"`
	SpecialValue       int          `json:"specialValue"`
	TimezoneRawOffset  int64        `json:"timezoneRawOffset"`
	ModelNum           string       `json:"modelNum"`
	HardwareRevision   string       `json:"hardwareRevision"`
	FirmwareRevision   string       `json:"firmwareRevision"`
	Date               int64        `json:"date"`
}

type LockVersion struct {
	ProtocolType    int `json:"protocolType"`
	ProtocolVersion int `json:"protocolVersion"`
	Scene           int `json:"scene"`
	GroupId         int `json:"groupId"`
	OrgId           int `json:"orgId"`
}

//获取锁详细信息
//https://open.ttlock.com/doc/api/v3/lock/detail
func (s *Service) LockDetail(req *ReqParams) (resp *LockDetailResp) {
	var (
		u   = V3_URL + "/lock/detail"
		err error
	)
	resp = &LockDetailResp{}
	err = req.CheckKeys("accessToken", "lockId")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(GetRequest(u, req, resp))
	return
}

type LockListKeyResp struct {
	Err
	List []*LockListKeyItem `json:"list,omitempty"`
}

type LockListKeyItem struct {
	KeyId          int    `json:"keyId"`
	LockId         int    `json:"lockId"`
	OpenId         int    `json:"openId"`
	Username       string `json:"username"`
	KeyName        string `json:"keyName"`
	KeyStatus      string `json:"keyStatus"`
	StartDate      int64  `json:"startDate"`
	EndDate        int64  `json:"endDate"`
	KeyRight       int8   `json:"keyRight"`
	SenderUsername string `json:"senderUsername"`
	Remarks        string `json:"remarks"`
	Date           int64  `json:"date"`
}

//获取锁的普通钥匙列表
//https://open.ttlock.com/doc/api/v3/lock/listKey
func (s *Service) LockListKey(req *ReqParams) (resp *LockListKeyResp) {
	var (
		u   = V3_URL + "/lock/listKey"
		err error
	)
	resp = &LockListKeyResp{}
	err = req.CheckKeys("accessToken", "lockId", "pageNo", "pageSize")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(GetRequest(u, req, resp))
	return
}

//修改锁名称
//https://open.ttlock.com/doc/api/v3/lock/rename
func (s *Service) LockRename(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lock/rename"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "lockId", "lockAlias")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//转移锁
//https://open.ttlock.com/doc/api/v3/lock/transfer
func (s *Service) LockTransfer(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lock/transfer"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "receiverUsername", "lockIdList")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//上传锁电量
//https://open.ttlock.com/doc/api/v3/lock/updateElectricQuantity
func (s *Service) LockUpdateElectricQuantity(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lock/updateElectricQuantity"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "lockId", "ElectricQuantity")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//更新锁的特征值
//https://open.ttlock.com/doc/api/v3/lock/updateSpecialValue
func (s *Service) LockUpdateSpecialValue(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lock/updateSpecialValue"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "lockId", "specialValue")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

type LockQueryElectricQuantityResp struct {
	Err
	ElectricQuantity int `json:"electricQuantity"`
}

//查询锁电量
//https://open.ttlock.com/doc/api/v3/lock/queryElectricQuantity
func (s *Service) LockQueryElectricQuantity(req *ReqParams) (resp *LockQueryElectricQuantityResp) {
	var (
		u   = V3_URL + "/lock/queryElectricQuantity"
		err error
	)
	resp = &LockQueryElectricQuantityResp{}
	err = req.CheckKeys("accessToken", "lockId")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(GetRequest(u, req, resp))
	return
}

//删除锁
//https://open.ttlock.com/doc/api/v3/lock/delete
func (s *Service) LockDelete(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lock/delete"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "lockId")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//设置自动闭锁时间
//https://open.ttlock.com/doc/api/v3/lock/setAutoLockTime
func (s *Service) LockSetAutoLockTime(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lock/setAutoLockTime"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "lockId", "seconds", "type")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//设置常开模式
//https://open.ttlock.com/doc/api/v3/lock/configPassageMode
func (s *Service) LockConfigPassageMode(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lock/configPassageMode"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "lockId", "passageMode", "startDate", "endDate", "isAllDay", "weekDays", "type")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

type LockGetPassageModeConfigResp struct {
	Err
	PassageMode int   `json:"passageMode"`
	StartDate   int   `json:"startDate"`
	EndDate     int   `json:"endDate"`
	IsAllDay    int   `json:"isAllDay"`
	WeekDays    []int `json:"weekDays"`
}

//获取锁的常开模式设置
//https://open.ttlock.com/doc/api/v3/lock/getPassageModeConfig
func (s *Service) LockGetPassageModeConfig(req *ReqParams) (resp *LockGetPassageModeConfigResp) {
	var (
		u   = V3_URL + "/lock/getPassageModeConfig"
		err error
	)
	resp = &LockGetPassageModeConfigResp{}
	err = req.CheckKeys("accessToken", "lockId")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(GetRequest(u, req, resp))
	return
}

//更新锁数据
//https://open.ttlock.com/doc/api/v3/lock/updateLockData
func (s *Service) LockUpdateLockData(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/lock/updateLockData"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("accessToken", "lockId", "lockData")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

type LockUpgradeResp struct {
	Err
	NeedUpgrade     int8   `json:"needUpgrade"`
	FirmwareInfo    string `json:"firmwareInfo"`
	FirmwarePackage string `json:"firmwarePackage"`
	Version         string `json:"version"`
}

//检测锁固件是否需要升级
//https://open.ttlock.com/doc/api/v3/lock/upgradeCheck
func (s *Service) LockUpgradeCheck(req *ReqParams) (resp *LockUpgradeResp) {
	var (
		u   = V3_URL + "/lock/upgradeCheck"
		err error
	)
	resp = &LockUpgradeResp{}
	err = req.CheckKeys("accessToken", "lockId")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(GetRequest(u, req, resp))
	return
}

//再次检测锁固件是否需要升级
//https://open.ttlock.com/doc/api/v3/lock/upgradeRecheck
func (s *Service) LockUpgradeRecheck(req *ReqParams) (resp *LockUpgradeResp) {
	var (
		u   = V3_URL + "/lock/upgradeRecheck"
		err error
	)
	resp = &LockUpgradeResp{}
	err = req.CheckKeys("accessToken", "lockId", "firmwareInfo")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddPublicParams(&s.Config)
	resp.failed(GetRequest(u, req, resp))
	return
}
