package output

// Outputer ...
type Outputer interface {
	Output(msg, address string) error
}
