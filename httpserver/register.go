package main


import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	cfsslapi "github.com/cloudflare/cfssl/api"
	"github.com/cloudflare/cfssl/log"

	"github.com/hyperledger/fabric-ca/api"

)

// registerHandler for register requests
type registerHandler struct {
	server *Server
}

// newRegisterHandler is constructor for register handler
func newRegisterHandler(server *Server) (h http.Handler, err error) {
	// NewHandler is constructor for register handler
	return &cfsslapi.HTTPHandler{
		Handler: &registerHandler{server: server},
		Methods: []string{"POST"},
	}, nil
}

// Handle a register request
func (h *registerHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debug("Register request received")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body.Close()

	// Parse request body
	var req api.RegistrationRequestNet
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		return err
	}

	caname := r.Header.Get(caHdrName)

	// Register User
	callerID := r.Header.Get(enrollmentIDHdrName)


	resp := &api.RegistrationResponseNet{RegistrationResponse: api.RegistrationResponse{Secret: caname+callerID}}

	log.Debugf("Registration completed - sending response %+v", &resp)
	return cfsslapi.SendResponse(w, resp)
}

