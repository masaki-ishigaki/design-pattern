package facade

import (
	"fmt"
	"log"
	"os"
)

func MakeWelcomePage(mailaddr, filename string) {
	mailprop := GetProperties("maildata")

	username, ok := mailprop[mailaddr]
	if !ok {
		log.Fatal("specified user does not exist!")
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("file couldn't be created!")
	}
	writer := NewHtmlWriter(f)
	writer.Title(fmt.Sprintf("%s's web page", username))
	writer.Paragraph(fmt.Sprintf("Welcome to %s's web page!", username))
	writer.Paragraph("Nice to meet you!")
	writer.MailTo(mailaddr, username)
	defer writer.Close()

	fmt.Printf("%s is created for %s %s", filename, mailaddr, username)
}
