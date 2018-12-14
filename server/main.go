package main

import (
	"flag"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/golang/glog"
	"github.com/liuzl/goutil/rest"
)

var (
	addr = flag.String("addr", ":8080", "bind address")
	dir  = flag.String("dir", "data", "dir for files")
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("addr=%s  method=%s host=%s uri=%s",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	r.ParseForm()
	query := strings.TrimSpace(r.FormValue("query"))
	file := strings.TrimSpace(r.FormValue("file"))
	if query == "" || file == "" {
		rest.MustEncode(w, rest.RestMessage{"error", "query or file is empty"})
		return
	}
	file = filepath.Join(*dir, file)
	go run(query, file)
	rest.MustEncode(w, rest.RestMessage{"ok", "job submitted"})
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.Handle("/api/", rest.WithLog(ApiHandler))
	http.Handle("/", http.FileServer(rice.MustFindBox("ui").HTTPBox()))
	glog.Info("server listen on", *addr)
	glog.Error(http.ListenAndServe(*addr, nil))
}
