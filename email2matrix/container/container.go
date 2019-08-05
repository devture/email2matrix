package container

import (
	"devture-email2matrix/email2matrix/configuration"
	"devture-email2matrix/email2matrix/resolver"
	"devture-email2matrix/email2matrix/smtp"
	"fmt"

	"github.com/euskadi31/go-service"
	guerrilla "github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/sirupsen/logrus"
)

type ContainerShutdownHandler struct {
	destructors []func()
}

func (me *ContainerShutdownHandler) Add(destructor func()) {
	me.destructors = append(me.destructors, destructor)
}

func (me *ContainerShutdownHandler) Shutdown() {
	for i, _ := range me.destructors {
		me.destructors[len(me.destructors)-i-1]()
	}
}

func BuildContainer(
	configuration configuration.Configuration,
) (service.Container, *ContainerShutdownHandler) {
	container := service.New()
	shutdownHandler := &ContainerShutdownHandler{}

	// The logger is very crucial, so we're defining it outside
	logger := logrus.New()
	if configuration.Misc.Debug {
		logger.Level = logrus.DebugLevel
	}

	container.Set("logger", func(c service.Container) interface{} {
		return logger
	})

	container.Set("resolver", func(c service.Container) interface{} {
		return resolver.NewConfigurationBackedMailboxMappingInfoProvider(configuration.Matrix.Mappings)
	})

	container.Set("smtp.processor.email2matrix", func(c service.Container) interface{} {
		return smtp.Email2MatrixProcessor(logger, c.Get("resolver").(resolver.MailboxMappingInfoProvider))
	})

	container.Set("smtp.server", func(c service.Container) interface{} {
		cfg := &guerrilla.AppConfig{
			AllowedHosts: []string{configuration.Smtp.Hostname},
		}

		sc := guerrilla.ServerConfig{
			ListenInterface: configuration.Smtp.ListenInterface,
			IsEnabled:       true,
		}
		cfg.Servers = append(cfg.Servers, sc)

		additionalSaveProcess := ""
		if configuration.Misc.Debug {
			additionalSaveProcess = "|Debugger"
		}
		bcfg := backends.BackendConfig{
			"save_workers_size":  configuration.Smtp.Workers,
			"save_process":       fmt.Sprintf("HeadersParser|Header|Hasher%s|Email2Matrix", additionalSaveProcess),
			"log_received_mails": true,
		}
		cfg.BackendConfig = bcfg

		d := guerrilla.Daemon{Config: cfg}

		d.AddProcessor("Email2Matrix", c.Get("smtp.processor.email2matrix").(backends.ProcessorConstructor))

		shutdownHandler.Add(func() {
			logger.Debug("Shutdown SMTP server")
			d.Shutdown()
		})

		return d
	})

	return container, shutdownHandler
}
