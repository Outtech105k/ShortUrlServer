package utils

type Config struct {
	ServiceServerEndpoint  string `env:"MAIN_ENDPOINT"`
	RedirectServerEndpoint string `env:"REDIRECT_ENDPOINT"`
}
