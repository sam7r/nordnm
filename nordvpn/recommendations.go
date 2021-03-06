package nordvpn

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	body, err := makeRequestWithCallback(resourceURI, func(req *http.Request) {
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
	})

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &recommendations)
	return recommendations, err
}
