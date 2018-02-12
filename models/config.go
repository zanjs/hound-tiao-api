package models

// Config is
var Config = struct {
	APP struct {
		Name string
		Port string `default:"8080"`
	}
	Redis struct {
		Host     string
		Port     string `default:"6379"`
		Password string `default:"root"`
		Expire   string `default:"120"`
	}
}{}
