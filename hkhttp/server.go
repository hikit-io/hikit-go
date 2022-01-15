package hkhttp

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/lucas-clemente/quic-go/http3"
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

func Serve(opts ...ServerOption) error {
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
		tcpConn, err := net.Listen("tcp", httpServer.Addr)
		if err != nil {
			return
		}
		tlsConn := tls.NewListener(tcpConn, httpServer.TLSConfig)
		hErr <- httpServer.Serve(tlsConn)
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
