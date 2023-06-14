package amqprpc

import (
	"github.com/sesigl/go-project-ddd-template/internal/application"
	"github.com/sesigl/go-project-ddd-template/pkg/rabbitmq/rmq_rpc/server"
	"sync"
)

var hdlOnce sync.Once
var amqpRpcRouter map[string]server.CallHandler

// NewRouter -.
func NewRouter(t *application.TranslationUseCase) map[string]server.CallHandler {

	hdlOnce.Do(func() {
		amqpRpcRouter = make(map[string]server.CallHandler)
		{
			newTranslationRoutes(amqpRpcRouter, t)
		}
	})

	return amqpRpcRouter
}
