package config

// Configuration ...
type Configuration struct {
	ScanList   []ScanObject   `json:"scanList"`
	OutputList []OutputObject `json:"outputList"`
}

// ScanObject ...
type ScanObject struct {
	IP   string   `json:"ip"`
	Port []string `json:"port"`
}

// OutputObject ...
type OutputObject struct {
	Name string `json:"name"`
	Ref  string `json:"ref"`
}
