package p24Client

import (
	"encoding/json"
	p24Error "github.com/arturwwl/p24-golang-client/error"
	"github.com/arturwwl/p24-golang-client/model/transaction"
	"github.com/arturwwl/p24-golang-client/model/transactionNotification"
	"github.com/arturwwl/p24-golang-client/model/transactionVerify"
	p24Path "github.com/arturwwl/p24-golang-client/path"
)

func (c *P24Client) VerifyTransaction(transactionNoti *transactionNotification.TransactionNotification) (*transaction.Verify, p24Error.P24Error) {

	tVerify := &transactionVerify.TransactionVerify{
		MerchantID: &transactionNoti.MerchantID,
		PosID:      transactionNoti.PosID,
		SessionID:  transactionNoti.SessionID,
		Amount:     transactionNoti.Amount,
		Currency:   transactionNoti.Currency,
		OrderID:    &transactionNoti.OrderID,
		Sign:       "", //don't forget calc it!
	}
	tVerify.CalcSign(c.Config.CrcKey)

	bodyBytes, err := c.MakeRequest("PUT", p24Path.TransactionVerify, tVerify)
	if err != nil {
		return nil, err
	}

	tVerifyResponse := &transaction.Verify{}
	//omit error
	_ = json.Unmarshal(bodyBytes, tVerifyResponse)

	return tVerifyResponse, nil
}
