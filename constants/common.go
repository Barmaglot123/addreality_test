package constants

const (
    GetMetrics = "SELECT * FROM \"device_metrics\"  WHERE (server_time > $1) ORDER BY server_time"
    GetEmailFromDevice = "SELECT email FROM users INNER JOIN devices ON users.id = devices.user_id WHERE devices.id = $1"
    SetDeviceAlert = "INSERT INTO device_alerts(device_id, message) VALUES($1, $2)"

    OutOfRangeAlert = "Value in the %s field is outside the range"
    AlertMailBody = "Values ​​came out of limits in these fields: "
    RedisAlertKey = "_device_alert"
)
