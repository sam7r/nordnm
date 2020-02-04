package nordvpn

import (
	"encoding/json"
)

// Technology nordvpn technology type def
type Technology struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

// GetTechnologies gets a list of supported technologies used by NordVPN
func GetTechnologies() (technologies []Technology, err error) {
	resourceURI := "/technologies"
	body, err := makeRequest(resourceURI)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &technologies)
	return technologies, err
}
