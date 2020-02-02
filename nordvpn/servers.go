package nordvpn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nordnm/logger"
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

	resourceURI := "/servers/recommendations"
	req, err := http.NewRequest("GET", nordBaseURL+resourceURI, nil)

	// start appending query parmas if present in filters
	q := req.URL.Query()

	if filters.ServerGroupID != "" {
		q.Add("filters[servers_groups][identifier]", filters.ServerGroupID)
	}
	if filters.TechnologyID != "" {
		q.Add("filters[servers_technologies][identifier]", filters.TechnologyID)
	}
	if filters.CountryID != 0 {
		q.Add("filters[country_id]", fmt.Sprint(filters.CountryID))
	}

	if filters.Limit == 0 {
		filters.Limit = 10
	}

	q.Add("limit", fmt.Sprint(filters.Limit))
	req.URL.RawQuery = q.Encode()

	logger.STDout.Info(req.URL.String())

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
