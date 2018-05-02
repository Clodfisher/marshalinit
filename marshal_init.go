package marshalinit

func MarshalFile(filename string, data interface{}) error {

}

/*
 序列化：将数据结构中的数据，序列化成切片数据，可以输出到文件或是进行网络传输。
*/
func Marshal(data interface{}) ([]byte, error) {

}

func UnMarshalFile(filename string, result interface{}) error {

}

/*
	反序列化：从文件或网络中获取到的切片数据，将其反序列化成数据结构的过程。
*/
func UnMarshal(data []byte, result interface{}) error {

}
