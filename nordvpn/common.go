package nordvpn

import (
	"github.com/sam7r/nordnm/utils"
	"io/ioutil"
	"net/http"
)

const nordBaseURL = "https://api.nordvpn.com/v1"

// makeRequest execute GET request with the given URI
func makeRequest(resourceURI string) (body []byte, err error) {
	req, err := http.NewRequest("GET", nordBaseURL+resourceURI, nil)
	return execReq(req)
}

// makeRequestWithCallback execute GET request with given URI, includes callback to modify the request before execution
func makeRequestWithCallback(resourceURI string, callback func(*http.Request)) (body []byte, err error) {
	req, err := http.NewRequest("GET", nordBaseURL+resourceURI, nil)
	callback(req)
	return execReq(req)
}

// execReq do the request
func execReq(req *http.Request) (body []byte, err error) {
	utils.Logger.Info(req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	utils.Logger.Info(resp)

	return ioutil.ReadAll(resp.Body)
}
