package main

import (
	"io/ioutil"
	"net/http"

	cfapi "github.com/cloudflare/cfssl/api"
	"github.com/cloudflare/cfssl/log"
	"github.com/hyperledger/fabric-ca/api"
	"github.com/hyperledger/fabric-ca/util"
)


func newEnrollHandler(server *Server) (h http.Handler, err error) {
	return newSignHandler(server, "enroll")
}


// signHandler for enroll or reenroll requests
type signHandler struct {
	server *Server
	// "enroll" or "reenroll"
	endpoint string
}



// newSignHandler is the constructor for an enroll or reenroll handler
func newSignHandler(server *Server, endpoint string) (h http.Handler, err error) {
	// NewHandler is constructor for register handler
	return &cfapi.HTTPHandler{
		Handler: &signHandler{server: server, endpoint: endpoint},
		Methods: []string{"POST"},
	}, nil
}



// Handle an enroll or reenroll request.
// Authentication has already occurred for both enroll and reenroll prior
// to calling this function in auth.go.
func (sh *signHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Received request for endpoint %s", sh.endpoint)
	err := sh.handle(w, r)
	if err != nil {
		log.Errorf("Enrollment failure: %s", err)
	}
	return err
}

func (sh *signHandler) handle(w http.ResponseWriter, r *http.Request) error {

	// Read the request's body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body.Close()

	var req api.EnrollmentRequestNet

	err = util.Unmarshal(body, &req, sh.endpoint)
	if err != nil {
		return err
	}

	log.Debugf("Enrollment request: %+v\n", req)

	caname := r.Header.Get(caHdrName)
	enrollmentID := r.Header.Get(enrollmentIDHdrName)

	byte := []byte(caname + enrollmentID)
	resp := &enrollmentResponseNet{Cert: util.B64Encode(byte)}
	return cfapi.SendResponse(w, resp)
}

