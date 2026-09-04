package optionfmt

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dasMonthMap = map[byte]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'A': 10,
	'B': 11,
	'C': 12,
}

var dasDayMap = map[byte]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
	'G': 16,
	'H': 17,
	'I': 18,
	'J': 19,
	'K': 20,
	'L': 21,
	'M': 22,
	'N': 23,
	'O': 24,
	'P': 25,
	'Q': 26,
	'R': 27,
	'S': 28,
	'T': 29,
	'U': 30,
	'V': 31,
}

var dasYearMap = map[byte]int{
	'0': 2010,
	'1': 2011,
	'2': 2012,
	'3': 2013,
	'4': 2014,
	'5': 2015,
	'6': 2016,
	'7': 2017,
	'8': 2018,
	'9': 2019,
	'A': 2020,
	'B': 2021,
	'C': 2022,
	'D': 2023,
	'E': 2024,
	'F': 2025,
	'G': 2026,
	'H': 2027,
	'I': 2028,
	'J': 2029,
	'K': 2030,
	'L': 2031,
	'M': 2032,
	'N': 2033,
	'O': 2034,
	'P': 2035,
	'Q': 2036,
	'R': 2037,
	'S': 2038,
	'T': 2039,
	'U': 2040,
	'V': 2041,
	'W': 2042,
	'X': 2043,
	'Y': 2044,
	'Z': 2045,
}

type bovespaOptionType struct {
	Month time.Month
	Type  OptionType
}

var bovespaMap = map[byte]bovespaOptionType{
	'A': {Month: time.January, Type: Call},
	'B': {Month: time.February, Type: Call},
	'C': {Month: time.March, Type: Call},
	'D': {Month: time.April, Type: Call},
	'E': {Month: time.May, Type: Call},
	'F': {Month: time.June, Type: Call},
	'G': {Month: time.July, Type: Call},
	'H': {Month: time.August, Type: Call},
	'I': {Month: time.September, Type: Call},
	'J': {Month: time.October, Type: Call},
	'K': {Month: time.November, Type: Call},
	'L': {Month: time.December, Type: Call},
	'M': {Month: time.January, Type: Put},
	'N': {Month: time.February, Type: Put},
	'O': {Month: time.March, Type: Put},
	'P': {Month: time.April, Type: Put},
	'Q': {Month: time.May, Type: Put},
	'R': {Month: time.June, Type: Put},
	'S': {Month: time.July, Type: Put},
	'T': {Month: time.August, Type: Put},
	'U': {Month: time.September, Type: Put},
	'V': {Month: time.October, Type: Put},
	'W': {Month: time.November, Type: Put},
	'X': {Month: time.December, Type: Put},
}

// ParseOSI parses an OSI formatted option string and returns success or failure. It populates the OptionContract struct with the parsed data.
func (o *OptionContract) ParseOSI(osiString string) bool {
	contract := OptionContract{}
	var err error

	if len(osiString) != 21 {
		return false
	}
	contract.raw = osiString
	contract.symbol = strings.TrimSpace(osiString[:6])

	contract.expiration.Year, err = strconv.Atoi(osiString[6:8])
	if err != nil {
		return false
	}
	contract.expiration.Year += 2000
	contract.expiration.Month, err = strconv.Atoi(osiString[8:10])
	if err != nil {
		return false
	}
	contract.expiration.Day, err = strconv.Atoi(osiString[10:12])
	if err != nil {
		return false
	}

	optionTypeChar := osiString[12]
	contract.optionType = OptionType(optionTypeChar)

	strikePriceStr := strings.TrimLeft(osiString[13:], "0")
	strikeNumber, err := strconv.ParseFloat(strikePriceStr, 64)
	if err != nil {
		return false
	}
	contract.strikePrice = strikeNumber / 1000.0

	*o = contract
	return true
}

