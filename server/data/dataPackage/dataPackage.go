package dataPackage

import (
	"bytes"

	"github.com/fmyxyz/connectYou/server/data"
)

//打包
//包的大小小于128k
func Packing(md data.Metadata) (*bytes.Buffer, error) {
	//128k
	if md.Length > 1<<17 {
		return nil, DataTooLengthError{}
	}
	bs := make([]byte, 0, 6+md.Length)
	buff := bytes.NewBuffer(bs)
	buff.Write(data.Int32ToBytes(md.Length))
	buff.Write(data.Uint8ToBytes(md.ReqType))
	buff.Write(data.Uint8ToBytes(md.ResType))
	buff.Write(md.Data)
	return buff, nil
}

//解包
//包的大小小于128k
func Unpacking(buff *bytes.Buffer) (data.Metadata, error) {
	bs := make([]byte, 6, 6)
	i, err := buff.Read(bs)
	if err != nil {
		return nil, err
	} else if i != len(bs) {
		return nil, ProtocolError{}
	}
	md := data.Metadata{}
	md.Length = data.BytesToInt32(bs[:4])
	md.ReqType = data.BytesToUint8(bs[4:5])
	md.ResType = data.BytesToUint8(bs[5:6])
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
