package data

import (
	"bytes"
)

//打包
//包的大小小于128k
func Packing(md Metadata) (*bytes.Buffer, error) {
	//128k
	if md.Length > 1<<17 {
		return nil, DataTooLengthError{}
	}
	bs := make([]byte, 0, 6+md.Length)
	buff := bytes.NewBuffer(bs)
	buff.Write(Int32ToBytes(md.Length))
	buff.Write(Uint8ToBytes(md.ReqType))
	buff.Write(Uint8ToBytes(md.ResType))
	buff.Write(md.Data)
	return buff, nil
}

//解包
//包的大小小于128k
func Unpacking(buff *bytes.Buffer) (Metadata, error) {
	bs := make([]byte, 6, 6)
	i, err := buff.Read(bs)
	if err != nil {
		return nil, err
	} else if i != len(bs) {
		return nil, ProtocolError{}
	}
	md := Metadata{}
	md.Length = BytesToInt32(bs[:4])
	md.ReqType = BytesToUint8(bs[4:5])
	md.ResType = BytesToUint8(bs[5:6])
	//128k
	if md.Length > 1<<17 {
		return md, DataTooLengthError{}
	}
	md.Data = make([]byte, md.Length)
	buff.Read(md.Data)
	return md, nil
}

type ProtocolError struct {
}

func (p ProtocolError) Error() string {
	return "接受协议头错误。"
}

type DataTooLengthError struct {
}

func (p DataTooLengthError) Error() string {
	return "接受协议头错误。"
}
