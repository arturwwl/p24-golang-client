package p24Client

import (
	"encoding/json"
	"github.com/arturwwl/p24-golang-client/model/transaction"
	"github.com/arturwwl/p24-golang-client/model/trasactionVerify"
	p24Path "github.com/arturwwl/p24-golang-client/path"
)

func (c *P24Client) VerifyTransaction(transactionData *transaction.Transaction) (*transaction.Verify, error) {

	if true { //tVerify does not need to be accessed later
		tVerify := &trasactionVerify.TransactionVerify{
			PosID:      transactionData.PosID,
			SessionID:  transactionData.SessionID,
			Amount:     transactionData.Amount,
			Currency:   transactionData.Currency,
			OrderID:    transactionData.OrderID,
			MerchantID: &transactionData.MerchantID,
			Sign:       "", //don't forget calc it!
		}
		tVerify.CalcSign(c.Config.CrcKey)

		transactionData.Sign = tVerify.Sign
	}
	bodyBytes, err := c.MakeRequest("PUT", p24Path.TransactionVerify, transactionData)
	if err != nil {
		return nil, err
	}

	tVerify := &transaction.Verify{}
	//omit error
	_ = json.Unmarshal(bodyBytes, tVerify)

	return tVerify, nil
}
