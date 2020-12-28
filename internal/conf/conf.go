package conf

import (
	"github.com/go-ini/ini"
)

type database struct {
	Driver   string `json:"driver";ini:"driver"`
	Host     string `json:"host";ini:"host"`
	Port     string `json:"port";ini:"port"`
	User     string `json:"user";ini:"user"`
	Password string `json:"pass";ini:"pass"`
	Path     string `json:"path";ini:"path"`
}

type mail struct {
	Host     string `json:"host";ini:"host"`
	Port     string `json:"port";ini:"port"`
	Login    string `json:"login";ini:"login"`
	Password string `json:"password";ini:"password"`
}

type admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Config struct {
	File     *ini.File
	Prod     bool     `json:"prod";ini:"prod"`
	Database database `json:"database";ini:"database"`
	Mail     mail     `json:"mail";ini:"mail"`
	Admin    admin    `json:admin`
}

var (
	Cfg      *Config
	confPath string = "conf/app.ini"
)

func Read() error {

	var err error

	Cfg = new(Config)
	if Cfg.File, err = ini.InsensitiveLoad(confPath); err != nil {
		return err
	}

	Cfg.Prod = Cfg.File.Section("").Key("prod").MustBool(false)

	Cfg.Database.Driver = Cfg.File.Section("database").Key("driver").In(
		"sqlite3", []string{"sqlite3", "postgres", "mysql"},
	)
	Cfg.Database.Host = Cfg.File.Section("database").Key("host").String()
	Cfg.Database.Port = Cfg.File.Section("database").Key("port").String()
	Cfg.Database.User = Cfg.File.Section("database").Key("user").String()
	Cfg.Database.Password = Cfg.File.Section("database").Key("password").String()
	Cfg.Database.Path = Cfg.File.Section("database").Key("path").String()

	Cfg.Mail.Host = Cfg.File.Section("mail").Key("host").String()
	Cfg.Mail.Port = Cfg.File.Section("mail").Key("port").String()

	return nil

}

func Save() {

	// TODO Переделать на цикл
	Cfg.File.Section("database").Key("driver").SetValue(Cfg.Database.Driver)
	Cfg.File.Section("database").Key("host").SetValue(Cfg.Database.Host)
	Cfg.File.Section("database").Key("port").SetValue(Cfg.Database.Port)
	Cfg.File.Section("database").Key("user").SetValue(Cfg.Database.User)
	Cfg.File.Section("database").Key("password").SetValue(Cfg.Database.Password)
	Cfg.File.Section("database").Key("path").SetValue(Cfg.Database.Path)

	Cfg.File.Section("mail").Key("host").SetValue(Cfg.Mail.Host)
	Cfg.File.Section("mail").Key("port").SetValue(Cfg.Mail.Port)

	Cfg.File.SaveTo(confPath)
}

func Prod() bool {
	return Cfg.Prod
}
