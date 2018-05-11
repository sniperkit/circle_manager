package synchronization

import (
	"time"

	"github.com/fatih/structs"
)

func IsEqual(obj1 interface{}, obj2 interface{}) bool {
	obj2Map := structs.Map(obj2)

	for key, value := range structs.Map(obj1) {
		if key == "ID" || key == "CreatedAt" || key == "UpdatedAt" {
			continue
		}

		switch obj2Value := obj2Map[key].(type) {
		case time.Time:
			if _, ok := value.(time.Time); !ok {
				return false
			}
			if obj2Value.Unix() != value.(time.Time).Unix() {
				return false
			}
		case uint:
			if _, ok := value.(uint); !ok {
				return false
			}
			if obj2Value != value.(uint) {
				return false
			}
		case string:
			if _, ok := value.(string); !ok {
				return false
			}
			if obj2Value != value.(string) {
				return false
			}
		case int:
			if _, ok := value.(int); !ok {
				return false
			}
			if obj2Value != value.(int) {
				return false
			}
		}
	}
	return true
}
