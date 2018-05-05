# marshalinit   

### 创建意图        
本程序利用反射原理，将固有格式的配置文件加载到数据结构中，以及将固的数据结构数据加载到配置文件中。主要用于将公司将配置文件加载到程序内部。         

### 名词解释        
* 序列化：将数据结构中的数据，序列化成切片数据，可以输出到文件或是进行网络传输。    
* 反序列化：从文件或网络中获取到的切片数据，将其反序列化成数据结构的过程。    
* section：配置文件中的部分数据即[xxx]。    
* item： 部分数据下的子项，即key=value对。       

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

### 其主函数实现    
 对于此项目，主要的实现是两个函数，一个是反序列化，一个是序列化，其核心实现思想如下所示：    

* 反序列化    
  函数声明：    
```
func UnMarshal(data []byte, result interface{}) (err error)
```
  1. 将通过网络或文件读取获取到的切片数据，以及将要赋值的结构数据类型指针传入此函数。        
  2. 以回车符为每一行标识进行数据的切分，过滤行数据中的注释和空行部分。       
  4. 依据section的名字在result类型中，找到结构类型名字。      
  3. 判断和解析是否是切片数据中的section数据，提取section的名字。    
  5. 通过结构名字找到result中section的value部分。   
  6. 通过每行的item数据，提取出key和vlaue对，以及获取到key的名字。    
  7. 通过key名字，找到在相应section中的数据类型名字。        
  8. 根据section中的数据类型名字，找到section的value字段。    
  9. 将相应的值，根据不同的类型，赋值给相应的value字段。    
  10. 重复执行相应的每一行解析。     
  
* 序列化     
  函数声明：    
```
func Marshal(data interface{}) (result []byte, err error)
```    
 1. 将需要序列化成切片的结构体类型传入此函数。    
 2. 利用反射判定datal类型是否为结构体类型。    
 3. 获取每一个section的类型，判定是否为结构体类型。    
 4. 通过section类型中的tag获取其在切片数据中应该具有的名字。    
 5. 将名字和[]组成在切片中的section内容。    
 6. 将每个section下的所有item以同样的方式组成key=value对。    
 7. 将所有的序列化好的数据按照解析的顺序，写入到result切片中。    
 8. 完成真个序列化的过程。    