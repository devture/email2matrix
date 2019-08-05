package resolver

type MailboxMappingInfoProvider interface {
	Resolve(mailboxName string) *MailboxMappingInfo
}

type MailboxMappingInfo struct {
	MailboxName         string
	MatrixHomeserverUrl string
	MatrixRoomId        string
	MatrixUserId        string
	MatrixAccessToken   string
	IgnoreSubject       bool
	IgnoreBody          bool
	SkipMarkdown        bool
}
