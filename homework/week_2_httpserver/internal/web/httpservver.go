package web

import (
	"net/http"
	"sync"
)

type Server interface {
	Start(address string) error
	Route(pattern string ,handFunc http.HandlerFunc)
	Shutdown()
}

type httpservr struct {
	name string
	status int
	mux sync.Mutex
}

func (h *httpservr) Shutdown() {
	_ = h.SetStatus(http.StatusInternalServerError)
}

func (h *httpservr) Start(address string) error {
	return http.ListenAndServe(address,nil)
}

func (h *httpservr) Route(pattern string, handFunc http.HandlerFunc) {
	http.HandleFunc(pattern,handFunc)
}

func (h *httpservr) GetName() string {
	return h.name
}
func (h *httpservr) GetStatus() int {
	return  h.status
}
func (h *httpservr) SetStatus(status int) error {
	h.status = status
	return nil
}

func NewHttpServer(name string) Server{
	return &httpservr{
		name: name,
		status: http.StatusOK,
		mux: sync.Mutex{},
	}
}


