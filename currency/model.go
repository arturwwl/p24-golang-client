package currency

type Currency string

const CurrencyAED = "AED"
const CurrencyAUD = "AUD"
const CurrencyAUR = "AUR"
const CurrencyBGN = "BGN"
const CurrencyCHF = "CHF"
const CurrencyCNY = "CNY"
const CurrencyCZK = "CZK"
const CurrencyEUR = "EUR"
const CurrencyGBP = "GBP"
const CurrencyHUF = "HUF"
const CurrencyINR = "INR"
const CurrencyIDR = "IDR"
const CurrencyJPY = "JPY"
const CurrencyKRW = "KRW"
const CurrencyMKD = "MKD"
const CurrencyMYR = "MYR"
const CurrencyNOK = "NOK"
const CurrencyNZD = "NZD"
const CurrencyPHP = "PHP"
const CurrencyPLN = "PLN"
const CurrencyRON = "RON"
const CurrencyRUB = "RUB"
const CurrencySEK = "SEK"
const CurrencyTHB = "THB"
const CurrencyTRY = "TRY"
const CurrencyUAH = "UAH"
const CurrencyZAR = "ZAR"

func (c Currency) IsValid() bool {
	switch c {
	case CurrencyAED, CurrencyAUD, CurrencyAUR, CurrencyBGN, CurrencyCHF, CurrencyCNY, CurrencyCZK, CurrencyEUR, CurrencyGBP, CurrencyHUF, CurrencyINR, CurrencyIDR, CurrencyJPY, CurrencyKRW, CurrencyMKD, CurrencyMYR, CurrencyNOK, CurrencyNZD, CurrencyPHP, CurrencyPLN, CurrencyRON, CurrencyRUB, CurrencySEK, CurrencyTHB, CurrencyTRY, CurrencyUAH, CurrencyZAR:
		return true
	}
	return false
}

func (c Currency) ToString() string {
	if c.IsValid() {
		return string(c)
	}
	panic("invalid currency")
}
