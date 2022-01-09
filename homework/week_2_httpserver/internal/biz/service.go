package biz

import (
	"fmt"
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	fmt.Printf("Client IP: %s,HTTP Code: %d\n",r.RemoteAddr,http.StatusOK)
	for key,values := range headers{
		for _,value := range values{
			w.Header().Add(key,value)
		}
	}
	w.WriteHeader(http.StatusOK)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	for key,values := range headers{
		for _,value := range values{
			w.Header().Add(key,value)
		}
	}
	version := os.Getenv("Version")
	w.Header().Add("Version",version)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Client IP: %s,HTTP Code: %d,Writer Version: %s\n",r.RemoteAddr,http.StatusOK,version)


}

func Healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Client IP: %s,HTTP Code: %d",r.RemoteAddr,http.StatusOK)
	w.WriteHeader(http.StatusOK)
}
