package server

import (
	"net/http"
	"path"
	"runtime"
)

// go:generate vfsgendev -source="github.com/syifan/goseth/server".Assets

var assets http.FileSystem

func init() {
	_, filename, _, _ := runtime.Caller(0)
	assets = http.Dir(path.Dir(filename) + "/web")
}
