package ttlock

type Oauth2TokenResp struct {
	Err
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	Uid          int    `json:"uid,omitempty"`
}

//获取访问令牌
//https://open.ttlock.com/doc/oauth2
func (s *Service) AccessToken(req *ReqParams) (resp *Oauth2TokenResp) {
	var (
		u   = URL + "/oauth2/token"
		err error
	)
	resp = &Oauth2TokenResp{}
	err = req.CheckKeys("username", "password", "redirect_uri")
	if err != nil {
		resp.failed(err)
		return
	}
	req.Set("grant_type", "password")
	req.AddClientIdAndSecretSnake(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}

//刷新访问令牌
//https://open.ttlock.com/doc/oauth2/refreshToken
func (s *Service) RefreshToken(req *ReqParams) (resp *Oauth2TokenResp) {
	var (
		u   = URL + "/oauth2/token"
		err error
	)
	resp = &Oauth2TokenResp{}
	err = req.CheckKeys("refresh_token", "redirect_uri")
	if err != nil {
		resp.failed(err)
		return
	}
	req.Set("grant_type", "refresh_token")
	req.AddClientIdAndSecretSnake(&s.Config)
	resp.failed(PostForm(u, req, resp))
	return
}
