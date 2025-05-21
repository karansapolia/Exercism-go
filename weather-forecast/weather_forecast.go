// Package weather provides the Forecast function to return a weather forecast.
package weather

// CurrentCondition represents the current weather condition in the provided city.
var CurrentCondition string

// CurrentLocation represents the current city.
var CurrentLocation string

// Forecast take city name and condition strings as variables and returns a string with city name and current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
