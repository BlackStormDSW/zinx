package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unsafe"
	"zinx/utils"
	"zinx/ziface"
)

//封包拆包类实例，暂时不需要成员
type DataPack struct {}

//封包拆包实例初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包头长度方法
func(dp *DataPack) GetHeadLen() uint32 {
	//Id uint32(4字节) +  DataLen uint32(4字节) + Crc32Num uint32(4字节)
	return uint32(unsafe.Sizeof(MsgHead{}))
}
//封包方法(压缩数据)
func(dp *DataPack) Pack(msg ziface.IMessage)([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写msgID
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写dataLen
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写Crc32校验码
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetCrc32Num()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetData()); err != nil {
		return nil ,err
	}

	return dataBuff.Bytes(), nil
}
//拆包方法(解压数据)
func(dp *DataPack) Unpack(binaryData []byte)(ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head的信息，得到dataLen和msgID
	msg := &Message{}

	//读msgID
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.Head.Id); err != nil {
		return nil, err
	}

	//读dataLen
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.Head.DataLen); err != nil {
		return nil, err
	}

	//读Crc32校验码
	if err := binary.Read(dataBuff, binary.BigEndian, &msg.Head.Crc32Num); err != nil {
		return nil, err
	}

	//判断dataLen的长度是否超出我们允许的最大包长度
	if (utils.GlobalObject.MaxPacketSize > 0 && msg.Head.DataLen > utils.GlobalObject.MaxPacketSize) {
		return nil, errors.New("Too large msg data recieved")
	}

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
