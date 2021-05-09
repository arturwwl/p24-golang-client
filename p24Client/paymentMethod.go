package p24Client

import (
	"encoding/json"
	"fmt"
	p24Error "github.com/arturwwl/p24-golang-client/error"
	"github.com/arturwwl/p24-golang-client/model/paymentMethod"
	p24Path "github.com/arturwwl/p24-golang-client/path"
)

func (c *P24Client) GetPaymentMethods(lang string) (*paymentMethod.PaymentMethodList, p24Error.P24Error) {
	bodyBytes, err := c.MakeRequest("GET", fmt.Sprintf(p24Path.PaymentMethods, lang), nil)
	if err != nil {
		return nil, err
	}

	listO := paymentMethod.PaymentMethodList{}
	//omit error, because struct is not unified in p24
	_ = json.Unmarshal(bodyBytes, &listO)

	return &listO, nil
}
