package output

import "log"

// STDOutOutputer ...
type STDOutOutputer struct{}

// Output ...
func (o STDOutOutputer) Output(msg string) error {
	log.Printf("扫描失败\t\t\t%s", msg)
	return nil
}
