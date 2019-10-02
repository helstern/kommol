package headers

import (
	"net/http"
	"strconv"
	"time"
)

func FormatTime(value time.Time) string {
	return value.UTC().Format(http.TimeFormat)
}

func FormatInt(value int64) string {
	return strconv.FormatInt(value, 10)
}
