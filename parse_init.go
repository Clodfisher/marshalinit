package marshalinit

import (
	"fmt"
	"reflect"
	"strings"
)

func parseSection(line string, typeInfo reflect.Type) (fieldName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax error, invalid section %s", line)
		return
	}

	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section %s", line)
	}

	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if 0 == len(fieldName) {
			err = fmt.Errorf("syntax error, invalid section %s", line)
			return
		}

		for i := 0; i < typeInfo.NumField(); i++ {
			filed := typeInfo.Field(i)
			tagValue := filed.Tag.Get("init")
			if tagValue == sectionName {
				fieldName = filed.Name
				break
			}
		}
	}
	return
}
