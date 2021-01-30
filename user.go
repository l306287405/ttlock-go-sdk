package ttlock

type UserRegisterResp struct {
	Err
	Username string `json:"username"`
}

//新用户注册
//https://open.ttlock.com/doc/api/v3/user/register
func (s *Service) UserRegister(req *ReqParams) (resp *UserRegisterResp) {
	var (
		u   = V3_URL + "/user/register"
		err error
	)
	resp = &UserRegisterResp{}
	err = req.CheckKeys("username", "password")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddClientIdAndSecret(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//用户重置密码
//https://open.ttlock.com/doc/api/v3/user/resetPassword
func (s *Service) UserResetPassword(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/user/resetPassword"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("username", "password")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddClientIdAndSecret(&s.Config)

	resp.failed(PostForm(u, req, resp))
	return
}

type UserListResp struct {
	Err
	List []*struct {
		Userid  string `json:"userid"`
		Regtime int64  `json:"regtime"`
	} `json:"list"`
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Pages    int `json:"pages"`
	Total    int `json:"total"`
}

//获取用户列表
//https://api.ttlock.com/v3/user/list
func (s *Service) UserList(req *ReqParams) (resp *UserListResp) {
	var (
		u   = V3_URL + "/user/list"
		err error
	)
	resp = &UserListResp{}
	err = req.CheckKeys("startDate", "endDate", "pageNo", "pageSize")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddClientIdAndSecret(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//删除用户
//https://api.ttlock.com/v3/user/delete
func (s *Service) UserDelete(req *ReqParams) (resp *Err) {
	var (
		u   = V3_URL + "/user/delete"
		err error
	)
	resp = &Err{}
	err = req.CheckKeys("username")
	if err != nil {
		resp.failed(err)
		return
	}
	req.AddClientIdAndSecret(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}
