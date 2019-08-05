package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Smtp   ConfigurationSmtp
	Misc   ConfigurationMisc
	Matrix ConfigurationMatrix
}

type ConfigurationSmtp struct {
	ListenInterface string
	Hostname        string
	Workers         int
}

type ConfigurationMatrix struct {
	Mappings []ConfigurationMatrixMapping
}

type ConfigurationMatrixMapping struct {
	MailboxName         string
	MatrixRoomId        string
	MatrixHomeserverUrl string
	MatrixUserId        string
	MatrixAccessToken   string
	IgnoreSubject       bool
	IgnoreBody          bool
	SkipMarkdown        bool
}

type ConfigurationMisc struct {
	Debug bool
}

func LoadConfiguration(filePath string) (*Configuration, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return nil, err
	}

	if len(configuration.Matrix.Mappings) == 0 {
		return nil, fmt.Errorf("There should be at least one entry in Matrix.Mappings")
	}

	return &configuration, nil
}
