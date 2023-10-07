# Weather Report CLI

A command-line tool for retrieving and displaying weather information for a specified location using the WeatherAPI.

## Features

- **Weather Information:** Retrieve and display current weather conditions and a one-hour weather forecast for a specified location.
- **Location Customization:** You can specify the location as a command-line argument, or if none is provided, it defaults to "Berhampur."

## Getting Started



   ```bash
    # Navigate to the project directory:
    cd 6-weather
    # Build the Go application:
    go build -o weather
    # Run the app
    # Without specifying a location (defaults to Berhampur):
    ./weather
    Berhampur, India: 26C, Cloudy
    # Specify a location (e.g., Delhi):
    ./weather Delhi
    Delhi, Canada: 12C, Sunny
    18:00 - 13C, 0%, Partly cloudy
    19:00 - 12C, 76%, Patchy rain possible
    20:00 - 11C, 87%, Patchy rain possible
    21:00 - 11C, 0%, Cloudy
    22:00 - 10C, 0%, Partly cloudy
    23:00 - 10C, 76%, Patchy rain possible
    00:00 - 10C, 0%, Partly cloudy
    01:00 - 9C, 80%, Patchy rain possible
    02:00 - 9C, 63%, Patchy rain possible
    03:00 - 9C, 0%, Partly cloudy

   ```


## Dependencies

- This project relies on the [WeatherAPI](http://api.weatherapi.com) to retrieve weather data.

