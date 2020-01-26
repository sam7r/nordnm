package nordvpn

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const nordBaseURL = "https://api.nordvpn.com/v1"

// Recommendations request to NordVPN web api for recommended servers
func Recommendations() {
	resp, err := http.Get(nordBaseURL + "/servers/recommendations")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))
}

func groups() {}

func countries() {}
