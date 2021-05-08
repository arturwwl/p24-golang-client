package p24Client

import (
	"encoding/json"
	"github.com/arturwwl/p24-golang-client/model/transaction"
	"github.com/arturwwl/p24-golang-client/model/trasactionVerify"
	p24Path "github.com/arturwwl/p24-golang-client/path"
)

func (c *P24Client) RegisterTransaction(transactionData *transaction.Transaction) (interface{}, error) {

	if true { //tVerify does not need to be accessed later
		tVerify := &trasactionVerify.TransactionVerify{
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

	listO := transaction.Created{}
	//omit error
	_ = json.Unmarshal(bodyBytes, &listO)

	return &listO, nil
}
