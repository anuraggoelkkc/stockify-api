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

func updateFirebaseCollections() error {
	ctx := context.Background()
	client, err := getFirebaseClient(ctx)
	defer client.Close()
/*	_, _, err = client.Collection("instruments").Add(ctx, map[string]interface{}{
		"Instrument_ID": 277876486,
		"Exchange_ID":  1085455,
		"Symbol":  "EURAPR22JUL22",
		"Name": "EURINR",
		"Price": 0,
		"Exchange": "BCD",
	})
	if err != nil {
		log.Panicf("Failed updating instruments collection: %v", err)
	}

	return err*/

	//Open a File
	csv_file, err := os.Open("instruments.csv")
	if err != nil {
		log.Panic(err)
	}
	defer csv_file.Close()
	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		log.Panic(err)
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
				_, _, err = client.Collection("instruments").Add(ctx, map[string]interface{}{
					"Instrument_ID": instrument.Instrument_ID,
					"Exchange_ID": instrument.Exchange_ID,
					"Symbol": instrument.Symbol,
					"Name": instrument.Name,
					"Price": instrument.Price,
					"Exchange":instrument.Exchange,
				})
				if err != nil {
					log.Panicf("Failed updating instruments collection: %v", err)
				}
			}

		}
	}

	return err

/*	_, _, err = client.Collection("intruments").Add(ctx, map[string]interface{}{
		"instrument_token": 277876486,
		"exchange_token":  1085455,
		"tradingsymbol":  "EURAPR22JUL22",
		"name": "EURINR",
		"last_price": 0,
	})
	if err != nil {
		log.Panic("Failed updating instruments collection: %v", err)
	}
	return err*/
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

func Init(){
	err := downloadInstrumentCSV("$HOME/stockify-api/instruments.csv","https://api.kite.trade/instruments")
	if err != nil {
		log.Panic(err)
	} else {
		err := updateFirebaseCollections()
		if err != nil {
			log.Panic(err)
		}
	}
}
