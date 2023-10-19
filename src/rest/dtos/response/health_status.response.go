package response

type HealthStatusResponse struct {
	ServerStatus   string `json:"server_status"`
	DatabaseStatus string `json:"database_status"`
	DatabaseStats  Stats
}

type Stats struct {
	MaxOpenConnections int `json:"max_open_connections"`
	OpenConnections    int `json:"open_connections"`
	InUse              int `json:"in_use"`
	Idle               int `json:"idle"`
}
