package configuration

// VERSION contain version of application
const VERSION = "development"

// Configuration represents configuration of application
type Configuration struct {
	Host string
	Port int

	EnableHTTPS        bool
	EnableRateLimit    bool
	EnableImageStorage bool
	MaxConnections     int
}
