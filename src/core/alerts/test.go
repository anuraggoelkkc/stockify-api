package alerts

func TestAlerts() {

	stockNames := []string{
		"SBIN", "DMART", "TITAN",
	}
	alertManager := GetAlertManager()

	alertManager.AddTopics(stockNames...)
	alertManager.AddTopic("IOC")
	alertManager.AddFCMKey("123456")

	alert_1 := AlertBuilder().
		Name("dmart_price_gt_150").
		Topic("DMART").
		ReceiverId("123").
		Trigger(AlertTrigger{"price", "int", ">", 150}).
		EnableFCMAlerts().
		Build()

	alert_2 := AlertBuilder().
		Name("sbin_price_lt_100").
		Topic("SBIN").
		ReceiverId("345").
		Trigger(AlertTrigger{"price", "int", "<", 100}).
		EnableFCMAlerts().
		Build()

	alertManager.SetAlert(alert_1)
	alertManager.SetAlert(alert_2)

}
