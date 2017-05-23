package data

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
}
