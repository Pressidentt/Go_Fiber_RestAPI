package config

import (
	"os"
	"strconv"
)

type Ports struct {
	PORT int
}

type DbCredentials struct {
	dbname     string
	dbpassword string
	dbhost     string
	dbport     string
}

type Config struct {
	Ports         Ports
	dbCredentials DbCredentials
}

func New() *Config {
	t := new(Ports)
	tmp := os.Getenv("PORT")
	tmpPort, err := strconv.Atoi(tmp)
	if err != nil || tmp == "" {
		tmpPort = 3000
	}
	t.PORT = tmpPort

	return &Config{
		Ports: *t,
	}
}
