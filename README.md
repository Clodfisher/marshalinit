# marshalinit   

### 创建意图        
本程序利用反射原理，将固有格式的配置文件加载到数据结构中，以及将固的数据结构数据加载到配置文件中。主要用于将公司将配置文件加载到程序内部。         

### 名词解释    
对于序列化和反序列化，不包括对于文件的读取和写入的过程。    
* 序列化：将配置文件的内容，加载到程序内部，最后转化为数据结构的过程。    
* 反序列化：将程序数据结构的内容，转换成固有的字符串，最终下发的配置文件中的整个过程。    

### 配置文件    
配置文件格式如下所示：    
```
      #this is comment
;this a comment
;[]表示一个section
[server]
ip=10.238.2.2
port = 8080

[    mysql]
username    =root
passwd = root
database=test
host=192.168.1.1
port=3838 
timeout=1.2
``` 

### 数据结构
数据结构格式如下所示：    
```
type Config struct {
	ServerConf ServerConfig
	MysqlConf  MysqlConfig 
}

type ServerConfig struct {
	Ip   string
	Port uint
}

type MysqlConfig struct {
	Username string 
	Passwd   string
	Database string
	Host     string
	Port     int   
	Timeout  float32
}
```

### 使用说明    
只要配置文件为【标题】加key-value格式接可以使用，结构体必须为一个大结构体，里面嵌套数个结构的格式。只需要调用此程序接口，既可以实现序列化和反序列化。      