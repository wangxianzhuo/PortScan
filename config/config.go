package config

// Configuration ...
type Configuration struct {
	ScanList   []ScanObject   `json:"scanList"`
	OutputList []OutputObject `json:"outputList"`
}

// ScanObject ...
type ScanObject struct {
	IP    string `json:"ip"`
	Ports []int  `json:"ports"`
}

// OutputObject ...
type OutputObject struct {
	Name string `json:"name"`
	Ref  string `json:"ref"`
}
