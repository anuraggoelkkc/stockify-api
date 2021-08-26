package zerodha

import (
	"fmt"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
	_struct "stockify-api/src/struct"
	"stockify-api/src/support_packs/firestore"
)

const (
	apiKey    string = "3ag3ura5bykyu8a5"
	apiSecret string = "g8du8pu39k86ubqwktuwtkpjrg2ce8p2"
)

func (z *Zerodha) getAccessToken(client *kiteconnect.Client) string {
	if client != nil {
		fmt.Println(client.GetLoginURL())
		requestToken := "l3BmkpNTUNOoEwtAQTag72iYRnvUo97k"
		// Get user details and access token
		data, err := client.GenerateSession(requestToken, apiSecret)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return ""
		} else {
			return data.AccessToken
		}
	} else {
		return "28VhkzVgY1OEKDMzFbhhC1FCh0g8vnqd"//need to update access token daily
	}
}

func (z *Zerodha) login() {
	// Create a new Kite connect instance
	kc := kiteconnect.New(apiKey)
	token := z.getAccessToken(nil) //to generate new token pass kc instead of nil
	kc.SetAccessToken(token)

	// Get margins
	margins, err := kc.GetUserMargins()
	if err != nil {
		fmt.Printf("Error getting margins: %v", err)
	}
	fmt.Println("margins: ", margins)
}

func (z *Zerodha) AddAlert(a _struct.Alert) error {
	return z.firebaseStore.AddAlert(a)
}

func (z *Zerodha) RemoveAlert(id string) error {
	return z.firebaseStore.RemoveAlert(id)
}

func (z *Zerodha) FetchAlerts(u string) ([]_struct.Alert, error) {
	//return z.firebaseStore.FetchAlerts(u)
	return nil,nil
}

func (z *Zerodha) ReloadInstrumentsInFirebase() error {
	return z.firebaseStore.UpdateFirebaseCollections()
}

func (z *Zerodha) FetchInstrumentDetails(exchange string, symbol string) (_struct.InstrumentDetail, error) {
	return z.firebaseStore.FetchInstrumentDetails(exchange+":"+symbol)
}

type Zerodha struct {
	firebaseStore *firestore.FireStore
}

func NewZerodha() *Zerodha {
	return &Zerodha{
		firebaseStore: firestore.NewFireStore(apiKey, "28VhkzVgY1OEKDMzFbhhC1FCh0g8vnqd", "https://api.kite.trade/instruments", "https://api.kite.trade/quote?i=EXCHANGE:SYMBOL", "$HOME/stockify-api/instruments.csv"),
	}
}

func (z *Zerodha) Init() {
	z.login()
//	z.firebaseStore.DownloadInstrumentCSV()
//	z.ReloadInstrumentsInFirebase()
	StartKiteTicker()
}
