package accounts

import (
	"log"
	"time"
)

// TimeTrack is to check how long the func running
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
