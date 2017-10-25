package main


import (

	"fmt"
	"io/ioutil"
	"net/http"


	"github.com/cloudflare/cfssl/api"
	"github.com/cloudflare/cfssl/log"
	"github.com/hyperledger/fabric-ca/util"
)



// Fabric CA authentication handler
type fcaAuthHandler struct {
	server   *Server
	authType authType
	next     http.Handler
}



func (ah *fcaAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := ah.serveHTTP(w, r)
	if err != nil {
		api.HandleError(w, err)
	} else {
		ah.next.ServeHTTP(w, r)
	}
}

// Handle performs authentication
func (ah *fcaAuthHandler) serveHTTP(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Received request\n%s", util.HTTPRequestToString(r))

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Debugf("Failed to read body: %s", err)
		return err
	}
	fmt.Sprint(body)

	return nil

}


