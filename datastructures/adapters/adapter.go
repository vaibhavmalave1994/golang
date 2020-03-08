package main

import "fmt"

type charger interface {
	UseIPhoneCharger()
}

type AndroidPhone struct {
}

func (a *AndroidPhone) UseMicroUsbCharger() {
	fmt.Println("Charging with micro USB")
}

type IPhone struct {
}

func (a *IPhone) UseIPhoneCharger() {
	fmt.Println("Charging with IPhone charger")
}

type AndroidChargerAdapter struct {
	AndroidPhone *AndroidPhone
}

func (androidp *AndroidChargerAdapter) UseIPhoneCharger() {
	androidp.AndroidPhone.UseMicroUsbCharger()
}

type Client struct {
}

func (c *Client) InsertIphoneCharger(char charger) {
	char.UseIPhoneCharger()
}

func main() {
	android := &AndroidPhone{}
	android.UseMicroUsbCharger()

	iphone := &IPhone{}
	iphone.UseIPhoneCharger()

	androidNew := &AndroidChargerAdapter{android}
	androidNew.UseIPhoneCharger()
	client := &Client{}
	client.InsertIphoneCharger(androidNew)
	client.InsertIphoneCharger(iphone)
}
