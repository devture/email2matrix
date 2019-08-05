package smtp

import (
	"devture-email2matrix/email2matrix/matrix"
	"devture-email2matrix/email2matrix/resolver"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"
)

var Email2MatrixProcessor = func(
	logger *logrus.Logger,
	resolver resolver.MailboxMappingInfoProvider,
) backends.ProcessorConstructor {
	return func() backends.Decorator {
		return func(p backends.Processor) backends.Processor {
			return backends.ProcessWith(
				func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
					if task == backends.TaskSaveMail {
						receiverMailbox := e.RcptTo[len(e.RcptTo)-1]

						mappingInfo := resolver.Resolve(receiverMailbox.User)
						if mappingInfo == nil {
							err := errors.New("Cannot find user")
							return backends.NewResult(fmt.Sprintf("450 Error: %s", err)), err
						}

						matrixEventId, err := matrix.Relay(e, *mappingInfo)
						if err != nil {
							return backends.NewResult(fmt.Sprintf("554 Error: %s", err)), err
						}

						logger.WithFields(logrus.Fields{
							"emailId":       e.QueuedId,
							"matrixRoomId":  mappingInfo.MatrixRoomId,
							"matrixEventId": matrixEventId,
						}).Infoln("Delivered")

						return p.Process(e, task)
					}
					return p.Process(e, task)
				},
			)
		}
	}
}
