package data

import (
	"bufio"
	"bytes"
	"net"
)

type Metadata struct {
	//数据长度
	//四个字节
	Length int32
	//请求类型
	//一个字节
	ReqType uint8
	//响应类型
	//一个字节
	ResType uint8
	//数据
	//Length个字节
	Data []byte
	//是否处理数据
	HasHandleData bool
	//大数据长度
	BigDataLength int32
	//链接
	Conn net.Conn
}

//bigDataLength 默认 128k
func NewMetadata(bigDataLength int32) Metadata {
	return Metadata{BigDataLength: bigDataLength}
}

//打包
//包的大小小于128k
func (md *Metadata) Packing() []byte {
	if md.BigDataLength == 0 {
		//128k
		md.BigDataLength = 1 << 17
	}
	bs := make([]byte, 0, 6+md.Length)
	buff := bytes.NewBuffer(bs)
	buff.Write(Int32ToBytes(md.Length))
	buff.Write(Uint8ToBytes(md.ReqType))
	buff.Write(Uint8ToBytes(md.ResType))
	if md.Length > md.BigDataLength {
		md.HasHandleData = false
	} else {
		md.HasHandleData = true
		buff.Write(md.Data)
	}
	return buff.Bytes()
}

//解包
//包的大小小于128k
func (md *Metadata) Unpacking(bufReader *bufio.Reader) error {
	bs := make([]byte, 6, 6)
	i, err := bufReader.Read(bs)
	if err != nil {
		return err
	} else if i != len(bs) {
		return ProtocolError{}
	}
	md.Length = BytesToInt32(bs[:4])
	md.ReqType = BytesToUint8(bs[4:5])
	md.ResType = BytesToUint8(bs[5:6])
	if md.BigDataLength == 0 {
		//128k
		md.BigDataLength = 1 << 17
	}
	if md.Length > md.BigDataLength {
		md.HasHandleData = false
	} else {
		md.HasHandleData = true
		md.Data = make([]byte, md.Length)
		bufReader.Read(md.Data)
	}
	return nil
}
