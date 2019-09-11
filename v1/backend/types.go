package backend

// RedisOption redis options
type RedisOption struct {
	// Addr host:port address.
	Addr string
}

// QueueOption of storage
type QueueOption struct {
	// DefaultQueue name. it will be "default" when you using default config
	Name string `json:"name"`
}
