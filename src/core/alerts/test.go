package alerts

func TestAlerts() {

	alertManager := GetAlertManager()
	alertManager.AddFirebaseKeyPath("123456")

	recevierIdMap := make(map[ChannelType][]string)
	recevierIdMap[ChannelTypeFCM] = []string{"1234", "5678"}

	alert_1, _ := AlertBuilder().
		Name("dmart_price_gt_150").
		Topic("DMART").
		Trigger(AlertTrigger{"price", "int", ">", 150}).
		EnableAlerts(recevierIdMap).
		Build()

	alert_2, _ := AlertBuilder().
		Name("sbin_price_lt_100").
		Topic("SBIN").
		Trigger(AlertTrigger{"price", "int", "<", 100}).
		EnableAlerts(recevierIdMap).
		Build()

	alertManager.SetAlert(alert_1)
	alertManager.SetAlert(alert_2)

}
