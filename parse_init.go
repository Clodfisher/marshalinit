package marshalinit

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func parseSection(line string, typeStruct reflect.Type) (fieldName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error, invalid section %s", line)
		return
	}

	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section %s", line)
	}

	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if 0 == len(sectionName) {
			err = fmt.Errorf("syntax error, invalid section %s", line)
			return
		}

		for i := 0; i < typeStruct.NumField(); i++ {
			filed := typeStruct.Field(i)
			tagValue := filed.Tag.Get("init")
			if tagValue == sectionName {
				fieldName = filed.Name
				break
			}
		}
	}
	return
}

func parseItem(lastSectionName string, line string, result interface{}) error {
	index := strings.Index(line, "=")
	if -1 == index {
		return fmt.Errorf("sytax error, line:%s", line)
	}

	key := strings.TrimSpace(line[0:index])
	value := strings.TrimSpace(line[index+1:])

	if len(key) == 0 {
		return fmt.Errorf("sytax error, line:%s", line)
	}

	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastSectionName)
	if sectionValue == reflect.ValueOf(nil) {
		return nil
	}

	if sectionValue.Kind() != reflect.Struct {
		return fmt.Errorf("field:%s must be struct", lastSectionName)

	}

	sectionType := sectionValue.Type()
	var keyItemName string
	for i := 0; i < sectionType.NumField(); i++ {
		item := sectionType.Field(i)
		itemName := item.Tag.Get("init")
		if itemName == key {
			keyItemName = item.Name
			break
		}
	}

	if len(keyItemName) == 0 {
		return nil
	}

	itemValue := sectionValue.FieldByName(keyItemName)
	if itemValue == reflect.ValueOf(nil) {
		return nil
	}

	switch itemValue.Type().Kind() {
	case reflect.String:
		itemValue.SetString(value)
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		intVal, errRet := strconv.ParseInt(value, 10, 64)
		if errRet != nil {
			return errRet
		}
		itemValue.SetInt(intVal)

	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		intVal, errRet := strconv.ParseUint(value, 10, 64)
		if errRet != nil {
			return errRet

		}
		itemValue.SetUint(intVal)
	case reflect.Float32, reflect.Float64:
		floatVal, errRet := strconv.ParseFloat(value, 64)
		if errRet != nil {
			return errRet
		}

		itemValue.SetFloat(floatVal)

	default:
		return fmt.Errorf("unsupport type:%v", itemValue.Type().Kind())
	}

	return nil
}
