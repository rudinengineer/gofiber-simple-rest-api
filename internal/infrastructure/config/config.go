package config

type Config struct {
	Host     Host
	Database Database
}

type Host struct {
	Host string
	Port string
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
