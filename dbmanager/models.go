package dbmanager

type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type DbModel struct {
}
