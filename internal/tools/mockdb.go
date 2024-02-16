package tools

import "time"

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "5678DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "9012GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate a database query
	time.Sleep(1 * time.Second)
	var clientData = LoginDetails{}
	clientData = mockLoginDetails[username]
	if clientData.Username == "" {
		return nil
	}
	return &clientData
}
func (d *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(1 * time.Second)
	var clientData = CoinDetails{}
	clientData = mockCoinDetails[username]
	if clientData.Username == "" {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
