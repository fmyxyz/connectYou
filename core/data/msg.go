package data

type Message struct {
	MsgType    string `json:"msg_type"`
	Msg        string `json:"msg"`
	FromUserId string `json:"from_user_id"`
	ToUserId   string `json:"to_user_id"`
	FromConnId string `json:"from_conn_id"`
	ToConnId   string `json:"to_conn_id"`
}

type Cmd struct {
	CMD string
	Who string
}
