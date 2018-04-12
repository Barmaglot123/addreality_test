package config

import "github.com/spf13/viper"

type MetricConfig interface {
    MinLim() int64
    MaxLim() int64
}

var Metric metric

type metric struct {
    minLim int64
    maxLim int64
}

func loadMetric() metric {
    s := viper.Sub("metric")
    return metric{
        minLim: s.GetInt64("min_lim"),
        maxLim: s.GetInt64("max_lim"),
    }
}

func(m metric) MinLim() int64 {
    return m.minLim
}

func(m metric) MaxLim() int64{
    return m.maxLim
}