// ParseDAS parses a DAS formatted option string and returns success or failure. It populates the OptionContract struct with the parsed data.
func (o *OptionContract) ParseDAS(dasString string) bool {
	contract := OptionContract{}
	var err error

	re := `\+([A-Z]+)([\^*])([A-Z0-9]{3})([0-9.]+)`

	matches := regexp.MustCompile(re).FindStringSubmatch(dasString)
	if len(matches) != 5 {
		return false
	}

	contract.raw = dasString
	contract.symbol = matches[1]

	optionTypeChar := matches[2]
	switch optionTypeChar {
	case "^":
		contract.optionType = Call
	case "*":
		contract.optionType = Put
	default:
		return false
	}

	expirationStr := matches[3]
	contract.expiration.Year = dasYearMap[expirationStr[0]]
	contract.expiration.Month = dasMonthMap[expirationStr[1]]
	contract.expiration.Day = dasDayMap[expirationStr[2]]

	strikePriceStr := matches[4]
	strikeNumber, err := strconv.ParseFloat(strikePriceStr, 64)
	if err != nil {
		return false
	}
	contract.strikePrice = strikeNumber

	*o = contract
	return true
}

// ParseEUREXInfront parses a EUREX Infront/Market Data formatted option string and returns success or failure. It populates the OptionContract struct with the parsed data.
func (o *OptionContract) ParseEUREXInfront(eurexString string) bool {
	contract := OptionContract{}
	var err error

	parts := strings.Split(eurexString, " ")

	if len(parts) != 6 {
		return false
	}

	contract.raw = eurexString
	contract.symbol = parts[0]
	//discard parts[1] as it is not needed for the OptionContract struct

	contract.optionType = OptionType(parts[2])

	strikePriceStr := parts[3]
	strikeNumber, err := strconv.ParseFloat(strikePriceStr, 64)
	if err != nil {
		return false
	}
	contract.strikePrice = strikeNumber

	// month/year parsing logic here
	my := strings.Split(parts[4], "/")
	if len(my) != 2 {
		return false
	}

	month, err := strconv.Atoi(my[0])
	if err != nil {
		return false
	}
	year, err := strconv.Atoi(my[1])
	if err != nil {
		return false
	}
	contract.expiration.Year = year + 2000
	contract.expiration.Month = month

	//third friday of the month
	firstOfMonth := time.Date(year+2000, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	weekday := firstOfMonth.Weekday()
	offset := (time.Friday - weekday + 7) % 7
	thirdFriday := firstOfMonth.AddDate(0, 0, int(offset)+14)
	contract.expiration.Day = thirdFriday.Day()

	*o = contract
	return true
}

// ParseEUREXT7 parses a EUREX T7 formatted option string and returns success or failure. It populates the OptionContract struct with the parsed data.
func (o *OptionContract) ParseEUREXT7(eurexString string) bool {
	contract := OptionContract{}
	var err error

	parts := strings.Split(eurexString, " ")

	if len(parts) != 8 {
		return false
	}

	contract.raw = eurexString
	contract.symbol = parts[0]

	expirationDate, err := time.Parse("20060102", parts[2])
	if err != nil {
		return false
	}
	contract.expiration.Year = expirationDate.Year()
	contract.expiration.Month = int(expirationDate.Month())
	contract.expiration.Day = expirationDate.Day()

	contract.optionType = OptionType(parts[5])

	strikePriceStr := parts[6]
	strikeNumber, err := strconv.ParseFloat(strikePriceStr, 64)
	if err != nil {
		return false
	}
	contract.strikePrice = strikeNumber

	*o = contract
	return true
}

// ParseBovespa parses a Bovespa formatted option string and returns success or failure. It populates the OptionContract struct with the parsed data.
func (o *OptionContract) ParseBovespa(bovespaString string) bool {
	contract := OptionContract{}
	var err error

	if len(bovespaString) != 7 {
		return false
	}

	contract.raw = bovespaString
	contract.symbol = bovespaString[:4]

	datetype := bovespaMap[bovespaString[4]]
	contract.optionType = datetype.Type

	currentYear := time.Now().Year()
	contract.expiration.Month = int(datetype.Month)
	contract.expiration.Year = currentYear

	//day is the third monday of the month, suppose current year
	thirdMonday := time.Date(currentYear, time.Month(datetype.Month), 1, 0, 0, 0, 0, time.UTC)
	for thirdMonday.Weekday() != time.Monday {
		thirdMonday = thirdMonday.AddDate(0, 0, 1)
	}
	thirdMonday = thirdMonday.AddDate(0, 0, 14) // move to the third Monday
	contract.expiration.Day = thirdMonday.Day()

	strikePriceStr := bovespaString[5:]
	strikeNumber, err := strconv.ParseFloat(strikePriceStr, 64)
	if err != nil {
		return false
	}
	contract.strikePrice = strikeNumber

	*o = contract
	return true

}
