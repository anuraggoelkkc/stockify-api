package alerts

import "errors"

type AlertTrigger struct {
	FieldName string
	FieldType string
	Operator  string
	Value     interface{}
}

func (t *AlertTrigger) Validate() error {
	if t.FieldName == "" {
		return errors.New("FieldName not defined")
	} else if t.FieldType == "" {
		return errors.New("FieldType not defined")
	} else if t.Operator == "" {
		return errors.New("Operator not defined")
	} else if t.Value == nil {
		return errors.New("Value not defined")
	}
	return nil
}

type alert struct {
	AlertName    string
	AlertTopic   string
	ReceiverId   map[ChannelType][]string
	TriggerLogic AlertTrigger
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

func (a *alert) EnableAlerts(receiverIdMap map[ChannelType][]string) *alert {
	a.ReceiverId = receiverIdMap
	return a
}

// Save to db, implment this
func (a *alert) saveToDB() error {
	return nil
}

func (a *alert) Build() (alert, error) {
	if a.AlertName == "" {
		return alert{}, errors.New("Name not defined")
	} else if a.AlertTopic == "" {
		return alert{}, errors.New("Topic not defined")
	} else if a.TriggerLogic.Validate() != nil {
		return alert{}, errors.New("TriggerLogic not defined")
	}
	return *a, nil
}
