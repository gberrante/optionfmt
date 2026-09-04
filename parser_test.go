package optionfmt

import "testing"

func TestOSIParser(t *testing.T) {
	osiString := "AAPL  261120C00150000"
	var contract OptionContract
	ok := contract.ParseOSI(osiString)
	if !ok {
		t.Errorf("Failed to parse OSI string: %s", osiString)
	}

	if contract.Symbol() != "AAPL" {
		t.Errorf("Expected symbol 'AAPL', got '%s'", contract.Symbol())
	}

	if contract.Type() != Call {
		t.Errorf("Expected option type 'C', got '%s'", contract.Type())
	}

	if contract.StrikePrice() != 150.0 {
		t.Errorf("Expected strike price 150.0, got %f", contract.StrikePrice())
	}

	if contract.Expiration().Year() != 2026 || contract.Expiration().Month() != 11 || contract.Expiration().Day() != 20 {
		t.Errorf("Expected expiration date 2026-11-20, got %s", contract.Expiration().Format("2006-01-02"))
	}

	wrongOSIString := "AAPL  261120C0015000" // Missing one digit in strike price
	ok = contract.ParseOSI(wrongOSIString)
	if ok {
		t.Errorf("Expected failure to parse invalid OSI string: %s", wrongOSIString)
	}
}

func TestDASParser(t *testing.T) {
	dasString := "+MSFT^G7D300"
	var contract OptionContract
	ok := contract.ParseDAS(dasString)
	if !ok {
		t.Errorf("Failed to parse DAS string: %s", dasString)
	}

	if contract.Symbol() != "MSFT" {
		t.Errorf("Expected symbol 'MSFT', got '%s'", contract.Symbol())
	}

	if contract.Type() != Call {
		t.Errorf("Expected option type 'C', got '%s'", contract.Type())
	}

	if contract.StrikePrice() != 300.0 {
		t.Errorf("Expected strike price 300.0, got %f", contract.StrikePrice())
	}

	if contract.Expiration().Year() != 2026 || contract.Expiration().Month() != 7 || contract.Expiration().Day() != 13 {
		t.Errorf("Expected expiration date 2026-07-13, got %s", contract.Expiration().Format("2006-01-02"))
	}

}

func TestEUREXInfrontParser(t *testing.T) {
	eurexString := "DBK XEUR C 28 12/21 2"
	var contract OptionContract
	ok := contract.ParseEUREXInfront(eurexString)
	if !ok {
		t.Errorf("Failed to parse EUREX Infront string: %s", eurexString)
	}

	if contract.Symbol() != "DBK" {
		t.Errorf("Expected symbol 'DBK', got '%s'", contract.Symbol())
	}

	if contract.Type() != Call {
		t.Errorf("Expected option type 'C', got '%s'", contract.Type())
	}

	if contract.StrikePrice() != 28.0 {
		t.Errorf("Expected strike price 28.0, got %f", contract.StrikePrice())
	}

	if contract.Expiration().Year() != 2021 || contract.Expiration().Month() != 12 || contract.Expiration().Day() != 17 {
		t.Errorf("Expected expiration date 2021-12-17, got %s", contract.Expiration().Format("2006-01-02"))
	}
}

func TestEUREXT7Parser(t *testing.T) {
	eurexString := "PROD FI 20260917 SM ES C 28.00 CNG"

	var contract OptionContract
	ok := contract.ParseEUREXT7(eurexString)
	if !ok {
		t.Errorf("Failed to parse EUREX T7 string: %s", eurexString)
	}

	if contract.Symbol() != "PROD" {
		t.Errorf("Expected symbol 'PROD', got '%s'", contract.Symbol())
	}

	if contract.Type() != Call {
		t.Errorf("Expected option type 'C', got '%s'", contract.Type())
	}

	if contract.StrikePrice() != 28.00 {
		t.Errorf("Expected strike price 28.00, got %f", contract.StrikePrice())
	}

	if contract.Expiration().Year() != 2026 || contract.Expiration().Month() != 9 || contract.Expiration().Day() != 17 {
		t.Errorf("Expected expiration date 2026-09-17, got %s", contract.Expiration().Format("2006-01-02"))
	}

}

func TestBovespaParser(t *testing.T) {
	bovespaString := "BBDCF26"
	var contract OptionContract
	ok := contract.ParseBovespa(bovespaString)

	if !ok {
		t.Errorf("Failed to parse Bovespa string: %s", bovespaString)
	}

	if contract.Symbol() != "BBDC" {
		t.Errorf("Expected symbol 'BBDC', got '%s'", contract.Symbol())
	}

	if contract.Type() != Call {
		t.Errorf("Expected option type 'C', got '%s'", contract.Type())
	}

	if contract.StrikePrice() != 26.0 {
		t.Errorf("Expected strike price 26.0, got %f", contract.StrikePrice())
	}

	if contract.Expiration().Year() != 2026 || contract.Expiration().Month() != 6 || contract.Expiration().Day() != 15 {
		t.Errorf("Expected expiration date 2026-06-15, got %s", contract.Expiration().Format("2006-01-02"))
	}
}
