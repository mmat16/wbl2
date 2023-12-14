package ntp

import (
	"time"

	"github.com/beevik/ntp"
)

// GetPreciseTime returns the current time from the NTP server
func GetPreciseTime(ntpPool string) (time.Time, error) {
	// Get the (local) time from the NTP server
	ntpTime, err := ntp.Time(ntpPool)
	if err != nil {
		return time.Now(), err
	}
	return ntpTime, nil
}
