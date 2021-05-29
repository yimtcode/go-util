package event

// Event is interface
type Event interface {
	SetEvent(key interface{}, callback func()) bool
}