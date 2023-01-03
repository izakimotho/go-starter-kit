package log

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"github.com/rs/zerolog/log"
)

// NewHandlerWrapper return Log HandlerWrapper which  log Request with Context metadata
func NewHandlerWrapper() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
			md, _ := metadata.FromContext(ctx)
			log.Debug().Fields(map[string]interface{}{
				"category": "LogWrapper",
				"service":  req.Service(),
				"method":   req.Method(),
				"ctx":      md,
			}).Msg("Server-Side Handler")
			err = fn(ctx, req, rsp)
			//securityLog(ctx, req, rsp)
			return
		}
	}
}

// NewSubscriberWrapper return Log SubscriberWrapper which  log Request with Context metadata
func NewSubscriberWrapper() server.SubscriberWrapper {
	return func(fn server.SubscriberFunc) server.SubscriberFunc {
		return func(ctx context.Context, p server.Message) (err error) {
			md, _ := metadata.FromContext(ctx)
            log.Debug().Fields(map[string]interface{}{
				"category":    "LogWrapper",
				"Topic":       p.Topic(),
				"ContentType": p.ContentType(),
				"Payload":     p.Payload(),
				"ctx":         md,
			}).Msg("Server-Side Subscriber")
			err = fn(ctx, p)
			return
		}
	}
}

type clientWrapper struct {
	client.Client
}

func (l *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) (err error) {
	md, _ := metadata.FromContext(ctx)
    log.Debug().Fields(map[string]interface{}{
		"category": "LogWrapper",
		"service":  req.Service(),
		"method":   req.Method(),
		"ctx":      md,
	}).Msg("Client-Side Call")
	err = l.Client.Call(ctx, req, rsp, opts...)
	return
}

func (l *clientWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) (err error) {
	md, _ := metadata.FromContext(ctx)
    log.Debug().Fields(map[string]interface{}{
		"category":    "LogWrapper",
		"Topic":       p.Topic(),
		"ContentType": p.ContentType(),
		"Payload":     p.Payload(),
		"ctx":         md,
	}).Msg("Client-Side Publish")
	err = l.Client.Publish(ctx, p, opts...)
	return
}

// NewClientLogWrapper return client.Wrapper which log Requests with Context metadata
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
