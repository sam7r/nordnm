package nordvpn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"nordnm/logger"
)

// Group nordvpn technology type def
type Group struct {
	ID         uint32 `json:"id"`
	Title      string `json:"title"`
	Identifier string `json:"identifier"`
}

// GetGroups gets a list of supported groups used by NordVPN
func GetGroups() (groups []Group, err error) {

	resourceURI := "/servers/groups"
	req, err := http.NewRequest("GET", nordBaseURL+resourceURI, nil)

	logger.Stdout.Info(req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	logger.Stdout.Info(resp)
	err = json.Unmarshal(body, &groups)

	return groups, err
}
