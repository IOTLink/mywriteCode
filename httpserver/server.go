package main

import (
	"net"
	"net/http"
	"log"
	"fmt"
)
// Server is the fabric-ca server
type Server struct {
	// The server mux
	mux *http.ServeMux
	// The current listener for this server
	listener net.Listener
}

// AuthType is the enum for authentication types: basic and token
type authType int

const (
	noAuth authType = iota
	basic           // basic = 1
	token           // token = 2
)

const (
	enrollmentIDHdrName = "__eid__"
	caHdrName           = "__caname__"
)


// Register all endpoint handlers
func (s *Server) registerHandlers() {
	s.mux = http.NewServeMux()
	s.registerHandler("register", newRegisterHandler, token)
	s.registerHandler("enroll", newEnrollHandler, basic)
}

// Register an endpoint handler
func (s *Server) registerHandler(
	path string,
	getHandler func(server *Server) (http.Handler, error),
	at authType) {

	var handler http.Handler

	handler, err := getHandler(s)
	if err != nil {
		log.Printf("Endpoint '%s' is disabled: %s", path, err)
		return
	}

	handler = &fcaAuthHandler{
		server:   s,
		authType: at,
		next:     handler,
	}
	s.mux.Handle("/"+path, handler)
	s.mux.Handle("/api/v1/"+path, handler)
}


func (s *Server) listenAndServe(addr string) (err error) {
	addrStr := fmt.Sprintf("http://%s", addr)

	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("TCP listen failed for %s: %s", addrStr, err)
	}
	serveError := http.Serve(s.listener, s.mux)
	if serveError != nil {
		fmt.Sprint(serveError.Error())
		return serveError
	}
	return nil
}

func main() {
	var s Server
	// Register http handlers
	s.registerHandlers()

	log.Printf("%s","instance(s) running on server")

	// Start listening and serving
	s.listenAndServe("127.0.0.1:7055")
}

/*

 tcp 实现的htto服务
 */