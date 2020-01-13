package auto

import (
	"net/http"
	"strconv"
)

type Sling struct {
	// http.Client for request
	client		*http.Client
	// GET, POST, etc.
	HttpMethod	string
	Path		string
	Port 		int
	SSL 		bool
	Header 		http.Header

}

func (r *Sling) Request(Host string) (*http.Request, error) {
	var url string

	switch r.SSL {
	case true:
		url = "https://" + Host + ":" + strconv.Itoa(r.Port) + r.Path
	case false:
		fallthrough
	default:
		url = "http://" + Host + ":" + strconv.Itoa(r.Port) + r.Path
	}
	request, err := http.NewRequest(r.HttpMethod, url, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (r *Sling) Send(Request *http.Request) (*http.Response, error) {
	return nil, nil
}

func (r *Sling) Check(Response *http.Response, Regex string) (bool, error) {
	return false, nil
}
