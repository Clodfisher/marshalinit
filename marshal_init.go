package marshalinit

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

/*
	将结构体中的数据，进行序列化成切换，将切片内容写入到文件内
*/
func MarshalFile(filename string, data interface{}) error {
	result, err := Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, result, 0755)
}

/*
 序列化：将数据结构中的数据，序列化成切片数据，可以输出到文件或是进行网络传输。
*/
func Marshal(data interface{}) ([]byte, error) {
	return nil, nil
}

/*
	从文件中读取数据，将读取的数据反序列化成结构体中的数据
*/
func UnMarshalFile(filename string, result interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return UnMarshal(data, result)
}

/*
	反序列化：从文件或网络中获取到的切片数据，将其反序列化成数据结构的过程。
*/
func UnMarshal(data []byte, result interface{}) error {

	//用于存储反序列化结果的变量是否为指针类型，否直接报错
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		return errors.New("please pass address for receive result")
	}

	//判断指针执行的内存存储的类型是否为结构体，否则直接包括
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		return errors.New("please pass address for struct")
	}

	lineArr := strings.Split(string(data), "\n")
	var lastSectionName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)

		//空行
		if 0 == len(line) {
			continue
		}

		//如果是注释，直接忽略
		if '#' == line[0] || ';' == line[0] {
			continue
		}

		//解析section，即[XXX]
		if '[' == line[0] {
			var err error
			lastSectionName, err = parseSection(line, typeStruct)
			if err != nil {
				return fmt.Errorf("%v lineno:%d", err, index+1)
			}
			continue
		}

		//解析item，即 key=value
		err := parseItem(lastSectionName, line, result)
		if err != nil {
			return fmt.Errorf("%v lineno:%d", err, index+1)
		}
	}

	return nil
}
