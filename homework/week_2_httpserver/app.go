package main

import (
	"fmt"
	"kcamp/homework/week_2_httpserver/internal/biz"
	"kcamp/homework/week_2_httpserver/internal/web"
	"net/http"
	"net/http/pprof"
)
type RunFunc func(<-chan struct{}) error

type AppInfo struct {
	fns []RunFunc
}

func (a *AppInfo) Add(fn RunFunc)  {
	a.fns = append(a.fns, fn)
}

func (a *AppInfo) Run() error {
	if len(a.fns) == 0{
		return nil
	}
	stop := make(chan struct{})
	done := make(chan error,len(a.fns))
	defer close(done)
	for _,fn := range a.fns{
		go func(fn RunFunc) {
			done <- fn(stop)
		}(fn)
	}
	var err error
	for i := 0; i < cap(done); i++ {
		if err == nil{
			err = <- done
		}else {
			<- done
		}
		if i == 0{
			close(stop)
		}
	}
	return err
}

func serveApp( stop <- chan struct{}) error{
	httpServer := web.NewHttpServer("HttpServer")
	httpServer.Route("/home",biz.Home)
	httpServer.Route("/version",biz.GetVersion)
	httpServer.Route("/healthz",biz.Healthz)
	fmt.Println("Starting http server...")
	if err := httpServer.Start("127.0.0.1:8000");err != nil{
		return err
	}
	<- stop
	return nil
}

func serverDebug(stop <- chan struct{}) error{
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	server := &http.Server{
		Addr: "127.0.0.1:8001",
		Handler: mux,
	}
	fmt.Println("Starting http debug server...")
	if err := server.ListenAndServe();err != nil{
		return err
	}
	<- stop
	return nil
}