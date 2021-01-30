package ttlock

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func PostForm(url string, R *ReqParams, resp interface{}) error {
	r, err := http.PostForm(url, R.Values)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return errors.New("通通锁请求http状态非200:" + r.Status)
	}

	return json.NewDecoder(r.Body).Decode(resp)
}

func GetRequest(url string, req *ReqParams, response interface{}) error {
	if !strings.HasSuffix(url, "?") {
		url += "?"
	}
	r, err := http.Get(url + req.Encode())
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return errors.New("通通锁请求http状态非200:" + r.Status + " url:" + url + " values:" + req.Encode())
	}
	return json.NewDecoder(r.Body).Decode(response)
}

func RandStr(len int) string {
	rand.Seed(time.Now().Unix())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
