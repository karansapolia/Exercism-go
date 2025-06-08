package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	gigaValue := math.Pow(10, 9)
	oneGigasecond := time.Duration(gigaValue) * time.Second
	oneGigaSecondLater := t.Add(oneGigasecond)
	return oneGigaSecondLater
}
