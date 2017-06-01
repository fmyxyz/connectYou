package core

import (
	"bufio"
	"net"
)

type ConnContext struct {
	Reader     *bufio.Reader
	Writer     *bufio.Writer
	LocalAddr  net.Addr
	RemoteAddr net.Addr
}
