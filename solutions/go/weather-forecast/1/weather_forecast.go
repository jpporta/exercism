// Package weather is used to display the forecast of a city.
package weather

// CurrentCondition will display the current condition of provided city.
var CurrentCondition string
// CurrentLocation will display the city provided to see the forecast.
var CurrentLocation string

// Forecast is a function that receives city and condition as strings that later return
// a string with human readable forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
