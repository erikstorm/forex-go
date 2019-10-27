package main

// FxFn - The function of choice
type FxFn string

const (
	FxDaily        FxFn = "FX_DAILY"
	FxWeekly       FxFn = "FX_WEEKLY"
	FxMonthly      FxFn = "FX_MONTHLY"
	FxExchangeRate FxFn = "CURRENCY_EXCHANGE_RATE"
	FxIntraDay     FxFn = "FX_INTRADAY"
)

// FxIntraDayInterval - Time interval between two consecutive data points in the time series.
type FxIntraDayInterval string

const (
	FxIntraDay1Min  FxIntraDayInterval = "1min"
	FxIntraDay5Min  FxIntraDayInterval = "5min"
	FxIntraDay15Min FxIntraDayInterval = "15min"
	FxIntraDay30Min FxIntraDayInterval = "30min"
	FxIntraDay60Min FxIntraDayInterval = "60min"
)

// QueryStruct -
type QueryStruct struct {
	FromSymbol   string
	ToSymbol     string
	FromCurrency string
	ToCurrency   string
	Interval     string
}
