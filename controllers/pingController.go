package controller

import "net/http"

var (
	//PingController visible ping controller
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

func (p *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
