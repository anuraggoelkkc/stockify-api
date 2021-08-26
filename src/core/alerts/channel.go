package alerts

type ChannelType string

const (
	FCMChannelType   ChannelType = "fcm"
	EmailChannelType ChannelType = "email"
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
