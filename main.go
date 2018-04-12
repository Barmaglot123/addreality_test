package main

import (
    . "addreality_test/datasource"
    . "addreality_test/constants"
    "addreality_test/config"
    "addreality_test/model"
    _ "github.com/lib/pq"
    "database/sql"
    "net/smtp"
    "strconv"
    "time"
    "fmt"
    "log"
)


func main() {
    config.Load()
    SetupSql()
    defer DB.Close()
    SetupRedis()
    defer Redis.Close()
    var t time.Time = time.Time{}

    for {
        getMetrics(&t)
    }
}

func getMetrics(t *time.Time) {
    rows, err := DB.Query(GetMetrics, t)
    defer rows.Close()
    m := model.DeviceMetric{}

    if err != nil {
        log.Fatal(err)
    }

    for rows.Next() {
        err := rows.Scan(&m.ID, &m.DeviceID, &m.Metric1, &m.Metric2, &m.Metric3,
            &m.Metric4, &m.Metric5, &m.LocalTime, &m.CreatedAt)

        if err != nil {
            log.Fatal(err)
        }
        checkMetricsValue(m)
    }

    if m.CreatedAt.After(*t) {
        *t = m.CreatedAt
    }
    err = rows.Err()

    if err != nil {
        log.Fatal(err)
    }
}

func checkMetricsValue(m model.DeviceMetric) {
    var sendMail bool
    var mailBody string = AlertMailBody

    for i, v := range []interface{}{m.Metric1, m.Metric2, m.Metric3, m.Metric4, m.Metric5} {
        a := v.(sql.NullInt64)
        if !a.Valid { continue }

        if a.Int64 > config.Metric.MaxLim() || a.Int64 < config.Metric.MinLim(){
            sendMail = true
            fn := "Metric" + strconv.Itoa(i + 1)
            mailBody = mailBody + fn + " "
            setAlertToRedis(m.DeviceID, fn)
            setAlertToDatabase(m.DeviceID, fn)
        }
    }

    if sendMail {
        go sendEmail(m.DeviceID, mailBody)
    }
}

func setAlertToRedis(deviceID int, fieldName string){
    strID := strconv.FormatUint(uint64(deviceID), 10)
    alert := fmt.Sprintf(OutOfRangeAlert, fieldName)
    Redis.Set(strID + RedisAlertKey, alert, 0)
}

func setAlertToDatabase(deviceID int, fieldName string){
    message := fmt.Sprintf(OutOfRangeAlert, fieldName)
    p, err := DB.Prepare(SetDeviceAlert)

    if err != nil {
        log.Fatal(err)
    }

    _, err = p.Exec(deviceID, message)

    if err != nil {
        log.Fatal(err)
    }
}

func sendEmail (deviceID int, mailBody string) error {
    email := ""
    err := DB.QueryRow(GetEmailFromDevice, deviceID).Scan(&email)

    if err != nil {
        log.Fatal(err)
    }

    auth := smtp.PlainAuth("", config.Smtp.Username(), config.Smtp.Password(), config.Smtp.Host())
    err = smtp.SendMail(
        config.Smtp.Host() + ":" + strconv.Itoa(config.Smtp.Port()),
        auth,
        config.Smtp.Username(),
        []string{email},
        []byte(mailBody),
    )
    return nil
}