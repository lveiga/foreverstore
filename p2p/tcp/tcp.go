package p2p

import (
	"fmt"
	"net"
	"sync"

	p2p "github.com/desenv/foreverstore/p2p"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.Mutex
	peers         map[net.Addr]p2p.Peer
}

func NewTCPeer(c net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     c,
		outbound: outbound,
	}
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	conn, err := t.listener.Accept()
	if err != nil {
		fmt.Printf("TCP accept error:, %s\n", err)
	}

	t.handleConn(conn)
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("new incoming connection %v\n", conn)
}
