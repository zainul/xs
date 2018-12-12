package counter

// Counter is package use for counting some value , can added from cache or storage
type Counter interface {
	// Increament set counter value increase
	Increament(key string) (int64, error)
	// Decreament set counter value decrease
	Decreament(key string) (int64, error)
	// Get counter value
	Get(key string) (int64, error)
}
