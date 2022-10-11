package helpers

import (
	"fmt"
	"time"
)

func ConvertDate(date string) string {
	myDate, _ := time.Parse("2006-01-02", date)
	return fmt.Sprintf("%s", myDate.Format("January 02, 2006"))
}
