package hkhttp

import (
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
)

func NewServer(opts ...ServerOption) *http.Server {
	options := &serverOptions{}
	for _, opt := range opts {
		opt.apply(options)
	}
	return &http.Server{
		Addr:              options.addr,
		Handler:           options.handler,
		TLSConfig:         options.tlsConfig,
		ReadTimeout:       options.readTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      options.writeTimeout,
		IdleTimeout:       options.idleTimeout,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
}

func Serve() error {
	httpServer := NewServer()
	quicServer := &http3.Server{
		Server: httpServer,
	}
	handler := httpServer.Handler
	httpServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quicServer.SetQuicHeaders(w.Header())
		handler.ServeHTTP(w, r)
	})

	hErr := make(chan error)
	qErr := make(chan error)
	go func() {
		//hErr <- ListenTLS(httpServer)
	}()
	go func() {
		qErr <- quicServer.ListenAndServe()
	}()

	select {
	case err := <-hErr:
		quicServer.Close()
		return err
	case err := <-qErr:
		return err
	}
}
