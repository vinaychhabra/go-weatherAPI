# Go-WeatherAPI

 Go-WeatherAPI is a RESTful API server built with Go. It provides current weather information for a specified city by integrating with an external weather API. This project demonstrates the use of Go for backend development, including handling HTTP requests, integrating with external APIs, and structuring a Go application.

## Features

- Fetch current weather data for cities.
- RESTful API supporting GET and POST requests.
- Integration with an external weather API.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installing

A step-by-step guide to setting up a development environment.

1. Clone the repository

```bash
git clone https://github.com/vinaychhabra/go-weatherAPI.git
cd go-weatherAPI
go run main.go

2. Usage
To fetch weather data:

GET Request: GET /city?name=Toronto
POST Request: Send a POST request to /city with a JSON body: {"name": "Toronto"}

