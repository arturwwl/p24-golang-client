package p24Client

import (
	"encoding/json"
	"fmt"
	p24error "github.com/arturwwl/p24-golang-client/error"
	"github.com/arturwwl/p24-golang-client/model/transaction"
	"github.com/arturwwl/p24-golang-client/model/transactionVerify"
	p24Path "github.com/arturwwl/p24-golang-client/path"
)

func (c *P24Client) GenerateTransactionClientLink(token string) string {
	return fmt.Sprintf(getUrl(c.Config.IsProd, p24Path.TransactionClientLink), token)
}

func (c *P24Client) RegisterTransaction(transactionData *transaction.Transaction) (*transaction.Created, p24error.P24Error) {

	if true { //tVerify does not need to be accessed later
		tVerify := &transactionVerify.TransactionVerify{
			PosID:      transactionData.PosID,
			SessionID:  transactionData.SessionID,
			Amount:     transactionData.Amount,
			Currency:   transactionData.Currency,
			OrderID:    nil,
			MerchantID: &transactionData.MerchantID,
			Sign:       "", //don't forget create it!
		}
		tVerify.CreateSign(c.Config.CrcKey)

		transactionData.Sign = tVerify.Sign
	}
	bodyBytes, err := c.MakeRequest("POST", p24Path.TransactionRegister, transactionData)
	if err != nil {
		return nil, err
	}

	tCreated := &transaction.Created{}
	//omit error
	_ = json.Unmarshal(bodyBytes, tCreated)

	return tCreated, nil
}
