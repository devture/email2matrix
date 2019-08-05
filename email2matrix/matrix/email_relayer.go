package matrix

import (
	"devture-email2matrix/email2matrix/mime"
	"devture-email2matrix/email2matrix/resolver"
	"errors"
	"fmt"
	"strings"

	"github.com/matrix-org/gomatrix"
	blackfriday "gopkg.in/russross/blackfriday.v2"

	"github.com/flashmob/go-guerrilla/mail"
)

func Relay(envelope *mail.Envelope, mappingInfo resolver.MailboxMappingInfo) (string, error) {
	subject, body, err := mime.ExtractContentFromEmail(envelope.NewReader())
	if err != nil {
		return "", fmt.Errorf("Failed to parse email: %s", err)
	}

	messagePlainOrMarkdown := GenerateMessage(
		subject,
		body,
		mappingInfo.IgnoreSubject,
		mappingInfo.IgnoreBody,
		mappingInfo.SkipMarkdown,
	)

	if messagePlainOrMarkdown == "" {
		return "", fmt.Errorf("Refusing to send an empty message")
	}

	client, err := gomatrix.NewClient(mappingInfo.MatrixHomeserverUrl, mappingInfo.MatrixUserId, mappingInfo.MatrixAccessToken)
	if err != nil {
		return "", err
	}

	var event *gomatrix.RespSendEvent
	if mappingInfo.SkipMarkdown {
		event, err = client.SendText(mappingInfo.MatrixRoomId, messagePlainOrMarkdown)
	} else {
		messageAsHtml := string(blackfriday.Run([]byte(messagePlainOrMarkdown)))

		event, err = client.SendMessageEvent(mappingInfo.MatrixRoomId, "m.room.message", gomatrix.HTMLMessage{
			Body:          messagePlainOrMarkdown,
			MsgType:       "m.text",
			Format:        "org.matrix.custom.html",
			FormattedBody: messageAsHtml,
		})
	}

	if err != nil {
		errMessage := strings.Replace(err.Error(), mappingInfo.MatrixAccessToken, "REDACTED", -1)
		return "", errors.New(errMessage)
	}

	return event.EventID, nil
}
