package output

// STDOutOutputer ...
type STDOutOutputer struct{}

// Output ...
func (o STDOutOutputer) Output(msg string) error {
	return nil
}
