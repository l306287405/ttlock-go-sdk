package ttlock

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	URL = "https://api.ttlock.com"
	V3_URL = "https://api.ttlock.com/v3"
)

const (
	SUCCESS = 200
)

//配置
type Config struct {
	ClientId      string
	ClientSecret  string
}

func (s *Config) check() error {
	if s.ClientId == "" {
		return errors.New("ClientId is invalid")
	}
	if s.ClientSecret == "" {
		return errors.New("ClientSecret is invalid")
	}
	return nil
}

//服务
type Service struct {
	Config
}

func NewService(cfg Config) (*Service, error) {
	err := cfg.check()
	if err != nil {
		return nil, err
	}
	return &Service{cfg}, nil
}

//参数类型
type ReqParams struct {
	url.Values
}

func NewReqParams() *ReqParams {
	return &ReqParams{url.Values{}}
}

func (s *ReqParams) SetInt(key string, val int) {
	s.Set(key, strconv.Itoa(val))
}

func (s *ReqParams) SetList(key string, vals interface{}) {
	str, _ := json.Marshal(vals)
	s.Set(key, string(str))
}

func (s *ReqParams) SetMap(key string, mapVal interface{}) {
	str, _ := json.Marshal(mapVal)
	s.Set(key, string(str))
}

func (s *ReqParams) SetLong(key string, val int64) {
	s.SetInt64(key, val)
}

func (s *ReqParams) SetInt64(key string, val int64) {
	s.Set(key, strconv.FormatInt(val, 10))
}

func (s *ReqParams) SetInt8(key string, val int8) {
	s.Set(key, strconv.Itoa(int(val)))
}

func (s *ReqParams) SetFloat(key string, val float32) {
	s.SetFloat64(key,float64(val))
}

func (s *ReqParams) SetFloat64(key string, val float64) {
	s.Set(key, strconv.FormatFloat(val, 'G', -1, 64))
}

func (s *ReqParams) SetDouble(key string, val float64) {
	s.SetFloat64(key,val)
}

//获取参数签名
func (s *ReqParams) Sign(secret string) {
	var (
		keys      []string
		signStr   = secret
		k         string
		md5Ctx    = md5.New()
		cipherStr []byte
	)

	for k = range s.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k = range keys {
		if s.Get(k) != "" {
			signStr += k + s.Get(k)
		}
	}
	signStr += secret
	md5Ctx.Write([]byte(signStr))
	cipherStr = md5Ctx.Sum(nil)
	s.Set("sign", strings.ToLower(hex.EncodeToString(cipherStr)))
}

func (s *ReqParams) AddClientIdAndSecret(cfg *Config) {
	s.Set("clientId", cfg.ClientId)
	s.Set("clientSecret", cfg.ClientSecret)
	s.SetInt64("date", time.Now().UnixNano())
}

func (s *ReqParams) AddPublicParams(cfg *Config) {
	s.Set("clientId", cfg.ClientId)
	s.SetInt64("date", time.Now().UnixNano())
}

func (s *ReqParams) AddClientIdAndSecretSnake(cfg *Config) {
	s.Set("client_id", cfg.ClientId)
	s.Set("client_secret", cfg.ClientSecret)
	s.SetInt64("date", time.Now().UnixNano())
}

func (s *ReqParams) AddPublicParamsSnake(cfg *Config) {
	s.Set("client_id", cfg.ClientId)
	s.SetInt64("date", time.Now().UnixNano())
}

func (s *ReqParams) CheckKeys(keys ...string) error {
	if s == nil {
		return errors.New("Param is not init")
	}
	for _, key := range keys {
		if s.Get(key) == "" {
			return errors.New("Param is not found:" + key)
		}
	}
	return nil
}

func (s *ReqParams) ChooseOne(keys []string) error {
	for _, key := range keys {
		if s.Get(key) != "" {
			return nil
		}
	}
	return errors.New("Params must at least one: " + strings.Join(keys, ", "))

}
