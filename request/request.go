package request

import (
	"fmt"
	"net/http"
)

var accessToken string

type Request struct {
	endpoint string
	client   *http.Client
}

// Set sets the client attributes for the HTTP request.
func (r *Request) Set(e, access_token string) {
	accessToken = access_token
	r.endpoint = e
	r.client = &http.Client{
		Transport: &myTransport{},
	}
}

// Post sends an HTTP POST request to the specified route.
func (r *Request) Post(requestURL string) {
}

// Get sends an HTTP GET request to the specified route.
func (r *Request) Get(rourequestURLte string) {
}

// Put sends an HTTP PUT request to the specified route.
func (r *Request) Put(requestURL string) {
}

// Delete sends an HTTP DELETE request to the specified route.
func (r *Request) Delete(requestURL string) {
}

// Custom Transport with default headers
type myTransport struct {
	Transport http.RoundTripper
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Set the default headers here
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// If you want to add more default headers, you can set them here

	// Use the default transport if provided; otherwise, use http.DefaultTransport
	if t.Transport != nil {
		return t.Transport.RoundTrip(req)
	}
	return http.DefaultTransport.RoundTrip(req)
}
