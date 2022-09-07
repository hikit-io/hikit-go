package hkhttp

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/lucas-clemente/quic-go/http3"
)

func NewServer(opts ...ServerOption) *http.Server {
	options := &serverOptions{
		tlsConfig: GenTLSConfigNoErr(),
	}
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
	httpServer := NewServer(opts...)

	addr := httpServer.Addr
	handler := httpServer.Handler

	config := httpServer.TLSConfig

	if addr == "" {
		addr = ":https"
	}

	// Open the listeners
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer udpConn.Close()

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	tcpConn, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	defer tcpConn.Close()

	tlsConn := tls.NewListener(tcpConn, config)
	defer tlsConn.Close()

	if handler == nil {
		handler = http.DefaultServeMux
	}
	// Start the servers
	quicServer := &http3.Server{
		TLSConfig: config,
		Handler:   handler,
	}
	httpServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quicServer.SetQuicHeaders(w.Header())
		handler.ServeHTTP(w, r)
	})

	hErr := make(chan error)
	qErr := make(chan error)
	go func() {
		hErr <- httpServer.Serve(tlsConn)
	}()
	go func() {
		qErr <- quicServer.Serve(udpConn)
	}()

	select {
	case err := <-hErr:
		quicServer.Close()
		return err
	case err := <-qErr:
		return err
	}
}
