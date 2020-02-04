package nordvpn

import "encoding/json"

// Country nordvpn technology type def
type Country struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// GetCountries gets a list of supported countries used by NordVPN
func GetCountries() (countries []Country, err error) {
	resourceURI := "/servers/countries"
	body, err := makeRequest(resourceURI)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &countries)
	return countries, err
}
