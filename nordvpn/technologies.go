package nordvpn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"nordnm/logger"
)

// Technology nordvpn technology type def
type Technology struct {
	ID         uint32 `json:"id"`
	Name       string `json:"Name"`
	Identifier string `json:"Identifier"`
}

// GetTechnologies gets a list of supported technologies used by NordVPN
func GetTechnologies() (technologies []Technology, err error) {

	resourceURI := "/technologies"
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
	err = json.Unmarshal(body, &technologies)

	return technologies, err
}
