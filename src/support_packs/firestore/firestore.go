package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/csv"
	firebase "firebase.google.com/go"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	_struct "stockify-api/src/struct"
	"strconv"
	"strings"
)

func downloadInstrumentCSV(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(os.ExpandEnv(filepath))
	if err != nil  {
		return err
	}
	defer out.Close()


	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("X-Kite-Version", "3")
	req.Header.Set("Authorization", "token api_key:YsQ3Nw9Uf8yA2TbJtmQfQKYHo3m8Kw7P")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
		return err
	}

	return nil
}

func getFirebaseClient(ctx context.Context) (*firestore.Client, error) {
	// Use a service account
	conf := &firebase.Config{ProjectID: "stockify-8407f"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Panic(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Panic(err)
	}
	return client, err
}

func AddOrUpdateUser(u _struct.User) error {
	//Create Firebase client
	ctx := context.Background()
	client, err := getFirebaseClient(ctx)
	defer client.Close()

	_, err = client.Collection("user").Doc(u.ID).Set(ctx, map[string]interface{}{
		"ID": u.ID,
		"Name": u.Name,
		"Email": u.Email,
		"DeviceID": u.DeviceID,
	})
	if err != nil {
		log.Panicf("Failed adding/updating user: %v", err)
		return err
	}
	return err
}

func FetchAlerts(u string) ([]_struct.Alert, error) {
	//Create Firebase client
	ctx := context.Background()
	client, err := getFirebaseClient(ctx)
	defer client.Close()

	//Get Existing alerts for user
	dsnap, err := client.Collection("user_alert").Doc(u).Get(ctx)
	if err != nil {
		return nil, err
	}
	var existingAlerts []_struct.Alert
	err = dsnap.DataTo(&existingAlerts)
	fmt.Printf("Document data: %#v\n", existingAlerts)
	return existingAlerts, err
}

func AddAlert(a _struct.Alert) error {
	//Create Firebase client
	ctx := context.Background()
	client, err := getFirebaseClient(ctx)
	defer client.Close()

	//Get Existing alerts for user
	dsnap, err := client.Collection("user_alert").Doc(a.User_ID).Get(ctx)
	if err != nil {
		return err
	}
	var existingAlerts []_struct.Alert
	err = dsnap.DataTo(&existingAlerts)
	fmt.Printf("Document data: %#v\n", existingAlerts)
	if err != nil {
		return err
	}

	//Add new alert to existing alert
	existingAlerts = append(existingAlerts, a)

	_, err = client.Collection("user_alert").Doc(a.User_ID).Set(ctx, map[string]interface{}{
		"Alerts": existingAlerts,
	})
	if err != nil {
		log.Panicf("Failed adding/updating alerts: %v", err)
		return err
	}
	return err
}

//func RemoveAlert()

func UpdateFirebaseCollections() error {
	//Download CSV
	err := downloadInstrumentCSV("$HOME/stockify-api/instruments.csv","https://api.kite.trade/instruments")
	if err != nil {
		log.Panic(err)
		return err
	}

	//Create Firebase client
	ctx := context.Background()
	client, err := getFirebaseClient(ctx)
	defer client.Close()

	//Open File
	csv_file, err := os.Open("instruments.csv")
	if err != nil {
		log.Panic(err)
		return err
	}
	defer csv_file.Close()
	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		log.Panic(err)
		return err
	}

	var instrument = &_struct.Instruments{}
	for _, rec := range records {
		instrumentType := rec[9]
		if strings.EqualFold(instrumentType, "EQ") {
			instrument.Instrument_ID = rec[0]
			instrument.Exchange_ID = rec[1]
			instrument.Symbol = rec[2]
			instrument.Name = rec[3]
			if price, err := strconv.ParseFloat(rec[4], 64); err == nil {
				instrument.Price = price
			}
			instrument.Exchange = rec[11]

			if len(instrument.Name) > 0 && (strings.EqualFold(instrument.Exchange, "BSE") || strings.EqualFold(instrument.Exchange, "NSE")) {
				_, err = client.Collection("instruments").Doc(instrument.Instrument_ID).Set(ctx, map[string]interface{}{
					"Instrument_ID": instrument.Instrument_ID,
					"Exchange_ID": instrument.Exchange_ID,
					"Symbol": instrument.Symbol,
					"Name": instrument.Name,
					"Price": instrument.Price,
					"Exchange":instrument.Exchange,
				})
				if err != nil {
					log.Panicf("Failed updating instruments collection: %v", err)
					return err
				}
			}

		}
	}

	return nil
}

func Init(){
	UpdateFirebaseCollections()
}
