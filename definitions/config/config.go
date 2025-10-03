package config

type Config struct {
	Host     string         `json:"host"`
	Port     int            `json:"port"`
	User     string         `json:"user"`
	Password string         `json:"password"`
	Database string         `json:"database"`
	Params   map[string]any `json:"params"`
}
