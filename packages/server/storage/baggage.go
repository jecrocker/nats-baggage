package storage

// Baggage defines the interface to enable messages to be pulled from the store
type Baggage interface {
	// Store adds a message from a topic to that topics store
	Store(topic string, content string)
	// Get retrieves all messages from that timestamp
	Get(topic string, timestamp string)
	// Tidy removes stale components
	Tidy()
}
