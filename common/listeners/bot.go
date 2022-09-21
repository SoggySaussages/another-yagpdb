package listeners

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Devtest Query returned")
	res.Write(data)
}

func RegisterPlugin() {
	logrus.Info("Starting devtest listener")
	go startHTTPServer()
}

func startHTTPServer() {
	handler := HttpHandler{}
	http.ListenAndServe(":9000", handler)
}