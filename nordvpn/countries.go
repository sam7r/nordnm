package nordvpn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"nordnm/logger"
)

// Country nordvpn technology type def
type Country struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// GetCountries gets a list of supported countries used by NordVPN
func GetCountries() (countries []Country, err error) {

	resourceURI := "/servers/countries"
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
	err = json.Unmarshal(body, &countries)

	return countries, err
}
