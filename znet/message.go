package znet

import "hash/crc32"

type MsgHead struct {
	Id      	uint32 //消息的ID
	DataLen 	uint32 //消息的长度
	Crc32Num	uint32 //Crc32校验码
}

type Message struct {
	Head	MsgHead
	Data    []byte //消息的内容
}

//创建一个Message消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Head: MsgHead{
			Id:     id,
			DataLen: uint32(len(data)),
			Crc32Num: crc32.Checksum(data, crc32.MakeTable(0xD5828281)),
		},
		Data:   data,
	}
}

//获取消息ID
func (msg *Message) GetMsgId() uint32 {
	return msg.Head.Id
}

//获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.Head.DataLen
}

//获取消息Crc32校验码
func (msg *Message) GetCrc32Num() uint32 {
	return msg.Head.Crc32Num
}

//获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

//设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.Head.DataLen = len
}

//设置消息ID
func (msg *Message) SetMsgId(msgId uint32) {
	msg.Head.Id = msgId
}

//设置消息Crc32校验码
func (msg *Message) SetCrc32Num(data []byte, poly uint32) {
	msg.Head.Crc32Num = crc32.Checksum(data, crc32.MakeTable(poly))
}

//设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
