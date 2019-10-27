# forex-go

Trying out golang, didn't feel like writing another hello world.

Serves a few routes for realtime and historical forex (FX) rates.

Uses in-memory cache to limit requests and not hit the free tier limit at alphavantage. Requests or data are not persisted outside of the cache.
The standard API call frequency is 5 calls per minute
and 500 calls per day, so we cache all requests for 10 minutes not to hit this limit too soon. Hence, data is not realtime.

## Get started

- Golang installed and configured
- Alphavantage api key, set in project `.env`
- run `go build && ./forex-go`


## Api

#### /forex/daily

Returns the realtime exchange rate for any pair of digital currency (e.g., Bitcoin) and physical currency (e.g., USD). Data returned for physical currency (Forex) pairs also include realtime bid and ask prices.

##### params
```
fromSymbol   from symbol
toSymbol     to symbol

# example
GET http://localhost:8080/forex/daily?fromSymbol=USD&toSymbol=EUR
```

#### /forex/weekly

Returns the weekly time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.

The latest data point is the price information for the week (or partial week) containing the current trading day, updated realtime.

##### params
```
fromSymbol   from symbol
toSymbol     to symbol

# example
GET http://localhost:8080/forex/weekly?fromSymbol=USD&toSymbol=EUR
```
#### /forex/monthly

Returns the monthly time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.

The latest data point is the prices information for the month (or partial month) containing the current trading day, updated realtime.

##### params
```
fromSymbol   from symbol
toSymbol     to symbol

# example
GET http://localhost:8080/forex/monthly?fromSymbol=USD&toSymbol=EUR
```
#### /forex/intraday

Returns intraday time series (timestamp, open, high, low, close) of the FX currency pair specified, updated realtime.

##### params
```
fromSymbol   from symbol
toSymbol     to symbol
interval     time interval between time series (1min, 5min, 15min, 30min 60min)

# example
GET http://localhost:8080/forex/intraday?fromSymbol=USD&toSymbol=EUR&interval=60min
```
#### /forex/exchange-rate

Returns the realtime exchange rate for any pair of digital currency (e.g., Bitcoin) and physical currency (e.g., USD). Data returned for physical currency (Forex) pairs also include realtime bid and ask prices.

##### params
```
fromCurrency   from currency
toCurrency     to currency

# example
GET http://localhost:8080/forex/exchange-rate?fromCurrency=USD&toCurrency=EUR
```
