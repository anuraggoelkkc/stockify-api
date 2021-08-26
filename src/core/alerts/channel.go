package alerts

type ChannelType string

const (
	ChannelTypeFCM   ChannelType = "fcm"
	ChannelTypeEmail ChannelType = "email"
	ChannelTypeSMS   ChannelType = "sms"
)

type AlertChannel interface {
	AddProperty(string, interface{}) error
}

// FCM channel
type FCMChannel struct {
	apiKey string
}

func (f *FCMChannel) AddProperty(s string, i interface{}) error {
	f.apiKey = i.(string)
	return nil
}

// Mail channel
type MailChannel struct {
	apiKey string
}

func (m *MailChannel) AddProperty(s string, i interface{}) error {
	m.apiKey = i.(string)
	return nil
}

// SMS channel
type SMSChannel struct {
	apiKey string
}

func (s *SMSChannel) AddProperty(k string, i interface{}) error {
	s.apiKey = i.(string)
	return nil
}
