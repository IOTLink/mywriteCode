package main

type RegistrationRequest struct {
	// Name is the unique name of the identity
	Name string `json:"id" help:"Unique name of the identity"`
	// Type of identity being registered (e.g. "peer, app, user")
	Type string `json:"type" def:"user" help:"Type of identity being registered (e.g. 'peer, app, user')"`
}

type RegistrationResponse struct {
	// The secret returned from a successful registration response
	Secret string `json:"secret"`
}












// EnrollmentRequest is a request to enroll an identity
type EnrollmentRequest struct {
	// The identity name to enroll
	Name string `json:"name" skip:"true"`
	// The secret returned via Register
	Secret string `json:"secret,omitempty" skip:"true" mask:"password"`
}

// The enrollment response from the server
type enrollmentResponseNet struct {
	// Base64 encoded PEM-encoded ECert
	Cert string

}
