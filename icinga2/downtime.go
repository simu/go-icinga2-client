package icinga2

type Downtime struct {
	Active    bool    `json:"active"`
	Author    string  `json:"author"`
	Comment   string  `json:"comment"`
	EndTime   float64 `json:"end_time"`
	Fixed     bool    `json:"fixed"`
	Host      string  `json:"host_name"`
	Name      string  `json:"name,omitempty"`
	Service   string  `json:"service_name"`
	StartTime float64 `json:"start_time"`
	Type      string  `json:"type"`
	Zone      string  `json:"zone"`
}

type DowntimeResults struct {
	Results []struct {
		Downtime Downtime `json:"attrs"`
	} `json:"results"`
}

func (s *WebClient) ListDowntimes(query string) (downtimes []Downtime, err error) {
	var dtResults DowntimeResults
	downtimes = []Downtime{}

	_, err = s.napping.Get(s.URL+"/v1/objects/downtimes?"+query, nil, &dtResults, nil)
	if err != nil {
		return
	}
	for _, result := range dtResults.Results {
		if result.Downtime.Type == "Downtime" {
			if s.Zone == "" || s.Zone == result.Downtime.Zone {
				downtimes = append(downtimes, result.Downtime)
			}
		}
	}

	return
}

func (s *MockClient) ListDowntimes(query string) ([]Downtime, error) {
	// Readonly mocked ListDowntimes never returns a downtime
	return []Downtime{}, nil
}
