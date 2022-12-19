// Package weather implements weather forecasting.
package weather

// CurrentCondition tells about current weather.
var CurrentCondition string

// CurrentLocation tells about the current location.
var CurrentLocation string

// Forecast returns the current location and it's condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
