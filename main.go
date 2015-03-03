package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
)

// http://api.openweathermap.org/data/2.5/weather?q=Melbourne,au

func main() {
    http.HandleFunc("/", weatherHandler)
    http.ListenAndServe(":5000", nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
    body, err := getWeatherResponseBody()

    melbourne := City{}
    err = json.Unmarshal(body, &melbourne)

    if err != nil {
        panic(err)
    }
    fmt.Fprintf(w, "The weather in %v is %v\n", melbourne.Name, melbourne.Weather.NormalisedCurrentTemp())
}

func getWeatherResponseBody() ([]byte, error) {
    resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Melbourne,au")
    if err != nil {
        fmt.Printf("Error getting weather %v", err)
        return []byte(""), err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading weather $v", err)
        return []byte(""), err
    }
    defer resp.Body.Close()

    return body, nil
}

type City struct {
    Weather Weather `json:"main"`
    Name    string  `json:"name"`
}

type Weather struct {
    CurrentTemp float64 `json:"temp"`
    MaxTemp     float64 `json:"temp_max"`
}

func (w Weather) NormalisedCurrentTemp() float64 {
    return w.CurrentTemp - 273.15
}

