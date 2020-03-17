package utils

import (
	"encoding/json"
	"net/http"

	"gitlab.com/aubayaml/aubayaml-go/bookstore/utils-go/errors"
)

var (
	//HTTP utils
	HTTP httpUtilsInteface = &httpUtils{}
)

type httpUtilsInteface interface {
	ResponseJSON(http.ResponseWriter, int, interface{})
	ResponseError(http.ResponseWriter, *errors.RestErr)
}

type httpUtils struct{}

//ResponseJSON in JSON
func (u *httpUtils) ResponseJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

//ResponseError error response SOLID
func (u *httpUtils) ResponseError(w http.ResponseWriter, err *errors.RestErr) {
	u.ResponseJSON(w, err.Status, err)
}
