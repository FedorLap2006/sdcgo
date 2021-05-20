package sdcgo

import (
	"encoding/json"
	"fmt"
)

type RESTError struct {
	Message string `json:"msg"`
	Type string `json:"type"`
	Code uint `json:"code"`
}

func NewRESTError(b []byte) (*RESTError, error) {
	var v struct {
		Error *RESTError `json:"error"`
	}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}

	return v.Error, nil
}

func (e RESTError) Error() string {
	return fmt.Sprintf("msg: %s code: %d", e.Message, e.Code)
}
