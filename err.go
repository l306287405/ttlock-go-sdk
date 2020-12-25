package ttlock

import "strconv"

const (
	NONE_ERROR   = 0
	FAILED_ERROR = 1
)

type Err struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func (e *Err) String() string {
	return "Errmsg:" + e.Errmsg + " Errcode:" + strconv.Itoa(e.Errcode)
}

func (e *Err) failed(err error) {
	if err == nil {
		return
	}
	if e == nil {
		e = &Err{}
	}
	e.Errcode = FAILED_ERROR
	e.Errmsg = err.Error()
}

func (e *Err) Success() bool {
	if e.Errcode == NONE_ERROR {
		return true
	}
	return false
}
