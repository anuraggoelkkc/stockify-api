package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/csv"
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_struct "stockify-api/src/struct"
	"strconv"
	"strings"
)

func (f *FireStore) DownloadInstrumentCSV() (err error) {
	// Create the file
	out, err := os.Create(os.ExpandEnv(f.instrumentListFileLocation))
	if err != nil  {
		return err
	}
	defer out.Close()


	client := &http.Client{}
	req, err := http.NewRequest("GET", f.instrumentListUrl, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("X-Kite-Version", "3")
	req.Header.Set("Authorization", "token "+f.clientApiKey+":"+f.clientAccessToken)
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

func (f *FireStore) getFirebaseClient(ctx context.Context) (*firestore.Client, error) {
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

func (f *FireStore) AddOrUpdateUser(u _struct.User) error {
	//Create Firebase client
	ctx := context.Background()
	client, err := f.getFirebaseClient(ctx)
	defer client.Close()

	_, err = client.Collection("users").Doc(u.Id).Set(ctx, map[string]interface{}{
		"ID": u.Id,
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

/*func (f *FireStore) FetchAlerts(u string) ([]_struct.Alert, error) {

	//Create Firebase client
	ctx := context.Background()
	client, err := f.getFirebaseClient(ctx)
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
}*/

func (f *FireStore) AddAlert(a _struct.Alert) error {
	//Create Firebase client
	ctx := context.Background()
	client, err := f.getFirebaseClient(ctx)
	defer client.Close()

	//Add new Alert to alert collection
	_, _, err = client.Collection("alert").Add(ctx, map[string]interface{}{
		"Alert": a,
	})
	if err != nil {
		log.Panicf("Failed adding/updating alerts: %v", err)
		return err
	}
	return nil
}

func remove(s []_struct.Alert, i int) []_struct.Alert {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (f *FireStore) RemoveAlert(id string) error {
	//Create Firebase client
	ctx := context.Background()
	client, err := f.getFirebaseClient(ctx)
	defer client.Close()

	_, err = client.Collection("alert").Doc(id).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	//TODO: Remove from notification cache
	return nil
}

func (f *FireStore) UpdateInstrumentDetailToFireStore(o _struct.InstrumentDetail) error {

	return nil
}

func (f *FireStore) FetchInstrumentDetails(instrument string) (_struct.InstrumentDetail, error) {
	client := &http.Client{}
	url := f.instrumentDetailUrl
	res := _struct.InstrumentDetail{}
	url = strings.Replace(url, "EXCHANGE:SYMBOL", instrument, -1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("X-Kite-Version", "3")
	req.Header.Set("Authorization", "token "+f.clientApiKey+":"+f.clientAccessToken)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("unable to fetch details: %s", resp.Status)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var instrumentResp _struct.InstrumentDetailResponse
	err = json.Unmarshal(bodyBytes, &instrumentResp)
	if err == nil {
		return instrumentResp.Data[instrument], nil
	}
	return res, nil
}

func (f *FireStore) UpdateFirebaseCollections() error {

	//Create Firebase client
	ctx := context.Background()
	client, err := f.getFirebaseClient(ctx)
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
	var validEntries = 500
	for _, rec := range records {
		instrumentType := rec[9]
		if validEntries <= 0 {
			break
		}
		if strings.EqualFold(instrumentType, "EQ") {
			instrument.InstrumentID = rec[0]
			instrument.ExchangeID = rec[1]
			instrument.Symbol = rec[2]
			instrument.Name = rec[3]
			if price, err := strconv.ParseFloat(rec[4], 64); err == nil {
				instrument.Price = price
			}
			instrument.Exchange = rec[11]

			if len(instrument.Name) > 0 && (strings.EqualFold(instrument.Exchange, "BSE") || strings.EqualFold(instrument.Exchange, "NSE")) {
				_, err = client.Collection("instruments").Doc(instrument.InstrumentID).Set(ctx, map[string]interface{}{
					"Instrument_ID": instrument.InstrumentID,
					"Exchange_ID": instrument.ExchangeID,
					"Symbol": instrument.Symbol,
					"Name": instrument.Name,
					"Price": instrument.Price,
					"Exchange":instrument.Exchange,
				})
				if err != nil {
					log.Panicf("Failed updating instruments collection: %v", err)
					return err
				} else {
					validEntries--
				}
			}
		}
	}
	return nil
}

type FireStore struct {
	clientApiKey string
	clientAccessToken string
	instrumentListUrl string
	instrumentDetailUrl string
	instrumentListFileLocation string
}

func NewFireStore(apiKey string, accessToken string, instrumentListUrl string, instrumentDetailUrl string, instrumentListFileLocation string) *FireStore {
	return &FireStore{
		clientApiKey:      apiKey,
		clientAccessToken:  accessToken,
		instrumentListUrl: instrumentListUrl,
		instrumentDetailUrl: instrumentDetailUrl,
		instrumentListFileLocation: instrumentListFileLocation,
	}
}
