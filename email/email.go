package main

import (
	"bytes"
	"fmt"
	"github.com/keighl/mandrill"
	"html/template"
)

func main() {
	fmt.Println("Asd")

	// t := template.New("emailtemplating")

	t, err := template.ParseFiles("emailtemplating.html")
	if err != nil {
		fmt.Println("1 here : ", err)
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		fmt.Println("1 here : ", err)
	}

	result := tpl.String()
	fmt.Println(result)
	message := &mandrill.Message{}
	message.AddRecipient("vaibhavmalave@ayopop.com", "Vaibhav Malave", "to")
	message.Subject = "Welcome to Ayopop. Account activation link"
	message.HTML = result
	SendEmail(message)

}

func SendEmail(message *mandrill.Message) {
	fmt.Println("into send email")
	Key := "V_Lg1e4O5BS5qpWwwCK9dQ"
	fmt.Println("into send email 2")
	if message.FromEmail == "" {
		message.FromEmail = "info@ayopop.com"
	}
	fmt.Println("message", message)
	if message.FromEmail == "" {
		message.FromName = "Ayopop Admin"
	}
	fmt.Println("message", message)
	client := mandrill.ClientWithKey(Key)
	fmt.Println("client", client)
	responses, err := client.MessagesSend(message)
	fmt.Println(responses, err)
}
