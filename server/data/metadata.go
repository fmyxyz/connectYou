package data

import (
	"bufio"
	"bytes"
	"log"
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
	//连接对应id
	ConnId string
}

//bigDataLength 默认 128k
func NewMetadata(bigDataLength int32, connId string) *Metadata {
	return &Metadata{BigDataLength: bigDataLength, ConnId: connId}
}

//打包
//包的大小小于128k
func (md *Metadata) Packing() []byte {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("debug")

	if md.BigDataLength == 0 {
		//128k
		md.BigDataLength = 1 << 17
	}
	bs := make([]byte, 0, 6+md.Length)
	log.Println("数据长度：", len(bs))
	buff := bytes.NewBuffer(bs)
	log.Println("写入数据。。。")
	i, err := buff.Write(Int32ToBytes(md.Length))
	if err != nil || i != 4 {
		log.Println("打包错误：", err)
	}
	log.Println("写入数据：", i)
	i, err = buff.Write(Uint8ToBytes(md.ReqType))
	if err != nil || i != 1 {
		log.Println("打包错误：", err)
	}
	log.Println("写入数据：", i)
	i, err = buff.Write(Uint8ToBytes(md.ResType))
	if err != nil || i != 1 {
		log.Println("打包错误：", err)
	}
	log.Println("写入数据：", i)
	if md.Length > md.BigDataLength {
		md.HasHandleData = false
	} else {
		md.HasHandleData = true
		i, err = buff.Write(md.Data)
		if err != nil || i != int(md.Length) {
			log.Println("打包错误：", err)
		}
		log.Println("写入数据：", i)
	}
	log.Println("数据：", len(bs), cap(bs))
	bb := buff.Bytes()
	log.Println("数据：", len(bb), cap(bb))
	return buff.Bytes()
}

//解包
//包的大小小于128k
func (md *Metadata) Unpacking(bufReader *bufio.Reader) error {
	bs := make([]byte, 6, 6)
	i, err := bufReader.Read(bs)
	if err != nil {
		return err
	} else if i != 6 {
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
		i, err = bufReader.Read(md.Data)
		if err != nil || i != int(md.Length) {
			log.Println("打包错误：", err)
		}
		log.Println("数据长度：", i, "数据：", string(md.Data))
	}
	return nil
}
