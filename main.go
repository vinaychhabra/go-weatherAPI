package main

// WeatherResponse defines the structure of our API response
type WeatherResponse struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"` // We'll use the "Description" from the external API here
}

func main() {
}
