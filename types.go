package optionfmt

import "time"

type OptionType string

const (
	Call OptionType = "C"
	Put  OptionType = "P"
)

type OptionContract struct {
	symbol      string
	strikePrice float64
	expiration  struct {
		Year  int
		Month int
		Day   int
	}
	optionType OptionType
	raw        string
}

// get the underlying symbol of the option contract
func (o OptionContract) Symbol() string {
	return o.symbol
}

// get raw string representation of the option contract
func (o OptionContract) String() string {
	return o.raw
}

// get the expiration date of the option contract
func (o OptionContract) Expiration() time.Time {
	//end of day in UTC
	return time.Date(o.expiration.Year, time.Month(o.expiration.Month), o.expiration.Day, 23, 59, 59, 0, time.UTC)
}

// get the expiration date of the option contract as a formatted string (format can be "2006-01-02", "060102", etc.)
func (o OptionContract) ExpirationString(format string) string {
	return time.Date(o.expiration.Year, time.Month(o.expiration.Month), o.expiration.Day, 23, 59, 59, 0, time.UTC).Format(format)
}

// get the option type of the option contract
func (o OptionContract) Type() OptionType {
	return o.optionType
}

// get the option type of the option contract as a string ["C" or "P"]
func (o OptionContract) TypeString() string {
	return string(o.optionType)
}

// get the strike price of the option contract
func (o OptionContract) StrikePrice() float64 {
	return o.strikePrice
}
