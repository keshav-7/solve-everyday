package TimeConversion

import (
	"fmt"
	"time"
)

func main() {
	newtime := convertTimeFormat("02:30:45PM")
	fmt.Println(newtime)
}

func convertTimeFormat(twelveTime string) string {
	inFormat := "03:04:05PM"
	outFormat := "15:04:05"
	t, error := time.Parse(inFormat, twelveTime)
	if error != nil {
		fmt.Println("getting error while parsing input format: ", error)
		return ""
	}
	time24 := t.Format(outFormat)
	return time24
}
