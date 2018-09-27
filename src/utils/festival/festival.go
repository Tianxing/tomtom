package festival

import (
    "time"
    "strconv"
)

func Stage() int64 {
    t := time.Now()
    hour12,_ := time.ParseDuration("12h")
    stage, _ := strconv.ParseInt(t.Add(hour12).Format("20060102"), 10, 64)
    return stage
}
