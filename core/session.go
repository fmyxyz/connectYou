package core

import (
	"net"

	"bufio"

	"github.com/fmyxyz/connectYou/core/utils"
)

type Session struct {
	kv        map[string]interface{}
	id        string
	conn      net.Conn
	bufReader *bufio.Reader
	bufWriter *bufio.Writer
}

//生成session
func NewSession(conn net.Conn) *Session {
	bufReader := bufio.NewReader(conn)
	bufWriter := bufio.NewWriter(conn)
	session := &Session{
		kv:        make(map[string]interface{}),
		id:        utils.UniqueId(),
		conn:      conn,
		bufReader: bufReader,
		bufWriter: bufWriter,
	}
	return session
}

//获取session id
func (s *Session) GetSessionId() string {
	return s.id
}

//获取Conn
func (s *Session) GetConn() net.Conn {
	return s.conn
}

func (s *Session) GetAttr(k string) interface{} {
	return s.kv[k]
}

func (s *Session) PutAttr(k string, v interface{}) {
	s.kv[k] = v
}
