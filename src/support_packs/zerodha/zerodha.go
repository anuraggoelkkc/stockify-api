package zerodha

import (
	"fmt"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
)

const (
	apiKey    string = "3ag3ura5bykyu8a5"
	apiSecret string = "g8du8pu39k86ubqwktuwtkpjrg2ce8p2"
)

func GetAccessToken(client *kiteconnect.Client) string {
	if client != nil {
		fmt.Println(client.GetLoginURL())
		requestToken := "aUdk5fc3mAvrz8czJt4S1U1HpVyqnZQy"
		// Get user details and access token
		data, err := client.GenerateSession(requestToken, apiSecret)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return ""
		} else {
			return data.AccessToken
		}
	} else {
		return "YsQ3Nw9Uf8yA2TbJtmQfQKYHo3m8Kw7P"
	}
}

func Login() {
	// Create a new Kite connect instance
	kc := kiteconnect.New(apiKey)
	token := GetAccessToken(nil) //to generate new token pass kc instead of nil
	kc.SetAccessToken(token)

	// Get margins
	margins, err := kc.GetUserMargins()
	if err != nil {
		fmt.Printf("Error getting margins: %v", err)
	}
	fmt.Println("margins: ", margins)
}

func Init() {
	Login()
}
