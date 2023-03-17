package helpers

import (
	"fmt"
	"time"
)

func ConvertDate(date string) string {
	myDate, _ := time.Parse("2006-01-02", date)
	return fmt.Sprintf("%s", myDate.Format("02/01/2006"))
}

func YYYmmDD(date string) string {
	myDate, _ := time.Parse("02/04/2006", date)
	return myDate.Format("2006-01-02")
}
