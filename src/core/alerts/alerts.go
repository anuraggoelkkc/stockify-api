package alerts

type AlertTrigger struct {
	FieldName string
	FieldType string
	Operator  string
	Value     interface{}
}

type alert struct {
	AlertName          string
	AlertTopic         string
	RecevierId         string
	SubscribedChannels []ChannelType
	TriggerLogic       AlertTrigger
}

func AlertBuilder() *alert {
	return &alert{}
}

func (a *alert) Name(alertName string) *alert {
	a.AlertName = alertName
	return a
}

func (a *alert) Topic(topic string) *alert {
	a.AlertTopic = topic
	return a
}

func (a *alert) Trigger(triggerLogic AlertTrigger) *alert {
	a.TriggerLogic = triggerLogic
	return a
}

func (a *alert) ReceiverId(recieverId string) *alert {
	a.RecevierId = recieverId
	return a
}

func (a *alert) EnableFCMAlerts() *alert {
	a.SubscribedChannels = append(a.SubscribedChannels, FCMChannelType)
	return a
}

func (a *alert) Build() alert {
	return *a
}
