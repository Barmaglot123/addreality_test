package config

import "github.com/spf13/viper"

var Smtp smtp

type SmtpConfig interface {
    Username() string
    Password() string
    Host() string
    Port() int
    AuthType() string
}

type smtp struct {
    username string
    password string
    host     string
    port     int
    authType string
}

func loadSmtp() smtp{
    s := viper.Sub("smtp")
    return smtp{
        username: s.GetString("username"),
        password: s.GetString("password"),
        host:     s.GetString("host"),
        port:     s.GetInt("port"),
        authType: s.GetString("auth_type"),
    }
}

func (s smtp)Username() string {
    return s.username
}

func (s smtp)Password() string {
    return s.password
}

func (s smtp)Host() string {
    return s.host
}

func (s smtp)Port() int {
    return s.port
}

func (s smtp)AuthType() string {
    return s.authType
}