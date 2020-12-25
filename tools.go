package ttlock

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

func PostForm(url string, R *ReqParams, resp interface{}) error {
	r, err := http.PostForm(url, R.Values)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return errors.New("点评请求http状态非200:" + r.Status)
	}

	return json.NewDecoder(r.Body).Decode(resp)
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
