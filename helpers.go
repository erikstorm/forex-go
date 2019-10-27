package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func makeRequest(f FxFn, q QueryStruct) ([]byte, error) {
	resp, err := http.Get(getAPIURL(f, q))
	if err != nil {
		log.Println("Error while fetching from api")
		defer resp.Body.Close()
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response")
		return nil, err
	}
	return body, nil
}

func getAPIURL(function FxFn, q QueryStruct) string {
	buf := bytes.Buffer{}
	buf.WriteString(basePath)
	buf.WriteString(fmt.Sprint("?function=", function))
	if q.FromSymbol != "" {
		buf.WriteString(fmt.Sprint("&from_symbol=", q.FromSymbol))
	}
	if q.ToSymbol != "" {
		buf.WriteString(fmt.Sprint("&to_symbol=", q.ToSymbol))
	}
	if q.Interval != "" {
		buf.WriteString(fmt.Sprint("&interval=", q.Interval))
	}
	if q.FromCurrency != "" {
		buf.WriteString(fmt.Sprint("&from_currency=", q.FromCurrency))
	}
	if q.ToCurrency != "" {
		buf.WriteString(fmt.Sprint("&to_currency=", q.ToCurrency))
	}
	buf.WriteString(fmt.Sprint("&apikey=", getEnv("ALPHAVANTAGE_API_KEY", "")))
	log.Println("Calling")
	log.Println(buf.String())
	return buf.String()
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
