package test

import (
	"github.com/arturwwl/p24-golang-client/currency"
	"github.com/arturwwl/p24-golang-client/model/transaction"
	"github.com/arturwwl/p24-golang-client/model/trasactionVerify"
	"github.com/arturwwl/p24-golang-client/p24Client"
	"github.com/google/uuid"
	"testing"
)

func TestCalcSignAndToJSON(t *testing.T) {
	orderID := uint64(456)
	tVerify := &trasactionVerify.TransactionVerify{
		//MerchantID: 123,
		PosID:     234,
		SessionID: "abc",
		Amount:    345,
		Currency:  currency.CurrencyPLN,
		OrderID:   &orderID,
	}

	tVerify.CalcSign("test")
}

func TestEncodeAndDecodeTransaction(t *testing.T) {
	orderID := uint64(456)
	tVerify := &trasactionVerify.TransactionVerify{
		//MerchantID: 123,
		PosID:     234,
		SessionID: "abc",
		Amount:    345,
		Currency:  currency.CurrencyPLN,
		OrderID:   &orderID,
	}
	tvJson := tVerify.ToJSON()

	tVerifyFromJSON, err := trasactionVerify.FromJSON(tvJson)
	if err != nil {
		t.Error(err)
	}

	if tVerify != tVerifyFromJSON {
		t.Error("invalid")
	}
}

func TestClient(t *testing.T) {
	client, err := p24Client.New("conf/testsandbox.ini")
	if err != nil {
		t.Fatal(err)
	}

	response, err := client.TestConnection()
	if err != nil {
		t.Fatal(err)
	}
	if response.Data != true {
		t.Fatal("response data not true")
	}
}
func TestPaymentMethods(t *testing.T) {
	client, err := p24Client.New("conf/testsandbox.ini")
	if err != nil {
		t.Fatal(err)
	}

	methods, err := client.GetPaymentMethods("pl")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(methods)
}
func TestCreateTransaction(t *testing.T) {
	client, err := p24Client.New("conf/testsandbox.ini")
	if err != nil {
		t.Fatal(err)
	}
	transactionData := &transaction.Transaction{
		MerchantID:  uint64(client.Config.MerchantID),
		PosID:       uint64(client.Config.PosID),
		SessionID:   uuid.NewString(),
		Amount:      100,
		Currency:    currency.CurrencyPLN,
		Description: "test order",
		Email:       "text@example.com",
		//ClientName:        "Lorem",
		//Address:           "Test 1",
		//Zip:               "00-000",
		//City:              "City",
		Country: "PL",
		//Phone:             "123123123",
		Language: "pl",
		//Method:            0,
		UrlReturn: client.Config.ReturnUrl,
		//UrlStatus:         client.Config.StatusUrl,
		//TimeLimit:         client.Config.TransactionTimeLimit,
		//Channel:           0,
		//RegulationAccept:  false,
		//ShippingPrice:     0,
		//BankTransferLabel: "Testowa transakcja",
		//MobileLib:         0,
		//SDKVersion:        "",
		Sign: "", //filled later
		//Encoding:          "",
		//MethodRegID:       "",
		//Cart:              nil,
		//Additional:        nil,
	}
	methods, err := client.RegisterTransaction(transactionData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(methods)
}
