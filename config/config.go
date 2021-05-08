package config

import (
	"github.com/go-ini/ini"
)

type Config struct {
	ReturnUrl            string
	StatusUrl            string
	PosID                int
	MerchantID           int
	OrderKey             string
	CrcKey               string
	ApiKey               string
	IsProd               bool
	WaitForResult        bool
	TransactionTimeLimit uint16
}

func LoadConfig(configPath string) (p24Config Config, err error) {
	err = ini.MapTo(&p24Config, configPath)
	return
}
