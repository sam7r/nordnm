package nordvpn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const nordBaseURL = "https://api.nordvpn.com/v1"

// RecommendationFilters filters to apply to recommendation query
type RecommendationFilters struct {
	ServerGroupID string
	CountryID     uint8
	TechnologyID  string
	Limit         uint8
}

// Recommendation body response from recommendation query
type Recommendation struct {
	ID       uint16 `json:"id"`
	Name     string `json:"name"`
	Station  string `json:"station"`
	Hostname string `json:"hostname"`
	Load     string `json:"load"`
	Status   string `json:"status"`
}

// GetRecommendations request to NordVPN web api for recommended servers
func GetRecommendations(filters RecommendationFilters) (recommendations []Recommendation, err error) {

	resp, err := http.Get(nordBaseURL + "/servers/recommendations")

	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		print(err)
	}

	err = json.Unmarshal(body, &recommendations)

	return recommendations, err
}

// GetServerGroups gets a list of groups which are available within NordVPN
func GetServerGroups() {}

// GetServerCountries gets a list of countries where servers are available
func GetServerCountries() {}

// GetTechnologies gets a list of supported technologies used by NordVPN
func GetTechnologies() {}

// GetServerConfigFile downloads the OVPN file for the given server
func GetServerConfigFile() {}
