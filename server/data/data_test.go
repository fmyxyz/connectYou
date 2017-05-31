package data

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestBytesInt32Convert(t *testing.T) {
	var i32 int32
	i32 = 11
	t.Log(i32)
	bs := Int32ToBytes(i32)
	t.Log(bs)
	t.Log(len(bs))
	t.Log(cap(bs))
	buff := bytes.NewBuffer(bs)
	t.Log(buff.Bytes())
	inew := BytesToInt32(bs)
	t.Log(inew)
	if i32 == inew {
		t.Log(i32 == inew)
	} else {
		t.Fail()
	}
}

func TestMsg(t *testing.T) {
	msg := new(Message)
	json.Unmarshal([]byte(`{
	"msg_type": "msg_type",
	"msg": "msg",
	"from_user_id": "from_user_id",
	"to_user_id": "to_user_id",
	"from_conn_id": "from_conn_id",
	"to_conn_id": "to_conn_id"
}`), msg)
	t.Log(msg)
}
