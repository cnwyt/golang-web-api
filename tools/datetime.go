package tools

import "time"

func GetDate(params ...int64) string {
    res := ""
    format := "2006-01-02"
    if len(params) > 0 {
        timestamp := string(params[0])
        res = time.Unix(timestamp, 0).Format(format)
    } else {
        res = time.Now().Format(format)
    }

    return res;
}

func GetDatetime(params ...int64) string {
    res := ""
    format := "2006-01-02 15:04:05"

    if len(params) > 0 {
        timestamp := int64(params[0])
        res = time.Unix(timestamp, 0).Format(format)
    } else {
        res = time.Now().Format(format)
    }

    return res;
}
