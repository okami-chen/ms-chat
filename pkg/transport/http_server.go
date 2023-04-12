package transport

import (
	"context"
	"github.com/okamin-chen/chat/pkg/global"
	"github.com/okamin-chen/chat/pkg/initialize"
	"net/http"
)

type HttpServer struct {
	Addr string
	srv  *http.Server
}

func NewHttpServer() *HttpServer {
	router := initialize.InitRouter()
	addr := global.Conf.Server.GetPort()
	return &HttpServer{
		Addr: addr,
		srv: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *HttpServer) Type() Type {
	return TypeHTTP
}

func (s *HttpServer) Start(ctx context.Context) error {
	global.Log.Infof("HTTP Server listen: %s", s.Addr)
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			global.Log.Errorf("error http serve: %s", err)
		}
	}()
	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
