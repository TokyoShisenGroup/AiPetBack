package config


type Config struct {
	StaticPath     PathConfig
	MsgChannelType MsgChannelType
}

// 相关地址信息，例如静态文件地址
type PathConfig struct {
	FilePath string
}

// 消息队列类型及其消息队列相关信息
// gochannel为单机使用go默认的channel进行消息传递
// kafka是使用kafka作为消息队列，可以分布式扩展消息聊天程序
type MsgChannelType struct {
	ChannelType string
	KafkaHosts  string
	KafkaTopic  string
}

var c Config

func GetConfig() Config {
	return c
}