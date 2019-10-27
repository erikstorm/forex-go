package main

import (
	"net/http"
)

// GetDaily - returns the daily time series (timestamp, open, high, low, close) of the FX
// currency pair specified, updated realtime.
func GetDaily(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	query := QueryStruct{
		FromSymbol: queryValues.Get("fromSymbol"),
		ToSymbol:   queryValues.Get("toSymbol"),
	}
	if query.FromSymbol == "" || query.ToSymbol == "" {
		return
	}
	body, err := makeRequest(FxDaily, query)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

// GetWeekly - returns the weekly time series (timestamp, open, high, low, close) of the FX
// currency pair specified, updated realtime.
//
// The latest data point is the price information for the week (or partial week) containing
// the current trading day, updated realtime.
func GetWeekly(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	query := QueryStruct{
		FromSymbol: queryValues.Get("fromSymbol"),
		ToSymbol:   queryValues.Get("toSymbol"),
	}
	if query.FromSymbol == "" || query.ToSymbol == "" {
		return
	}
	body, err := makeRequest(FxWeekly, query)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

// GetMonthly - returns the monthly time series (timestamp, open, high, low, close) of the FX
// currency pair specified, updated realtime.
//
// The latest data point is the prices information for the month (or partial month) containing
// the current trading day, updated realtime.
func GetMonthly(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	query := QueryStruct{
		FromSymbol: queryValues.Get("fromSymbol"),
		ToSymbol:   queryValues.Get("toSymbol"),
	}
	if query.FromSymbol == "" || query.ToSymbol == "" {
		return
	}
	body, err := makeRequest(FxMonthly, query)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

// GetIntraDay - returns intraday time series (timestamp, open, high, low, close) of
// the FX currency pair specified, updated realtime.
func GetIntraDay(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	query := QueryStruct{
		FromSymbol: queryValues.Get("fromSymbol"),
		ToSymbol:   queryValues.Get("toSymbol"),
		Interval:   queryValues.Get("interval"),
	}
	if query.FromSymbol == "" || query.ToSymbol == "" || query.Interval == "" {
		return
	}
	body, err := makeRequest(FxIntraDay, query)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}

// GetCurrencyExchangeRate - returns the realtime exchange rate for any pair of digital currency
// (e.g., Bitcoin) and physical currency (e.g., USD). Data returned for physical currency (Forex)
// pairs also include realtime bid and ask prices.
func GetCurrencyExchangeRate(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	query := QueryStruct{
		FromCurrency: queryValues.Get("fromCurrency"),
		ToCurrency:   queryValues.Get("toCurrency"),
	}
	if query.FromCurrency == "" || query.ToCurrency == "" {
		return
	}

	body, err := makeRequest(FxExchangeRate, query)
	if err != nil {
		panic(err)
	}
	w.Write(body)
}
