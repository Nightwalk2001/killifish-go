package docs

type State struct {
	Born    int `json:"born"`
	Killed  int `json:"killed"`
	Current int `json:"current"`
	Total   int `json:"total"`
}

type Operation struct {
	Executor string  `json:"executor"`
	Type     string  `json:"type"`
	Quantity *int    `json:"quantity,omitempty"`
	Prev     *int    `json:"prev,omitempty"`
	Current  *int    `json:"current,omitempty"`
	Time     string  `json:"time"`
	Note     *string `json:"note,omitempty"`
}
