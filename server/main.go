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

func PilosaHandler(w http.ResponseWriter, r *http.Request) {
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
	go pilosa(query, file)
	rest.MustEncode(w, rest.RestMessage{"ok", "job submitted"})
}

func SetdiffHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("addr=%s  method=%s host=%s uri=%s",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	r.ParseForm()
	f1 := strings.TrimSpace(r.FormValue("f1"))
	f2 := strings.TrimSpace(r.FormValue("f2"))
	f3 := strings.TrimSpace(r.FormValue("f3"))
	if f1 == "" || f2 == "" || f3 == "" {
		rest.MustEncode(w, rest.RestMessage{"error", "f1 or f2 or f3 is empty"})
		return
	}
	f1 = filepath.Join(*dir, f1)
	f2 = filepath.Join(*dir, f2)
	f3 = filepath.Join(*dir, f3)
	err := setdiff(f1, f2, f3)
	if err != nil {
		rest.MustEncode(w, rest.RestMessage{"error", err.Error()})
		return
	}
	rest.MustEncode(w, rest.RestMessage{"ok", "done"})
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.Handle("/api/pilosa", rest.WithLog(PilosaHandler))
	http.Handle("/api/setdiff", rest.WithLog(SetdiffHandler))
	http.Handle("/", http.FileServer(rice.MustFindBox("ui").HTTPBox()))
	glog.Info("server listen on", *addr)
	glog.Error(http.ListenAndServe(*addr, nil))
}
