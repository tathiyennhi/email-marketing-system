package date

import (
	"time"
)

func FormatDate() string {
	return time.Now().Format("02 Jan 2006")
}
