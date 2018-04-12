package config

import (
    "github.com/spf13/viper"
    "log"
)

func Load() {
    viper.SetConfigName("config")
    viper.SetConfigType("json")
    viper.AddConfigPath("./resources")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatal(err)
    }
    reloadEveryConfig()
}

func reloadEveryConfig() {
    Metric = loadMetric()
    Smtp = loadSmtp()
}
