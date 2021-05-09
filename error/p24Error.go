package p24Error

import "fmt"

type P24Error error

const InvalidRequestSignature = "p24 - invalid request signature"

func Errorf(s string) P24Error {
	return P24Error(fmt.Errorf(s))
}
