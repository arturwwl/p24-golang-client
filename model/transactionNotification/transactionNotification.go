package transactionNotification

import (
	"fmt"
	"github.com/arturwwl/p24-golang-client/currency"
	"github.com/arturwwl/p24-golang-client/hash"

	_ "crypto/sha512"
)

type TransactionNotification struct {
	MerchantID   uint64            `json:"merchantId"`
	PosID        uint64            `json:"posId"`
	SessionID    string            `json:"sessionId"`
	Amount       uint64            `json:"amount"`
	OriginAmount uint64            `json:"originAmount"`
	Currency     currency.Currency `json:"currency"`
	OrderID      uint64            `json:"orderId"`
	MethodId     uint64            `json:"methodId"`
	Statement    string            `json:"statement"`
	Sign         string            `json:"sign"`
}

func (tv *TransactionNotification) IsSignValid(CRC string) bool {
	jsonString := fmt.Sprintf(`{"merchantId":%d,"posId":%d,"sessionId":"%s","amount":%d,"originAmount":%d,"currency":"%s","orderId":%d,"methodId":%d,"statement":"%s","crc":"%s"}`, tv.MerchantID, tv.PosID, tv.SessionID, tv.Amount, tv.OriginAmount, tv.Currency.ToString(), tv.OrderID, tv.MethodId, tv.Statement, CRC)
	return tv.Sign == hash.CalcSha384(jsonString)
}
