package main

import (
	"fmt"
	"math"
	"time"
)

const minPrice float64 = 99.0
const maxPrice float64 = 20000

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		price = minPrice
	}
	if price > maxPrice {
		price = maxPrice
	}
	return price
}

const pricePerKm float64 = 10.0
const pricePerMinute float64 = 2.0

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(tp TripParameters) float64 {
	return tp.Distance*pricePerKm + tp.Duration*pricePerMinute
}

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	var multiplier float64 = 1.0
	switch weather.Condition {
	case Rain:
		multiplier += 0.125
	case HeavyRain:
		multiplier += 0.2
	case Snow:
		multiplier += 0.15
	}
	if weather.WindSpeed > 15 {
		multiplier += 0.1
	}
	return multiplier
}

func GetTimeMultiplier(t time.Time) float64 {
	hour := t.Hour()
	isWeekend := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday // Проверка, что сегодня суббота или воскресенье (выходные)

	switch {
	case hour >= 0 && hour < 5:
		return 1.5 // Ночной тариф
	case hour >= 7 && hour < 10 && !isWeekend:
		return 1.3 // Утренний час пик
	case isWeekend:
		return 1.2 // Выходные
	default:
		return 1.0
	}
}

type TrafficClient interface {
	GetTrafficLevel(lat, lng float64) int // 1–5
}

func GetTrafficMultiplier(trafficLevel int) float64 {
	return 1.0 + float64(trafficLevel-1)*0.1
}

type PriceCalculator struct {
	TrafficClient TrafficClient
}

type RealTrafficClient struct{}

func (c *RealTrafficClient) GetTrafficLevel(lat, lng float64) int {
	return 3 // Константное значение в нашем примере, в реальности оно будет вычисляться сервисом Яндекс Карты
}

func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather WeatherData, lat, lng float64) float64 {
	base := CalculateBasePrice(trip)
	timeMult := GetTimeMultiplier(now)
	weatherMult := GetWeatherMultiplier(weather)
	trafficMult := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return ApplyPriceLimits(math.Round(finalPrice))
}

func main() {
	calculator := PriceCalculator{
		TrafficClient: &RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
	}

	price := calculator.CalculatePrice(
		TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		WeatherData{Snow, 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)

}
