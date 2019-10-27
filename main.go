package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	cache "github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	memcached, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LRU),
		memory.AdapterWithCapacity(10000000),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(memcached),
		cache.ClientWithTTL(cacheTTL),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/forex/daily", cacheClient.Middleware(loggerMiddleware(http.HandlerFunc(GetDaily)))).Methods("GET")
	router.Handle("/forex/weekly", cacheClient.Middleware(loggerMiddleware(http.HandlerFunc(GetWeekly)))).Methods("GET")
	router.Handle("/forex/monthly", cacheClient.Middleware(loggerMiddleware(http.HandlerFunc(GetMonthly)))).Methods("GET")
	router.Handle("/forex/intraday", cacheClient.Middleware(loggerMiddleware(http.HandlerFunc(GetIntraDay)))).Methods("GET")
	router.Handle("/forex/exchange-rate", cacheClient.Middleware(loggerMiddleware(http.HandlerFunc(GetCurrencyExchangeRate)))).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func loggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		log.Println(r.URL)
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}
