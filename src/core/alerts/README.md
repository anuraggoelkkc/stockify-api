# go-alerts
## A realtime alerting system build with golang and firestore.

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

go-alerts is a realtime alerting library, which support monitoring of streaming data. It provides methods to set trigger on relevant data point, and generates alerts on subscribed channels once data values matches trigger.

### Examples

- SBIN share price >=150 , trigger push notification
- DMART share price >=100 , alexa audio alert
- Temperture > 50 , send sms alert
- api response code = 401 , send mail alert

## Features
- Lightweight, Faster & efficient.
- Supports alerts on realtime streaming data.
- Golang's goroutines and channels to utlize maximum concurrency/parallelism.
- Firestore for persisting alerts and triggers.
- FCM channel support for mobile alerts.
- Support custom channels (Email, SMS, Alexa etc.) - TBD


## Tech
- Golang
- Firestore

## Usage
Setting an alert

```sh
	alertManager := GetAlertManager() // Get alertManager
	alertManager.AddFirebaseKeyPath("123456.txt") // Set firebase.json path

	recevierIdMap := make(map[ChannelType][]string)
	recevierIdMap[ChannelTypeFCM] = []string{"1234", "5678"} // ReceiverId and channel map

    // Building alert object
	alert_1, _ := AlertBuilder().
		Name("dmart_price_gt_150").
		Topic("DMART").
		Trigger(AlertTrigger{"price", "int", ">", 150}).
		EnableAlerts(recevierIdMap).
		Build()
    // Building alert object
	alert_2, _ := AlertBuilder().
		Name("sbin_price_lt_100"). // alert name
		Topic("SBIN").  // a topic for alerts
		Trigger(AlertTrigger{"price", "int", "<", 100}).// trigger logic
		EnableAlerts(recevierIdMap).
		Build()

	alertManager.SetAlert(alert_1) // adding alert object to alert Manager
	alertManager.SetAlert(alert_2)
```

Pushing streaming data
```sh
    // sending current price from data stream on a topic
    alertManager.pushData("DMART","price",150) 
    alertManager.pushData("DMART","volume",150)
```
## Flow Diagram

## License
**Free Software, Hell Yeah!**

