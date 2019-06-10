package server

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/syifan/goseth"
)

// Server runs an HTTP server and shows the serialization results over webpage
type Server struct {
	item       interface{}
	fileServer http.Handler
}

// NewServer creates a new server. The server always returns the json
// representation of the provided item.
func NewServer(item interface{}) *Server {
	s := &Server{
		item:       item,
		fileServer: http.FileServer(assets),
	}
	return s
}

// Run starts the server on a random port.
func (s *Server) Run() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	fmt.Println("Seth server running on:", listener.Addr().(*net.TCPAddr).Port)

	panic(http.Serve(listener, s))
}

// ServeHTTP respond to client http request
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/serialize" {
		goseth.MakeSerializer().Serialize(s.item, w)
		return
	}

	if strings.HasPrefix(req.URL.Path, "/") {
		s.fileServer.ServeHTTP(w, req)
		return
	}

	w.WriteHeader(404)
}
