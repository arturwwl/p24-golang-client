package p24Client

import (
	"encoding/json"
	p24Error "github.com/arturwwl/p24-golang-client/error"
	"github.com/arturwwl/p24-golang-client/model/transactionNotification"
)

func (c *P24Client) TransactionNotification(bodyBytes []byte) (*transactionNotification.TransactionNotification, p24Error.P24Error) {
	tNoti := &transactionNotification.TransactionNotification{}
	//omit error
	_ = json.Unmarshal(bodyBytes, tNoti)

	if !tNoti.IsSignValid(c.Config.CrcKey) {
		return nil, p24Error.Errorf(p24Error.InvalidRequestSignature)
	}

	return tNoti, nil
}
