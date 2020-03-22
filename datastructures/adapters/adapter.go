package adapter

type charger interface {
	UseIPhoneCharger()
}

type AndroidPhone struct {
}

func (a *AndroidPhone) UseMicroUsbCharger() string {
	return "Charging with micro USB"
}

type IPhone struct {
}

func (a *IPhone) UseIPhoneCharger() string {
	return "Charging with IPhone charger"
}

type AndroidChargerAdapter struct {
	AndroidPhone *AndroidPhone
}

func (androidp *AndroidChargerAdapter) UseIPhoneCharger() string {
	return androidp.AndroidPhone.UseMicroUsbCharger()
}

type Client struct {
}

func (c *Client) InsertIphoneCharger(char charger) string {
	return char.UseIPhoneCharger()
}

// func main() {
// 	android := &AndroidPhone{}
// 	android.UseMicroUsbCharger()

// 	iphone := &IPhone{}
// 	iphone.UseIPhoneCharger()

// 	androidNew := &AndroidChargerAdapter{android}
// 	androidNew.UseIPhoneCharger()
// 	client := &Client{}
// 	client.InsertIphoneCharger(androidNew)
// 	client.InsertIphoneCharger(iphone)
// }
