package configs

var Config Configuration

type ServerConfiguration struct {
	Host     string
	GRPCPort string
	HTTPPort string
}

type DatabaseConfiguration struct {
	Host            string
	Port            int
	User            string
	Password        string
	DB              string
	ConnMaxLifetime int
	MaxIdleConns    int
	MaxOpenConns    int
}

type Credential struct {
	TokenSecret string
}

type Configuration struct {
	Server     ServerConfiguration
	Database   DatabaseConfiguration
	Credential Credential
}
