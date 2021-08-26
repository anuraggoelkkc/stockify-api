package alerts

var alertWorkerPool WorkerPool

const (
	MaxAlertWorker = 10
)

type AlertManger struct {
	alerts   map[string][]alert           //cache alerts
	channels map[ChannelType]AlertChannel //cache channels
}

func (a *AlertManger) AddFirebaseKeyPath(apiKey string) {
	channel := FCMChannel{}
	channel.AddProperty("apiKey", apiKey)
	a.channels[ChannelTypeFCM] = &channel
}

func (a *AlertManger) SetAlert(alertObj alert) {
	if _, ok := a.alerts[alertObj.AlertTopic]; ok == true {
		a.alerts[alertObj.AlertTopic] = append(a.alerts[alertObj.AlertTopic], alertObj)
	} else {
		a.alerts[alertObj.AlertTopic] = []alert{alertObj}
	}
}

//ToDo: Implement this, load alerts from db to cache
func (a *AlertManger) loadAlerts() {

}

// ToDo:Implement this, these worker pool will listen to push data and send message to notification queue
func (a *AlertManger) startAlertWorkers() {
	alertWorkerPool = NewWorkerPool(MaxAlertWorker)
	alertWorkerPool.Run()
}
func SendAlert(alertObj alert) {
	for k, v := range alertObj.ReceiverIds {
		if k == ChannelTypeFCM {
			TriggerFCMNotification(v)
		}
	}
}

//This method is called by producer and checked with alerts data
func (a *AlertManger) pushData(alertTopic string, key string, value interface{}) {
	//ToDo, make it concurrent and modular using goroutines
	if _, ok := a.alerts[alertTopic]; ok == true {
		for _, v := range a.alerts[alertTopic] {
			//call process trigger, in concurrent tasks
			alertWorkerPool.AddTask(func() {
				processTrigger(key, v, value)
			})
		}
	}
}

func processTrigger(key string, v alert, value interface{}) {
	triggerLogic := v.TriggerLogic
	if key == triggerLogic.FieldName && triggerLogic.FieldType == "int" {
		value1 := triggerLogic.Value.(int)
		value2 := value.(int)
		result := CheckTriggerCondition(value1, value2, triggerLogic.Operator)
		if result {
			// send alert if condition valid
			SendAlert(v)
		}
	}
}

func CheckTriggerCondition(value1, value2 int, operator string) bool {
	switch operator {
	case "<":
		return value1 < value2
	case "<=":
		return value1 <= value2
	case ">":
		return value1 > value2
	case ">=":
		return value1 >= value2
	}
	return false
}

func GetAlertManager() *AlertManger {
	alertManager := new(AlertManger)
	alertManager.loadAlerts()        // load previously stored alert
	alertManager.startAlertWorkers() // start alert workers

	StartNotificationWorker() // Start Notification workers
	return alertManager
}
