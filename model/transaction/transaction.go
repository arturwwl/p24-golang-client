package transaction

import "github.com/arturwwl/p24-golang-client/currency"

type Transaction struct {
	MerchantID  uint64            `json:"merchantId"`          //ID Sklepu
	OrderID     *uint64           `json:"orderId"`             //ID Sklepu
	PosID       uint64            `json:"posId"`               //ID Sklepu (domyślnie ID Sprzedawcy)
	SessionID   string            `json:"sessionId"`           //Unikalny identyfikator z systemu sprzedawcy
	Amount      uint64            `json:"amount"`              //Kwota transakcji wyrażona w groszach, np. 1.23 PLN = 123
	Currency    currency.Currency `json:"currency"`            //Waluta, wartość zgodna z ISO np. PLN
	Description string            `json:"description"`         //Opis transakcji
	Email       string            `json:"email"`               //Email Klienta
	ClientName  *string           `json:"client,omitempty"`    //Imię i nazwisko Klienta
	Address     *string           `json:"address,omitempty"`   //Adres Klienta
	Zip         *string           `json:"zip,omitempty"`       //Kod pocztowy Klienta
	City        *string           `json:"city,omitempty"`      //Miasto Klienta
	Country     string            `json:"country"`             //"PL"
	Phone       *string           `json:"phone,omitempty"`     //Telefon klienta w formacie 481321132123
	Language    string            `json:"language"`            //"pl",
	Method      *uint8            `json:"method,omitempty"`    //Identyfikator metody płatności. Lista metod płatności widoczna w panelu lub dostępna przez API
	UrlReturn   string            `json:"urlReturn"`           //Adres powrotny po zakończeniu transakcji
	UrlStatus   *string           `json:"urlStatus,omitempty"` //Adres do przekazania statusu transakcji
	TimeLimit   *uint16           `json:"timeLimit,omitempty"` //Limit czasu na wykonanie transakcji, 0 - brak limitu, maks. 99 (w minutach)
	Channel     *uint8            `json:"channel,omitempty"`
	/*  1 - karty, 2 - przelewy, 4 - przelew tradycyjny, 8 - N/A, 16 - wszystkie 24/7 – udostępnia wszystkie metody płatności, 32 - użyj przedpłatę, 64 – tylko metody pay-by-link, 128 – formy ratalne, 256 – wallety Aby uruchomić poszczególne kanały, nalezy zsumowac ich wartości.
	Przykład: karty i przelew tradycyjny: channel=5
	WaitForResult bool `json:"waitForResult"`
	*/
	RegulationAccept *bool `json:"regulationAccept,omitempty"`
	/*
		Akceptacja regulaminu Przelewy24:
		false – wyświetl zgodę na stronie p24 (domyślna),
		true – akceptacja dokonana, nie wyświetlaj.
		W przypadku wysyłania parametru „true”, na stronie Partnera musi znaleźć się zgoda o treści: „Oświadczam, że zapoznałem się z regulaminem i obowiązkiem informacyjnym serwisu Przelewy24”.
		Pod słowami regulamin i obowiązek informacyjny musi być link do stron z tymi dokumentami. Checkbox nie może być odgórnie zaznaczony.
	*/
	ShippingPrice     *uint64                 `json:"shipping,omitempty"`      //Koszt dostawy/wysyłki
	BankTransferLabel *string                 `json:"transferLabel,omitempty"` //Opis przekazywany do tytułu przelewu
	MobileLib         *uint8                  `json:"mobileLib,omitempty"`     //Przesłanie tego parametru jest niezbędne przy wykorzystaniu bibliotek SDK. W mobileLib należy przesłać wartość 1, natomiast w parametrze sdkVersion należy wskazać wersję biblioteki, z której chcemy skorzystać.
	SDKVersion        *string                 `json:"sdkVersion,omitempty"`    //Wersja bibliotek mobilnych. Określa czy transakcja jest mobilna.
	Sign              string                  `json:"sign"`                    //Suma kontrolna parametrów
	Encoding          *string                 `json:"encoding,omitempty"`      //System kodowania przesyłanych znaków: ISO-8859-2, UTF-8, Windows-1250
	MethodRegID       *string                 `json:"methodRefId,omitempty"`   //Specjalny parametr wymagany dla niektórych procesów płatności, np. BLIK i Karty one-click.
	Cart              *map[string]interface{} `json:"cart,omitempty"`          //Koszyk
	Additional        *map[string]interface{} `json:"additional,omitempty"`    //Zbiór dodatkowych danych nt. transakcji i płatnika
}

type Created struct {
	Data         *CreatedData
	ResponseCode uint64
}

type CreatedData struct {
	Token string
}

type Verify struct {
	Data         *VerifyData
	ResponseCode uint64
}

type VerifyData struct {
	Status string
}
