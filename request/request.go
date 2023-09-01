package request

import (
	"net/http"
	"net/url"
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
func (r *Request) Post(endpoint string, body interface{}) {
	req, err := http.NewRequest(http.MethodPost, r.endpoint+endpoint, nil)
}

// Get sends an HTTP GET request to the specified route.
func (r *Request) Get(endpoint string, body interface{}) {
	req, err := http.NewRequest(http.MethodGet, r.endpoint+endpoint, nil)

}

// Put sends an HTTP PUT request to the specified route.
func (r *Request) Put(endpoint string, body interface{}) {
	req, err := http.NewRequest(http.MethodPut, r.endpoint+endpoint, nil)

}

// Delete sends an HTTP DELETE request to the specified route.
func (r *Request) Delete(endpoint string, body interface{}) {
	req, err := http.NewRequest(http.MethodDelete, r.endpoint+endpoint, nil)

}

// Custom Transport with default headers
type myTransport struct {
	Transport http.RoundTripper
}

func (t *myTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Set the default headers here
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	req.Header.Set("access_token", accessToken)

	// If you want to add more default headers, you can set them here

	// Use the default transport if provided; otherwise, use http.DefaultTransport
	if t.Transport != nil {
		return t.Transport.RoundTrip(req)
	}
	return http.DefaultTransport.RoundTrip(req)
}
