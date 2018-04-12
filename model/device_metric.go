package model

import (
    "time"
    "database/sql"
)

type DeviceMetric struct {
    ID        int
    DeviceID  int
    Metric1   sql.NullInt64
    Metric2   sql.NullInt64
    Metric3   sql.NullInt64
    Metric4   sql.NullInt64
    Metric5   sql.NullInt64
    LocalTime time.Time
    CreatedAt time.Time
}