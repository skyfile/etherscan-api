package etherscan

type geth struct {
	Cli    *Client
	Module string
}

func (c *Client) Geth() *geth {
	return &geth{
		Cli:    c,
		Module: "proxy",
	}
}

// Returns the current price per gas in wei
func (g *geth) GasPrice() (price GasPrice, err error) {
	err = g.Cli.call(g.Module, "eth_gasPrice", M{}, &price)
	return
}

// Returns the number of most recent block
func (g *geth) BlockNumber() (number BlockNumber, err error) {
	err = g.Cli.call(g.Module, "eth_gasPrice", M{}, &number)
	return
}
