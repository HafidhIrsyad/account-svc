package accounts

import (
	"log"
	"strconv"
	"time"
)

// TimeTrack is to check how long the func running
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func ParseStrToInt64(str string) int64 {
	num, _ := strconv.ParseInt(str, 10, 64)
	return num
}
