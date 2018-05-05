package marshalinit

import (
	"io/ioutil"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `init:"server"`
	MysqlConf  MysqlConfig  `init:"mysql"`
}

type ServerConfig struct {
	Ip   string `init:"ip"`
	Port uint   `init:"port"`
}

type MysqlConfig struct {
	Username string  `init:"username"`
	Passwd   string  `init:"passwd"`
	Database string  `init:"database"`
	Host     string  `init:"host"`
	Port     int     `init:"port"`
	Timeout  float32 `init:"timeout"`
}

func TestInitConfig(t *testing.T) {

	data, err := ioutil.ReadFile("./config.init")
	if err != nil {
		t.Error("read file failed")
	}
	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Errorf("unmarshal failed, err:%v", err)
		return
	}

	//以#v方式打印的uint是16进制数显示
	t.Logf("unmarshal success, conf:%#v, port:%v", conf, conf.ServerConf.Port)

	confData, err := Marshal(conf)
	if err != nil {
		t.Errorf("marshal failed, err:%v", err)
	}

	t.Logf("marshal succ, conf:%s", string(confData))

}

func TestIniConfigFile(t *testing.T) {

	//将数据结构的数据写入配置文件
	filename := "D:/tmp/test.conf"
	var conf Config
	conf.ServerConf.Ip = "localhost"
	conf.ServerConf.Port = 9009
	err := MarshalFile(filename, conf)
	if err != nil {
		t.Errorf("marshal failed, err:%v", err)
		return
	}

	//将配置文件的数据写入数据结构
	var conf2 Config
	err = UnMarshalFile(filename, &conf2)
	if err != nil {
		t.Errorf("unmarshal failed, err:%v", err)
	}

	t.Logf("unmarshal succ, conf:%#v", conf2)
}
