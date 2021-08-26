package alerts

type AlertManger struct {
	topics   []string
	alerts   map[string][]alert
	channels map[ChannelType]AlertChannel
}

func (a *AlertManger) AddFirebaseKeyPath(apiKey string) {
	channel := FCMChannel{}
	channel.AddProperty("apiKey", apiKey)
	a.channels[ChannelTypeFCM] = &channel
}

func (a *AlertManger) AddTopic(topic string) {
	a.topics = append(a.topics, topic)
}

func (a *AlertManger) AddTopics(topics ...string) {
	a.topics = append(a.topics, topics...)
}

func (a *AlertManger) SetAlert(alertObj alert) {
	if _, ok := a.alerts[alertObj.AlertTopic]; ok == true {
		a.alerts[alertObj.AlertTopic] = append(a.alerts[alertObj.AlertTopic], alertObj)
	} else {
		a.alerts[alertObj.AlertTopic] = []alert{alertObj}
	}
}

func (a *AlertManger) pushData(alertTopic string, key string, value interface{}) {
	//ToDo, make it concurrent and modular using goroutines
	if _, ok := a.alerts[alertTopic]; ok == true {
		for _, v := range a.alerts[alertTopic] {
			triggerLogic := v.TriggerLogic
			if key == triggerLogic.FieldName && triggerLogic.FieldType == "int" {
				value1 := triggerLogic.Value.(int)
				value2 := value.(int)
				result := ProcessTrigger(value1, value2, triggerLogic.Operator)
				if result {
					//push to notification queue
				}
			}
		}
	}
}

func ProcessTrigger(value1, value2 int, operator string) bool {
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
	return alertManager
}
