package date_utils

import "time"

const (
	apiDbLayout = "2006-01-02 10:10:10"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(time.RFC3339)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
