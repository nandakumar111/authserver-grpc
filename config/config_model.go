package config

type ServerConfig struct {
	Database       Database    `json:"database"`
	Server         ServerNodes `json:"server"`
	OptionalServer ServerNodes `json:"optional_server"`
	TLS            TLSNodes    `json:"tls"`
}

type ServerNodes struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Status string `json:"status"`
}

type TLSNodes struct {
	Active      bool   `json:"active"`
	Certificate string `json:"certificate"`
	KeyFile     string `json:"key_file"`
}

type Database struct {
	MongoDB mongo `json:"mongo_db"`
}

type mongo struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
