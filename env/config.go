package env

type Configuration struct {
	AppEnv      string `env:"APP_ENV"`
	AppName     string `env:"APP_NAME"`
	LogLevel    string `env:"LOG_LEVEL"`
	RunningMode string `env:"RUNNING_MODE"`
}
