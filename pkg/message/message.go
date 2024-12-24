package message

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

/* An example valid message string would be
 * altenhofen|S|Hello World!
 */
type Message struct {
	Sender    string
	Action    byte // S for send, R for retrieve user info
	Content   *string
	Timestamp time.Time
}

func ParseMessage(message string) (*Message, error) {
	splitted := strings.Split(message, "|")
	sender := splitted[0]
	if sender == "" {
		sender = strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	action := []byte(splitted[1])[0]
	content := splitted[2]

	mes := Message{
		Sender:    sender,
		Action:    action,
		Content:   &content,
		Timestamp: time.Now(),
	}

	return &mes, nil
}

func (mes Message) ParseCommands() string {
	content := strings.Split(*mes.Content, " ")
	mesRes := fmt.Sprintf("command '%s' not found", content[0])

	if content[0] == "/username" {
		// users = append(users, content[1])
		log.Printf("User %s registrated", content[1])
		mesRes = fmt.Sprintf("User %s registrated", content[1])
	}

	return mesRes
}
