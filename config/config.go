package config

// Server Config
type Server struct {
	Port int `mapstructure:"port"`
}

// Database Config
type Database struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	Driver   string `mapstructure:"driver"`
}

type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
	// File     `mapstructure:"file"`
	// Token    `mapstructure:"token"`
	// TokenConfig
	DefaultRowsPerPage string `mapstructure:"DEFAULT_ROWS_PER_PAGE"`
}
