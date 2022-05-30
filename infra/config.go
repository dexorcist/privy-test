package infra

type MergeConfig struct {
	AppConfig
	AppCustomConfig
}

type AppConfig struct {
	Host                    string `mapstructure:"host"`
	Port                    string `mapstructure:"port"`
	ReadTimeout             int    `mapstructure:"read_timeout"`
	WriteTimeout            int    `mapstructure:"write_timeout"`
	IdleTimeout             int    `mapstructure:"idle_timeout"`
	LoggerOutput            string `mapstructure:"logger_output"`
	LoggerLevel             string `mapstructure:"logger_level"`
	LoggerWithTimestamp     bool   `mapstructure:"logger_with_timestamp"`
	MysqlDatabaseConnection string `mapstructure:"mysql_database_connection"`
	CustomConfig            string `mapstructure:"custom_config"`
	Namespace               string
	GoVersion               string
	BuildVersion            string
	Environment             string `mapstructure:"environment"`
}

type AppCustomConfig struct {
	// Swagger
	Swagger SwaggerOptions
	//Database
	Database DatabaseOption
}

type SwaggerOptions struct {
	Enabled   bool
	BasicAuth BasicAuthOptions
	// Path defines endpoint for swagger
	Path string
	// DocFile defines file path location of swagger doc
	DocFile string
	// Custom swagger
	SwaggerTemplate SwaggerTemplateOptions
}

type BasicAuthOptions struct {
	Username, Password string
}

type SwaggerTemplateOptions struct {
	BasicAuth    BasicAuthOptions
	Enabled      bool
	TemplateFile string
	Path         string
	GoTemplate   GoTemplateOptions
}

type GoTemplateOptions struct {
	Description string
	Title       string
	Version     string
	Schemes     string
	Host        string
	BasePath    string
}

type DatabaseOption struct {
	Host     string
	Port     int32
	User     string
	Password string
	Name     string
	SslMode  string
}
