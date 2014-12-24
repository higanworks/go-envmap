package envmap

import (
	"os"
	"strings"
)

func All() map[string]string {
	data := os.Environ()
	items := make(map[string]string)
	for _, val := range data {
		splits := strings.SplitN(val, "=", 2)
		key := splits[0]
		value := splits[1]
		items[key] = value
	}
	return items
}

func ListKeys() []string {
	data := os.Environ()
	var keys []string
	for _, val := range data {
		splits := strings.SplitN(val, "=", 2)
		keys = append(keys, splits[0])
	}
	return keys
}
