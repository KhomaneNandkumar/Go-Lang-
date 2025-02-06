package config

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"db_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}
