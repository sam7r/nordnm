package nordvpn

import (
	"encoding/json"
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
	body, err := makeRequest(resourceURI)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &groups)
	return groups, err
}
