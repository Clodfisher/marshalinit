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

	t.Logf("unmarshal success, conf:%#v, port:%v", conf, conf.ServerConf.Port)

}
