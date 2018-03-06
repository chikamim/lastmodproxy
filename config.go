package main

type Config struct {
	Websites []WebSite `yaml:"websites"`
}

type WebSite struct {
	URLFilter  string `yaml:"url_filter"`
	DateMatch  string `yaml:"date_match"`
	DateLayout string `yaml:"date_layout"`
	TimeZone   string `yaml:"timezone"`
}

func NewConfig() *Config {
	return &Config{}
}
