package test

import (
	"github.com/arturwwl/p24-golang-client/currency"
	"github.com/arturwwl/p24-golang-client/model/transaction"
	"github.com/arturwwl/p24-golang-client/model/transactionVerify"
	"github.com/arturwwl/p24-golang-client/p24Client"
	"github.com/google/uuid"
	"testing"
)

func TestCalcSignAndToJSON(t *testing.T) {
	orderID := uint64(456)
	tVerify := &transactionVerify.TransactionVerify{
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
	tVerify := &transactionVerify.TransactionVerify{
		//MerchantID: 123,
		PosID:     234,
		SessionID: "abc",
		Amount:    345,
		Currency:  currency.CurrencyPLN,
		OrderID:   &orderID,
	}
	tvJson := tVerify.ToJSON()

	tVerifyFromJSON, err := transactionVerify.FromJSON(tvJson)
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
		UrlStatus: &client.Config.StatusUrl,
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
	response, err := client.RegisterTransaction(transactionData)
	if err != nil {
		t.Fatal(err)
	}

	if response.Data == nil || response.Data.Token == "" {
		t.Fatal("invalid response data")
	}

	link := client.GenerateTransactionClientLink(response.Data.Token)
	if link == "" {
		t.Fatal("client link cannot be null")
	}
}

var isManual bool

func TestStatus(t *testing.T) {
	if !isManual {
		t.Log("this can be run only with manually filled")
		t.SkipNow()
	}

	exampleString := `{"merchantId":123123,"posId":123123,"sessionId":"11223344-aaaa-bbbb-cccc-445566778899","amount":100,"originAmount":100,"currency":"PLN","orderId":111222333,"methodId":94,"statement":"p24-A11-A11-A11","sign":"596af9bc39271b4cfdab45937"}`
	exampleBytes := []byte(exampleString)

	client, err := p24Client.New("conf/testsandbox.ini")
	if err != nil {
		t.Fatal(err)
	}

	transactionNoti, err := client.TransactionNotification(exampleBytes)
	if err != nil {
		t.Fatal(err)
	}

	transactionNotiResp, err := client.VerifyTransaction(transactionNoti)
	if err != nil {
		t.Fatal(err)
	}

	if transactionNotiResp.Data == nil || transactionNotiResp.Data.Status == "" {
		t.Fatal("invalid response data")
	}

	if transactionNotiResp.Data.Status != "success" {
		t.Fatal("transaction status is not success")
	}

}
