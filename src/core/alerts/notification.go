package alerts

import (
	"github.com/appleboy/go-fcm"
	"log"
)

const (
	MaxNotificationWorker = 10
)

var notificationWorkerPool WorkerPool

//Todo: Implement this
func TriggerFCMNotification(recevierIds []string) {

}

func sendFCMNotification(title, body, device_token string) {
	notificationWorkerPool.AddTask(func() {
		msg := &fcm.Message{
			To: device_token,
			Data: map[string]interface{}{
				"Title": title,
				"Body":  body,
			},
			Notification: &fcm.Notification{
				Title: title,
				Body:  body,
			},
		}

		// Create a FCM client to send the message.
		client, err := fcm.NewClient("sample_api_key")
		if err != nil {
			log.Fatalln(err)
		}

		// Send the message and receive the response without retries.
		_, err = client.Send(msg)
		if err != nil {
			log.Fatalln(err)
		}
	})
}

func StartNotificationWorker() {

	// For monitoring purpose.
	//waitC := make(chan bool)
	//
	//go func() {
	//	for {
	//		log.Printf("[main] Total current goroutine: %d", runtime.NumGoroutine())
	//		time.Sleep(1 * time.Second)
	//	}
	//}()

	// Start Worker Pool.
	notificationWorkerPool = NewWorkerPool(MaxNotificationWorker)
	notificationWorkerPool.Run()

	//<-waitC
}
