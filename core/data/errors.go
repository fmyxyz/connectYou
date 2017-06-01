package data

type ProtocolError struct {
}

func (p ProtocolError) Error() string {
	return "接受协议头错误。"
}
