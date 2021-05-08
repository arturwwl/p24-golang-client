package p24Client

import (
	"encoding/json"
	"github.com/arturwwl/p24-golang-client/config"
	"github.com/arturwwl/p24-golang-client/model/testAccess"
	p24Path "github.com/arturwwl/p24-golang-client/path"
)

type P24Client struct {
	Config config.Config
}

func New(configPath string) (client *P24Client, err error) {
	client = new(P24Client)
	client.Config, err = config.LoadConfig(configPath)

	return
}

func (c *P24Client) TestConnection() (*testAccess.TestAccess, error) {
	bodyText, err := c.MakeRequest("GET", p24Path.TestAccess, nil)
	if err != nil {
		return nil, err
	}

	response := &testAccess.TestAccess{}
	_ = json.Unmarshal(bodyText, response)
	return response, nil
}
