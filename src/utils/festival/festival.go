package festival

import (
    "time"
)

func Stage() string{
    t := time.Now()
    hour12,_ := time.ParseDuration("12h")
    return t.Add(hour12).Format("20060102")
}
