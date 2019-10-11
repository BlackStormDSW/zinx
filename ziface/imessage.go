package ziface

/*
	将请求的一个消息封装到message中，定义抽象层接口
 */
type IMessage interface {
	GetMsgId() uint32	//获取消息ID
	GetDataLen() uint32	//获取消息数据段长度
	GetCrc32Num() uint32//获取消息Crc32校验码
	GetData() []byte	//获取消息内容

	SetMsgId(uint32)	//设计消息ID
	SetDataLen(uint32)	//设置消息数据段长度
	SetCrc32Num(data []byte, poly uint32)	//设置消息Crc32校验码
	SetData([]byte)		//设计消息内容
}
