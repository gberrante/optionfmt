# optionfmt

A Go package for parsing and formatting financial option contracts. This package provides utilities to parse various option contract formats and extract key information like symbol, strike price, expiration date, and option type.

## Features

- **Parse option contracts** from standardized formats
- **Extract contract details** including:
  - Underlying symbol
  - Strike price
  - Expiration date
  - Option type (Call or Put)
- **Support for multiple formats**:
  - OSI
  - DAS
  - EUREX Infront/T7
- **Expiration date handling** with customizable formatting
- **Type-safe** Go implementation

## Installation

```bash
go get github.com/gberrante/optionfmt
```

## Usage

### OSI Format
```go
package main

import (
	"fmt"
	"log"
	"optionfmt"
)

func main() {
	var contract optionfmt.OptionContract
	if !contract.ParseOSI("AAPL  261120C00150000") {
		log.Fatal("Failed to parse OSI string")
	}
	
	fmt.Println("Symbol:", contract.Symbol())           // Output: AAPL
	fmt.Println("Strike:", contract.StrikePrice())      // Output: 150.0
	fmt.Println("Type:", contract.Type())               // Output: C
	fmt.Println("Expiration:", contract.Expiration())   // Output: 2026-11-20
}
```

### EUREX Infront Format
```go
var contract optionfmt.OptionContract
contract.ParseEUREXInfront("DBK XEUR C 28 12/21 2")

fmt.Println("Symbol:", contract.Symbol())           // Output: DBK
fmt.Println("Strike:", contract.StrikePrice())      // Output: 28.0
fmt.Println("Expiration:", contract.Expiration())   // Output: 2021-12-17
```

### EUREX T7 Format
```go
var contract optionfmt.OptionContract
contract.ParseEUREXT7("PROD FI 20260917 SM ES C 28.00 CNG")

fmt.Println("Symbol:", contract.Symbol())           // Output: PROD
fmt.Println("Strike:", contract.StrikePrice())      // Output: 28.0
fmt.Println("Type:", contract.Type())               // Output: C
fmt.Println("Expiration:", contract.Expiration())   // Output: 2026-09-17
```

### Bovespa Format
```go
var contract optionfmt.OptionContract
contract.ParseBovespa("BBDCF26")

fmt.Println("Symbol:", contract.Symbol())           // Output: BBDC
fmt.Println("Strike:", contract.StrikePrice())      // Output: 26.0
fmt.Println("Type:", contract.Type())               // Output: C
fmt.Println("Expiration:", contract.Expiration())   // Output: 2026-06-15
```

## API Reference

### OptionContract

The main struct representing an option contract.

#### Methods

- `Symbol() string` - Returns the underlying asset symbol
- `String() string` - Returns the raw string representation
- `Expiration() time.Time` - Returns the expiration date as time.Time (23:59:59 UTC)
- `ExpirationString(format string) string` - Returns formatted expiration date
- `Type() OptionType` - Returns the option type (Call or Put)

### OptionType

Constants for option types:
- `Call` - Call option (value: "C")
- `Put` - Put option (value: "P")

## Testing

Run the test suite:

```bash
go test ./...
```

## License

MIT
