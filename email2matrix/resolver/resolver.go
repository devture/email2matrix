package resolver

import (
	"devture-email2matrix/email2matrix/configuration"
)

type ConfigurationBackedMailboxMappingInfoProvider struct {
	mappings []configuration.ConfigurationMatrixMapping
}

func NewConfigurationBackedMailboxMappingInfoProvider(mappings []configuration.ConfigurationMatrixMapping) *ConfigurationBackedMailboxMappingInfoProvider {
	return &ConfigurationBackedMailboxMappingInfoProvider{
		mappings: mappings,
	}
}

func (me *ConfigurationBackedMailboxMappingInfoProvider) Resolve(mailboxName string) *MailboxMappingInfo {
	for _, item := range me.mappings {
		if item.MailboxName == mailboxName {
			return &MailboxMappingInfo{
				MailboxName:         item.MailboxName,
				MatrixHomeserverUrl: item.MatrixHomeserverUrl,
				MatrixRoomId:        item.MatrixRoomId,
				MatrixUserId:        item.MatrixUserId,
				MatrixAccessToken:   item.MatrixAccessToken,
				IgnoreSubject:       item.IgnoreSubject,
				IgnoreBody:          item.IgnoreBody,
				SkipMarkdown:        item.SkipMarkdown,
			}
		}
	}

	return nil
}
