package accessibleLG

import "gopkg.in/ini.v1"

var Conf = new(Config)

type Config struct {
	AppName string `ini:"app_name"`
	Mode    string `ini:"mode"`

	MySQL MySQLConfig `ini:"mysql"`
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
