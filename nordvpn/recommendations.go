package nordvpn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nordnm/logger"
)

// RecommendationFilters filters to apply to recommendation query
type RecommendationFilters struct {
	ServerGroupID string
	CountryID     uint8
	TechnologyID  string
	Limit         uint8
}

// Recommendation body response from recommendation query
type Recommendation struct {
	ID           uint32       `json:"id"`
	Name         string       `json:"name"`
	Station      string       `json:"station"`
	Hostname     string       `json:"hostname"`
	Load         uint16       `json:"load"`
	Status       string       `json:"status"`
	Technologies []Technology `json:"technologies"`
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
	// set a default limit if none set
	if filters.Limit == 0 {
		filters.Limit = 10
	}
	q.Add("limit", fmt.Sprint(filters.Limit))

	// encode query into request
	req.URL.RawQuery = q.Encode()
	logger.Stdout.Info(req.URL.String())

	// make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// read out response and unmarshal into type
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &recommendations)
	logger.Stdout.Info(resp)

	return recommendations, err
}

// GetServerGroups gets a list of groups which are available within NordVPN
func GetServerGroups() {}

// GetServerCountries gets a list of countries where servers are available
func GetServerCountries() {}

// GetServerConfigFile downloads the OVPN file for the given server
func GetServerConfigFile() {}
