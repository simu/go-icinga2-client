package icinga2

import (
	"fmt"
)

type PerfData []string

type Command []string

// TODO: is this correct?
type TimeStamp string

type Action struct {
	ExitStatus      int       `json:"exit_status"`
	PluginOutput    string    `json:"plugin_output"`
	PerformanceData PerfData  `json:"performance_data,omitempty"`
	CheckCommand    Command   `json:"check_command,omitempty"`
	CheckSource     string    `json:"check_source,omitempty"`
	ExecutionStart  TimeStamp `json:"execution_start,omitempty"`
	ExecutionEnd    TimeStamp `json:"execution_end,omitempty"`
	TTL             int       `json:"ttl"`
	Filter          string    `json:"filter"`
	Type            string    `json:"type"`
}

func (s *WebClient) ProcessCheckResult(service Service, action Action) error {
	var results, errmsg Results

	path := s.URL + "/v1/actions/process-check-result"

	// Update action struct with filters
	action.Filter = fmt.Sprintf("host.name==\"%s\"&&service.name==\"%s\"",
		service.HostName, service.Name)
	action.Type = "Service"

	resp, err := s.napping.Post(path, action, &results, &errmsg)
	if err != nil {
		fmt.Printf("[PCR] Post failed: %v\n", err)
		return err
	}

	return s.handleResults("process-check-result", path, resp, &results, &errmsg, err)
}

func (s *MockClient) ProcessCheckResult(service Service, action Action) error {
	s.mutex.Lock()
	s.Actions[service.FullName()] = append(s.Actions[service.FullName()], action)
	s.mutex.Unlock()
	return nil
}
