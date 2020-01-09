package output

// Outputer ...
type Outputer interface {
	Output(msg string) error
}
