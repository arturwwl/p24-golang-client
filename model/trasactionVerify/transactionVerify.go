package trasactionVerify

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/arturwwl/p24-golang-client/currency"

	_ "crypto/sha512"
)

type TransactionVerify struct {
	PosID      uint64            `json:"posId"`
	SessionID  string            `json:"sessionId"`
	Amount     uint64            `json:"amount"`
	Currency   currency.Currency `json:"currency"`
	OrderID    *uint64           `json:"orderId"`
	MerchantID *uint64           `json:"merchantID"`
	Sign       string            `json:"sign"`
}

func (tv *TransactionVerify) CreateSign(CRC string) {
	jsonString := fmt.Sprintf(`{"sessionId":"%s","merchantId":%d,"amount":%d,"currency":"%s","crc":"%s"}`, tv.SessionID, *tv.MerchantID, tv.Amount, tv.Currency.ToString(), CRC)
	tv.calc(jsonString)
}

func (tv *TransactionVerify) CalcSign(CRC string) {
	jsonString := fmt.Sprintf(`{"sessionId":"%s","orderId":%d,"amount":%d,"currency":"%s","crc":"%s"}`, tv.SessionID, *tv.OrderID, tv.Amount, tv.Currency.ToString(), CRC)
	tv.calc(jsonString)
}

func (tv *TransactionVerify) calc(jsonString string) {
	sha384 := crypto.SHA384.New()
	sha384.Write([]byte(jsonString))
	signBytes := sha384.Sum(nil)

	tv.Sign = hex.EncodeToString(signBytes)
}

func (tv *TransactionVerify) ToJSON() string {
	b, err := json.Marshal(tv)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func FromJSON(jsonString string) (*TransactionVerify, error) {
	var tVerify TransactionVerify
	err := json.Unmarshal([]byte(jsonString), &tVerify)
	if err != nil {
		return nil, err
	}

	return &tVerify, nil
}
